package update

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
)

type UpdateTemplate struct{ flowx.BaseStep }

func (s *UpdateTemplate) Name() string { return "UpdateTemplate" }

func (s *UpdateTemplate) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[code_template_contract.TemplateInterface]()

	template, ok := flowx.Get[code_template_payload_structs.TemplateBaseStruct](fc, "template")
	if !ok {
		panic(fmt.Errorf("xxx: template parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", template.Header.Name)
	if _, err := service.SaveTemplate(template); err != nil {
		return err
	}

	return nil
}

func (s *UpdateTemplate) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
