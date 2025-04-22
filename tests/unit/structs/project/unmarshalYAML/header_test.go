package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs"

	"parsdevkit.net/core/errors"

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

func Test_UnMarshall_Header_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Type: Project
Kind: Application
Metadata:
`

	// Act

	var data structs.Header
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Name", fieldRequiredErr.FieldName)
}

func Test_UnMarshall_Header_WithoutType(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Metadata:
`

	// Act

	var data structs.Header
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Type", fieldRequiredErr.FieldName)
}
