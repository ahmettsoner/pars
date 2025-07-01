package orchestrator

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/core/utilities/file"
	"parsdevkit.net/core/utils"
	group "parsdevkit.net/modules/group/group"

	"parsdevkit.net/application/contracts"
	applicationproject "parsdevkit.net/structs/project/application-project"
	dataresource "parsdevkit.net/structs/resource/data-resource"
	objectsource "parsdevkit.net/structs/resource/object-resource"
	codetemplate "parsdevkit.net/structs/template/code-template"
	filetemplate "parsdevkit.net/structs/template/file-template"
	sharedtemplate "parsdevkit.net/structs/template/shared-template"
)

var registry = map[string]func() contracts.SchemaInterface{
	"Group":               func() contracts.SchemaInterface { return &group.GroupBaseStruct{} },
	"Project.Application": func() contracts.SchemaInterface { return &applicationproject.ProjectBaseStruct{} },
	"Resource.Data":       func() contracts.SchemaInterface { return &dataresource.ResourceBaseStruct{} },
	"Resource.Object":     func() contracts.SchemaInterface { return &objectsource.ResourceBaseStruct{} },
	"Template.Code":       func() contracts.SchemaInterface { return &codetemplate.TemplateBaseStruct{} },
	"Template.File":       func() contracts.SchemaInterface { return &filetemplate.TemplateBaseStruct{} },
	"Template.Shared":     func() contracts.SchemaInterface { return &sharedtemplate.TemplateBaseStruct{} },
}

func LoadTemplate(yamlData []byte) (contracts.SchemaInterface, error) {
	var header schemas.SchemaHeader
	if err := yaml.Unmarshal(yamlData, &header); err != nil {
		return nil, fmt.Errorf("failed to unmarshal header: %w", err)
	}

	var key string
	if header.Kind != "" {
		key = fmt.Sprintf("%s.%s", header.Type, header.Kind)
	} else {
		key = string(header.Type)
	}

	factory, ok := registry[key]
	if !ok {
		log.Fatalf("Unknown Type/Kind: %s", key)
	}

	target := factory()

	if err := yaml.Unmarshal(yamlData, target); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into %T: %w", target, err)
	}

	return target, nil
}

func GetAllManifestFilesInPath(path ...string) ([]contracts.SchemaInterface, error) {

	allFiles, err := file.GetAllFilesInPath(path...)
	if err != nil {
		return nil, fmt.Errorf("Error processing file paths: %v\n%w", allFiles, err)
	}

	schemas := make([]contracts.SchemaInterface, 0)
	for _, file := range allFiles {

		stringData, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		yamlLines := strings.Split(string(stringData), "---")

		for _, line := range yamlLines {
			schema, err := LoadTemplate([]byte(line))
			if err != nil {
				return nil, err
			}

			if schema != nil {
				schemas = append(schemas, schema)
			}
		}
	}

	return schemas, nil
}

func GenerateManifestFilesFromTemplate(data any, templateFiles ...string) ([]contracts.SchemaInterface, error) {

	schemas := make([]contracts.SchemaInterface, 0)
	logrus.Debugf("found %v template(s) to create project", len(templateFiles))
	for _, templateFilePath := range templateFiles {

		var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), templateFilePath)
		tmplContent, err := os.ReadFile(tmplFile)
		if err != nil {
			log.Fatal(err)
		}
		var outputBuffer bytes.Buffer
		err = template.Must(template.New("SchemaFromTemplate").Parse(string(tmplContent))).Execute(&outputBuffer, data)
		if err != nil {
			log.Fatal(err)
		}
		stringData := outputBuffer.String()

		yamlLines := strings.Split(string(stringData), "---")

		for _, line := range yamlLines {
			schema, err := LoadTemplate([]byte(line))
			if err != nil {
				return nil, err
			}

			if schema != nil {
				schemas = append(schemas, schema)
			}
		}
	}
	logrus.Debugf("found %v schema(s) to create", len(schemas))

	return schemas, nil
}
