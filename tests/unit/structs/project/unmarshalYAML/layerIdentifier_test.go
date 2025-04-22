package unmarshalYAML

import (
	"testing"

	layerPkg "parsdevkit.net/structs/layer"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_UnMarshall_LayerIdentifier_NameOnly(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Name: CMD
`

	// Act

	var data layerPkg.LayerIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := layerPkg.NewLayerIdentifier(0, "CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_LayerIdentifier_NameOnlyInline(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
CMD
`

	// Act

	var data layerPkg.LayerIdentifier
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := layerPkg.NewLayerIdentifier(0, "CMD")

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
