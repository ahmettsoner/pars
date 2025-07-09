package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/resource/object_resource_contract"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
)

type SaveResource struct{ flowx.BaseStep }

func (s *SaveResource) Name() string { return "SaveResource" }

func (s *SaveResource) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[object_resource_contract.ResourceInterface]()

	resource, ok := flowx.Get[object_resource_payload_structs.ResourceBaseStruct](fc, "resource")
	if !ok {
		panic(fmt.Errorf("xxx: resource parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", resource.Header.Name)
	if _, err := service.SaveResource(resource); err != nil {
		return err
	}

	return nil
}

func (s *SaveResource) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[object_resource_contract.ResourceInterface]()

	resource, ok := flowx.Get[object_resource_payload_structs.ResourceBaseStruct](fc, "resource")
	if !ok {
		panic(fmt.Errorf("xxx: resource parametresi hatalı tipte"))
	}

	logrus.Warnf("yyy: rolling back the creating %v!", resource.Header.Name)
	if _, err := service.UndoSaveResource(resource); err != nil {
		return fmt.Errorf("xxx: Object Resource Information kayıt sırasında meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", &resource.Header.Name, &err)
	}

	fc.Log("Rollback: Persist resource %s", resource.Header.Name)
	return nil
}
