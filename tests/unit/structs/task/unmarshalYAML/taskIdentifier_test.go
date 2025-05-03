package unmarshalYAML

import (
	"testing"

	commontask "parsdevkit.net/structs/task/common-task"

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

	var data commontask.TaskIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := commontask.NewTaskIdentifier(0, "CMD", "")

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

	var data commontask.TaskIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := commontask.NewTaskIdentifier(0, "CMD", "")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
