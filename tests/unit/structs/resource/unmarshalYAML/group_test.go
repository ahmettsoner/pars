package unmarshalYAML

import (
	"testing"

	"parsdevkit.net/structs/option"
	objectresource "parsdevkit.net/structs/resource/object-resource"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Group_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewGroup("CMD", objectresource.Message{}, 0, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Group_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewGroup("CMD", objectresource.Message{}, 0, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Group_Title_ByText(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Title: title_text
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewGroup("CMD", objectresource.NewMessage("title_text", objectresource.NewDictionaryIdentifier("")), 0, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Group_Title_ByDictionaryReference(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Title: 
  RefMessage: groups_patient_filter
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewGroup("CMD", objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("groups_patient_filter")), 0, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Group_Order(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Order: 3
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewGroup("CMD", objectresource.Message{}, 3, []option.Option(nil))

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Group_Options(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Order: 3
Options:
- foo=bar
- Key: foo2
  Value: bar2
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewGroup("CMD", objectresource.Message{}, 3, []option.Option{
		option.NewOption("foo", "bar"),
		option.NewOption("foo2", "bar2"),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Group_Complete(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Title:
  RefMessage: groups_patient_filter
Order: 3
Options:
- foo=bar
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := objectresource.NewGroup("CMD", objectresource.NewMessage("", objectresource.NewDictionaryIdentifier("groups_patient_filter")), 3, []option.Option{
		option.NewOption("foo", "bar"),
	})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Group_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Title:
  RefMessage: groups_patient_filter
Order: 3
`

	// Act

	var data objectresource.Group
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Group.Name", fieldRequiredErr.FieldName)
}
