package v2

import (
	"fmt"

	"parsdevkit.net/core"
	"parsdevkit.net/core/schemas"
	"parsdevkit.net/core/utils"
	group "parsdevkit.net/modules/group/group"
	projectApplication "parsdevkit.net/modules/project/application"
	resourceData "parsdevkit.net/modules/resource/data"
	resourceObject "parsdevkit.net/modules/resource/object"
	taskCommon "parsdevkit.net/modules/task/common"
	templateCode "parsdevkit.net/modules/template/code"
	templateFile "parsdevkit.net/modules/template/file"
	templateShared "parsdevkit.net/modules/template/shared"
)

var engineRegistry = map[string]Engine{
	"Group":               group.GroupEngine{},
	"Project.Application": projectApplication.ApplicationProjectEngine{},
	"Resource.Data":       resourceData.DataResourceEngine{},
	"Resource.Object":     resourceObject.ObjectResourceEngine{},
	"Template.Code":       templateCode.CodeTemplateEngine{},
	"Template.File":       templateFile.FileTemplateEngine{},
	"Template.Shared":     templateShared.SharedTemplateEngine{},
	"Task.Common":         taskCommon.CommonTaskEngine{},
}
var orderedKeys = []string{
	"Group",
	"Project.Application",
	"Resource.Data",
	"Resource.Object",
	"Template.File",
	"Template.Code",
	"Template.Shared",
	"Task.Common",
}

func DispatchEngineProcess(ctx *core.Context, t []schemas.Schema) error {

	schemaGroups := make(map[string][]schemas.Schema, 0)
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
			err := engine.Process(ctx, data)
			if err != nil {
				return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
			}
		}
	}
	return nil
}
func DispatchEngineDestroy(ctx *core.Context, t []schemas.Schema) error {

	schemaGroups := make(map[string][]schemas.Schema, 0)
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
			err := engine.Destroy(ctx, data)
			if err != nil {
				return fmt.Errorf("xxx %s(%s) işlemi sırasında engine hata verdi\n %w", key, key, err)
			}
		}
	}
	return nil
}
