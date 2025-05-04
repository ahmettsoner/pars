package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/option"
	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_AttributeGroup_GroupNameOnly_Inline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
RefGroup: CMD
`

	// Act

	var data objectresource.AttributeGroup
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAttributeGroup(objectresource.NewGroupIdentifier("CMD"), 0, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_AttributeGroup_GroupNameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
RefGroup:
  Name: CMD
`

	// Act

	var data objectresource.AttributeGroup
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAttributeGroup(objectresource.NewGroupIdentifier("CMD"), 0, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_AttributeGroup_Order(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
RefGroup: CMD
Order: 3
`

	// Act

	var data objectresource.AttributeGroup
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAttributeGroup(objectresource.NewGroupIdentifier("CMD"), 3, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_AttributeGroup_Options(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
RefGroup: CMD
Options:
- row=1
- column=3
`

	// Act

	var data objectresource.AttributeGroup
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewAttributeGroup(objectresource.NewGroupIdentifier("CMD"), 0, []option.Option{
		option.NewOption("row", "1"),
		option.NewOption("column", "3"),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_AttributeGroup_WithoutGroup(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Order: 3
`

	// Act

	var data objectresource.AttributeGroup
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Group", fieldRequiredErr.FieldName)
}
