package group

import (
	"os"
	"strings"

	"parsdevkit.net/structs"
	"parsdevkit.net/structs/group"

	"gopkg.in/yaml.v3"
)

type GroupSerializer struct{}

func (s GroupSerializer) GetGroupStructsFromString(data string) ([]group.GroupBaseStruct, error) {
	groups := make([]group.GroupBaseStruct, 0)

	yamlLines := strings.Split(string(data), "---")

	for _, line := range yamlLines {
		var header structs.Header
		if err := yaml.Unmarshal([]byte(line), &header); err != nil {
			return nil, err
		}

		if header.Type == structs.StructTypes.Group {
			var groupDefinitionStruct = group.GroupBaseStruct{Header: header}
			if err := yaml.Unmarshal([]byte(line), &groupDefinitionStruct); err != nil {
				return nil, err
			}

			rawGroup := groupDefinitionStruct
			if err := s.CompleteGroupInformation(&rawGroup); err != nil {
				return nil, err
			}

			groups = append(groups, rawGroup)
		}
	}

	return groups, nil
}

func (s GroupSerializer) GetGroupStructsFromFile(files ...string) ([]group.GroupBaseStruct, error) {
	groups := make([]group.GroupBaseStruct, 0)
	for _, file := range files {

		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		groupStructs, err := s.GetGroupStructsFromString(string(data))
		if err != nil {
			return nil, err
		}
		groups = append(groups, groupStructs...)

	}
	return groups, nil
}

func (s GroupSerializer) CompleteGroupInformation(group *group.GroupBaseStruct) error {

	return nil
}
