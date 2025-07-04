package unmarshalYAML

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
	applicationproject "parsdevkit.net/modules/project/application_project_payload"
)

func Test_UnMarshall_Header_Basic(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Project
Kind: Application
Name: CMD
Metadata:
`

	// Act

	var data schemas.SchemaHeader
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := schemas.SchemaHeader{
		Type:     schemas.StructTypes.Project,
		Kind:     applicationproject.PROJECT_KIND,
		Name:     "CMD",
		Metadata: schemas.Metadata{},
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Header_WithMetadataTags(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Project
Kind: Application
Name:  CMD
Metadata:
  Tags:
  - foo
  - bar
`

	// Act

	var data schemas.SchemaHeader
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := schemas.SchemaHeader{
		Type: schemas.StructTypes.Project,
		Kind: applicationproject.PROJECT_KIND,
		Name: "CMD",
		Metadata: schemas.Metadata{
			Tags: []string{"foo", "bar"},
		},
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
