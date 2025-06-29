package unmarshalYAML

import (
	"testing"

	applicationTask "parsdevkit.net/application/structs/task"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_TaskIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data applicationTask.TaskIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTask.NewTaskIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_TaskIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data applicationTask.TaskIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationTask.NewTaskIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
