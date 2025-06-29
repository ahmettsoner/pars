package unmarshalYAML

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	applicationProject "parsdevkit.net/application/structs/project"
)

func Test_UnMarshall_Layer_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Persistence:Data:Repository
`

	// Act

	var data applicationProject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewLayer(0, "Persistence:Data:Repository", "Persistence/Data/Repository", []string{"Persistence:Data:Repository"}, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Layer_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Persistence:Data:Repository
`

	// Act

	var data applicationProject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewLayer(0, "Persistence:Data:Repository", "Persistence/Data/Repository", []string{"Persistence:Data:Repository"}, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Layer_Path(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Path: parsdevkit.net/cmd
`

	// Act

	var data applicationProject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewLayer(0, "CMD", "parsdevkit.net/cmd", []string{"CMD"}, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Layer_Package_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Package: pars/cmd
`

	// Act

	var data applicationProject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewLayer(0, "CMD", "CMD", []string{"pars", "cmd"}, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Layer_Package_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
Package: 
- pars
- cmd
`

	// Act

	var data applicationProject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationProject.NewLayer(0, "CMD", "CMD", []string{"pars", "cmd"}, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
