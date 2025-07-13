package engines

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/pkg/utilities/array"
)

func DispatchEngineProcess(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			if engine, ok := engineModule.(ProcessorInterface); ok {
				err := engine.Process(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
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
			if engine, ok := engineModule.(DestroyerInterface); ok {
				err := engine.Destroy(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}

func DispatchEngineClean(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			if engine, ok := engineModule.(ApplicationProjectCleanerEngineInterface); ok {
				err := engine.Clean(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}
func DispatchEngineInstall(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			if engine, ok := engineModule.(ApplicationProjectInstallerEngineInterface); ok {
				err := engine.Install(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}
func DispatchEngineExecute(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			if engine, ok := engineModule.(ApplicationProjectExecuterEngineInterface); ok {
				err := engine.Execute(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}
func DispatchEngineOpen(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			if engine, ok := engineModule.(ApplicationProjectOpenerEngineInterface); ok {
				err := engine.Open(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}
func DispatchEngineTest(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			if engine, ok := engineModule.(ApplicationProjectTesterEngineInterface); ok {
				err := engine.Test(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}
func DispatchEngineList(ctx *application.ApplicationContext, t []string) error {

	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if ok := array.ContainsSlice(t, key); ok {
			if engine, ok := engineModule.(ListEngineInterface); ok {
				err := engine.List(ctx)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}
func DispatchEngineDescribe(ctx *application.ApplicationContext, t []string, args ...any) error {

	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if ok := array.ContainsSlice(t, key); ok {
			if engine, ok := engineModule.(DescribeEngineInterface); ok {
				err := engine.Describe(ctx, args...)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}
func DispatchEngineRemove(ctx *application.ApplicationContext, t []string, args ...any) error {

	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if ok := array.ContainsSlice(t, key); ok {
			if engine, ok := engineModule.(RemoveEngineInterface); ok {
				err := engine.Remove(ctx, args...)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
			}
		}
	}
	return nil
}

func DispatchEngineBrowse(ctx *application.ApplicationContext, t []schemas.SchemaInterface) error {

	schemaGroups := GroupSchemas(t)
	for _, engineModule := range AllSorted() {
		key := engineModule.GetConfig().Name
		if data, ok := schemaGroups[key]; ok {
			if engine, ok := engineModule.(ToolsEngineInterface); ok {
				err := engine.Browse(ctx, data)
				if err != nil {
					return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
				}
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
