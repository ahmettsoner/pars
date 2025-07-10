package update

import (
	"context"
	"fmt"

	"parsdevkit.net/application/bus"
	"parsdevkit.net/internal/flowx"
	object_resource_payload_commands "parsdevkit.net/modules/resource/object_resource_payload/commands"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
)

type GenerateResourceContent struct{ flowx.BaseStep }

func (s *GenerateResourceContent) Name() string { return "GenerateResourceContent" }

func (s *GenerateResourceContent) Run(ctx context.Context, fc *flowx.FlowContext) error {

	resource, ok := flowx.Get[object_resource_payload_structs.ResourceBaseStruct](fc, "resource")
	if !ok {
		panic(fmt.Errorf("xxx: resource parametresi hatalı tipte"))
	}

	err := bus.SendCommand(object_resource_payload_commands.GenerateResourceContents{Data: resource})
	if err != nil {
		return err
	}

	return nil
}

func (s *GenerateResourceContent) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
