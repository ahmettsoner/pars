package v2

import (
	"fmt"

	"parsdevkit.net/core/utils"
	engines "parsdevkit.net/engines/v2/engines"
	"parsdevkit.net/structs"
)

var engineRegistry = map[string]Engine{
	"Group":               engines.GroupEngine{},
	"Project.Application": engines.ApplicationProjectEngine{},
	"Resource.Data":       engines.DataResourceEngine{},
	"Resource.Object":     engines.ObjectResourceEngine{},
	"Template.Code":       engines.CodeTemplateEngine{},
	"Template.File":       engines.FileTemplateEngine{},
	"Template.Shared":     engines.SharedTemplateEngine{},
}
var orderedKeys = []string{
	"Group",
	"Project.Application",
	"Resource.Data",
	"Resource.Object",
	"Template.File",
	"Template.Code",
	"Template.Shared",
}

func DispatchEngineProcess(t []structs.Schema) error {

	schemaGroups := make(map[string][]structs.Schema, 0)
	for _, data := range t {
		header := data.GetHeader()

		var key string
		if header.Kind != "" {
			key = fmt.Sprintf("%s.%s", header.Type, header.Kind)
		} else {
			key = string(header.Type)
		}

		schemaGroups[key] = append(schemaGroups[key], data)
	}

	for _, key := range orderedKeys {
		if data, ok := schemaGroups[key]; ok {
			engine, ok := engineRegistry[key]
			if !ok {
				return fmt.Errorf("no engine found for typeKind %s", key)
			}

			fmt.Printf("Processing '%s' Schemas\n", key)
			err := engine.Process(data)
			if err != nil {
				return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
			}
		}
	}
	return nil
}
func DispatchEngineDestroy(t []structs.Schema) error {

	schemaGroups := make(map[string][]structs.Schema, 0)
	for _, data := range t {
		header := data.GetHeader()

		var key string
		if header.Kind != "" {
			key = fmt.Sprintf("%s.%s", header.Type, header.Kind)
		} else {
			key = string(header.Type)
		}

		schemaGroups[key] = append(schemaGroups[key], data)
	}

	for _, key := range utils.Reverse(orderedKeys) {
		if data, ok := schemaGroups[key]; ok {
			engine, ok := engineRegistry[key]
			if !ok {
				return fmt.Errorf("no engine found for typeKind %s", key)
			}

			fmt.Printf("Destroying '%s' Schemas\n", key)
			err := engine.Destroy(data)
			if err != nil {
				return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
			}
		}
	}
	return nil
}
