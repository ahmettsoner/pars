package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
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

	var data structs.Header
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.Header{
		Type:     structs.StructTypes.Project,
		Name:     "CMD",
		Metadata: structs.Metadata{},
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

	var data structs.Header
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := structs.Header{
		Type: structs.StructTypes.Project,
		Name: "CMD",
		Metadata: structs.Metadata{
			Tags: []string{"foo", "bar"},
		},
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
