package browse_tool

import (
	"fmt"

	"github.com/sirupsen/logrus"

	browse_tool_payload_structs "parsdevkit.net/modules/tool/browse_tool_payload/structs"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/internal/flowx"
	browse_steps "parsdevkit.net/modules/tool/browse_tool/flows/browse"
)

type ToolEngine struct{}

func (s ToolEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*browse_tool_payload_structs.ToolBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s ToolEngine) Browse(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.browse(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}
func (s ToolEngine) prepareToBrowse(ctx *application.ApplicationContext, tools []browse_tool_payload_structs.ToolBaseStruct) ([]browse_tool_payload_structs.ToolBaseStruct, error) {

	readyToCreateStructs := make([]browse_tool_payload_structs.ToolBaseStruct, 0)

	for _, tool := range tools {
		if err := s.completeInformation(ctx, &tool); err != nil {
			return nil, err
		}
		readyToCreateStructs = append(readyToCreateStructs, tool)
	}
	logrus.Debugf("'%d' tool(s) detected that will create", len(readyToCreateStructs))

	return readyToCreateStructs, nil
}
func (s ToolEngine) browse(ctx *application.ApplicationContext, tools []browse_tool_payload_structs.ToolBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToBrowse(ctx, tools)
	if err != nil {
		return err
	}

	for _, tool := range readyToCreateStructs {

		fmt.Printf("\n🛠️  Creating: %s.%s\n\n", tool.Header.Name, tool.GetKey())

		toolFlow := flowx.NewFlow("CreateNewTool").
			Step(&browse_steps.OpenBrowser{})

		fc := flowx.NewContextWithData(map[string]any{
			"tool": tool,
		})

		if err := toolFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Tool Create işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s ToolEngine) completeInformation(ctx *application.ApplicationContext, model *browse_tool_payload_structs.ToolBaseStruct) error {

	return nil
}

func (s ToolEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  browse_tool_payload_structs.MODULE_KEY,
		Order: 1000,
	}
}
func CastArrayToConcrate(data []schemas.SchemaInterface) ([]browse_tool_payload_structs.ToolBaseStruct, error) {
	r := make([]browse_tool_payload_structs.ToolBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*browse_tool_payload_structs.ToolBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected browse_tool_payload_structs.ToolBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}
