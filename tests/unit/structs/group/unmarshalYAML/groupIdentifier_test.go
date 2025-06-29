package unmarshalYAML

import (
	"testing"

	applicationGroup "parsdevkit.net/application/structs/group"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_GroupIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data applicationGroup.GroupIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationGroup.NewGroupIdentifier(0, "CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_GroupIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data applicationGroup.GroupIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := applicationGroup.NewGroupIdentifier(0, "CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
