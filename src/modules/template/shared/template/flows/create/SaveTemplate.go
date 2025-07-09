package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
)

type SaveTemplate struct{ flowx.BaseStep }

func (s *SaveTemplate) Name() string { return "SaveTemplate" }

func (s *SaveTemplate) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

func (s *SaveTemplate) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[code_template_contract.TemplateInterface]()

	template, ok := flowx.Get[code_template_payload_structs.TemplateBaseStruct](fc, "template")
	if !ok {
		panic(fmt.Errorf("xxx: template parametresi hatalı tipte"))
	}

	logrus.Warnf("yyy: rolling back the creating %v!", template.Header.Name)
	if _, err := service.UndoSaveTemplate(template); err != nil {
		return fmt.Errorf("xxx: Code Template Information kayıt sırasında meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", &template.Header.Name, &err)
	}

	fc.Log("Rollback: Persist template %s", template.Header.Name)
	return nil
}
