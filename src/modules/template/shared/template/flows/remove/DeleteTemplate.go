package remove

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
)

type DeleteTemplate struct{ flowx.BaseStep }

func (s *DeleteTemplate) Name() string { return "DeleteTemplate" }

func (s *DeleteTemplate) Run(ctx context.Context, fc *flowx.FlowContext) error {

	permanent, ok := flowx.Get[bool](fc, "permanent")
	if !ok {
		panic(fmt.Errorf("xxx: permanent parametresi hatalı tipte"))
	}

	if permanent {
		service := ioc.Get[code_template_contract.TemplateInterface]()

		template, ok := flowx.Get[code_template_payload_structs.TemplateBaseStruct](fc, "template")
		if !ok {
			panic(fmt.Errorf("xxx: template parametresi hatalı tipte"))
		}

		logrus.Debugf("trying to create %v", template.Header.Name)
		if _, err := service.DeleteTemplate(template); err != nil {
			return err
		}

	}
	return nil
}

func (s *DeleteTemplate) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
