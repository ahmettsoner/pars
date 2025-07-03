package unmarshalYAML

import (
	"testing"

	group_payload "parsdevkit.net/modules/group/basic_group_payload"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_GroupSpecification_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data group_payload.GroupSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group_payload.NewGroupSpecification(0, "CMD", "CMD", []string{"CMD"})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_GroupSpecification_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data group_payload.GroupSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group_payload.NewGroupSpecification(0, "CMD", "CMD", []string{"CMD"})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_GroupSpecification_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Path: parsdevkit.net/cmd
`

	// Act

	var data group_payload.GroupSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group_payload.NewGroupSpecification(0, "CMD", "parsdevkit.net/cmd", []string{"CMD"})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_GroupSpecification_Package_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Package: pars/cmd
`

	// Act

	var data group_payload.GroupSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group_payload.NewGroupSpecification(0, "CMD", "CMD", []string{"pars", "cmd"})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_GroupSpecification_Package_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Package: 
- pars
- cmd
`

	// Act

	var data group_payload.GroupSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group_payload.NewGroupSpecification(0, "CMD", "CMD", []string{"pars", "cmd"})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_GroupSpecification_ID_ShouldBeZero(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Id: 100
Name: CMD
`

	// Act

	var data group_payload.GroupSpecification
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := group_payload.NewGroupSpecification(0, "CMD", "CMD", []string{"CMD"})

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
