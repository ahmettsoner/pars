package schema

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
)

func LoadTemplate(yamlData []byte) (schemas.SchemaInterface, error) {
	var header schemas.SchemaHeader
	if err := yaml.Unmarshal(yamlData, &header); err != nil {
		return nil, fmt.Errorf("failed to unmarshal header: %w", err)
	}

	target, err := schemas.Get(header.GetKey())
	if err != nil {
		return nil, fmt.Errorf("Cannot load template for %s: %w", header.GetKey(), err)
	}

	if err := yaml.Unmarshal(yamlData, target); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body into %T: %w", target, err)
	}

	return target, nil
}

func GetAllManifestFilesInPath(path ...string) ([]schemas.SchemaInterface, error) {

	allFiles, err := file.GetAllFilesInPath(path...)
	if err != nil {
		return nil, fmt.Errorf("Error processing file paths: %v\n%w", allFiles, err)
	}

	schemas := make([]schemas.SchemaInterface, 0)
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

func GenerateManifestFilesFromTemplate(data any, templateFiles ...string) ([]schemas.SchemaInterface, error) {

	schemas := make([]schemas.SchemaInterface, 0)
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
