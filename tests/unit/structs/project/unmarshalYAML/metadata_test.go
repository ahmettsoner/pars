package unmarshalYAML

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"parsdevkit.net/application/schemas"
)

func Test_UnMarshall_Metadata_Tags_SingleLine(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Tags: tag1, tag2
`

	// Act

	var data schemas.Metadata
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := schemas.Metadata{
		Tags: []string{"tag1", "tag2"},
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
func Test_UnMarshall_Metadata_Tags_Array(t *testing.T) {

	// Arrange
	a := assert.New(t)
	yamlData := `
Tags: 
- tag1
- tag2
`

	// Act

	var data schemas.Metadata
	err := yaml.Unmarshal([]byte(yamlData), &data)

	expected := schemas.Metadata{
		Tags: []string{"tag1", "tag2"},
	}

	// Assert
	a.NoError(err)
	a.Equal(expected, data)
}
