package update

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/resource/object_resource_contract"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
)

type UpdateResource struct{ flowx.BaseStep }

func (s *UpdateResource) Name() string { return "UpdateResource" }

func (s *UpdateResource) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

func (s *UpdateResource) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
