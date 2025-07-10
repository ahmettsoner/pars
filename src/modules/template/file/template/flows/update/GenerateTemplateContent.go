package update

import (
	"context"
	"fmt"

	"parsdevkit.net/application/bus"
	"parsdevkit.net/internal/flowx"
	file_template_payload_commands "parsdevkit.net/modules/template/file_template_payload/commands"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
)

type GenerateTemplateContent struct{ flowx.BaseStep }

func (s *GenerateTemplateContent) Name() string { return "GenerateTemplateContent" }

func (s *GenerateTemplateContent) Run(ctx context.Context, fc *flowx.FlowContext) error {

	template, ok := flowx.Get[file_template_payload_structs.TemplateBaseStruct](fc, "template")
	if !ok {
		panic(fmt.Errorf("xxx: template parametresi hatalı tipte"))
	}

	err := bus.SendCommand(file_template_payload_commands.GenerateTemplateContents{Data: template})
	if err != nil {
		return err
	}

	return nil
}

func (s *GenerateTemplateContent) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
