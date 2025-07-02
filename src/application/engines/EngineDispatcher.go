package engines

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"
)

func DispatchEngineProcess(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			fmt.Printf("Processing '%s' Schemas\n", key)
			err := engineModule.Process(ctx, data)
			if err != nil {
				return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
			}
		}
	}
	return nil
}
func DispatchEngineDestroy(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSortedReverse() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			fmt.Printf("Destroying '%s' Schemas\n", key)
			err := engineModule.Destroy(ctx, data)
			if err != nil {
				return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
			}
		}
	}
	return nil
}

func GroupSchemas(t []schemas.SchemaInterface) map[string][]schemas.SchemaInterface {

	schemaGroups := make(map[string][]schemas.SchemaInterface, 0)
	for _, data := range t {
		header := data.GetHeader()

		schemaGroups[header.GetKey()] = append(schemaGroups[header.GetKey()], data)
	}

	return schemaGroups
}
