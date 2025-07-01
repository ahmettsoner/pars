package application

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
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utilities"
	"parsdevkit.net/core/utils"
	group "parsdevkit.net/modules/group/group"

	applicationproject "parsdevkit.net/structs/project/application-project"
	dataresource "parsdevkit.net/structs/resource/data-resource"
	objectsource "parsdevkit.net/structs/resource/object-resource"
	codetemplate "parsdevkit.net/structs/template/code-template"
	filetemplate "parsdevkit.net/structs/template/file-template"
	sharedtemplate "parsdevkit.net/structs/template/shared-template"
)

var registry = map[string]func() schemas.Schema{
	"Group":               func() schemas.Schema { return &group.GroupBaseStruct{} },
	"Project.Application": func() schemas.Schema { return &applicationproject.ProjectBaseStruct{} },
	"Resource.Data":       func() schemas.Schema { return &dataresource.ResourceBaseStruct{} },
	"Resource.Object":     func() schemas.Schema { return &objectsource.ResourceBaseStruct{} },
	"Template.Code":       func() schemas.Schema { return &codetemplate.TemplateBaseStruct{} },
	"Template.File":       func() schemas.Schema { return &filetemplate.TemplateBaseStruct{} },
	"Template.Shared":     func() schemas.Schema { return &sharedtemplate.TemplateBaseStruct{} },
}

func LoadTemplate(yamlData []byte) (schemas.Schema, error) {
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

func GetAllManifestFilesInPath(path ...string) ([]schemas.Schema, error) {

	allFiles, err := utilities.GetAllFilesInPath(path...)
	if err != nil {
		return nil, fmt.Errorf("Error processing file paths: %v\n%w", allFiles, err)
	}

	schemas := make([]schemas.Schema, 0)
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

func GenerateManifestFilesFromTemplate(data any, templateFiles ...string) ([]schemas.Schema, error) {

	schemas := make([]schemas.Schema, 0)
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
