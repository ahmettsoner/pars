package unmarshalYAML

import (
	"testing"

	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/core/errors"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_Layer_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: Persistence:Data:Repository
`

	// Act

	var data applicationproject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLayer(0, "Persistence:Data:Repository", "Persistence/Data/Repository", []string{"Persistence:Data:Repository"}, "")

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

	var data applicationproject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLayer(0, "Persistence:Data:Repository", "Persistence/Data/Repository", []string{"Persistence:Data:Repository"}, "")

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

	var data applicationproject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLayer(0, "CMD", "parsdevkit.net/cmd", []string{"CMD"}, "")

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

	var data applicationproject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLayer(0, "CMD", "CMD", []string{"pars", "cmd"}, "")

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

	var data applicationproject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationproject.NewLayer(0, "CMD", "CMD", []string{"pars", "cmd"}, "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}

func Test_UnMarshall_Layer_WithoutName(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Path: CMD
`

	// Act

	var data applicationproject.Layer
	err := yaml.Unmarshal([]byte(yamlData), &data)

	// Assert
	a.Error(err)
	var fieldRequiredErr *errors.ErrFieldRequired
	a.ErrorAs(err, &fieldRequiredErr)
	a.Equal("Name", fieldRequiredErr.FieldName)
}
