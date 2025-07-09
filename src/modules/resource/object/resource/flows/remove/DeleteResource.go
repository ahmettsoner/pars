package remove

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/resource/object_resource_contract"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
)

type DeleteResource struct{ flowx.BaseStep }

func (s *DeleteResource) Name() string { return "DeleteResource" }

func (s *DeleteResource) Run(ctx context.Context, fc *flowx.FlowContext) error {

	permanent, ok := flowx.Get[bool](fc, "permanent")
	if !ok {
		panic(fmt.Errorf("xxx: permanent parametresi hatalı tipte"))
	}

	if permanent {
		service := ioc.Get[object_resource_contract.ResourceInterface]()

		resource, ok := flowx.Get[object_resource_payload_structs.ResourceBaseStruct](fc, "resource")
		if !ok {
			panic(fmt.Errorf("xxx: resource parametresi hatalı tipte"))
		}

		logrus.Debugf("trying to create %v", resource.Header.Name)
		if _, err := service.DeleteResource(resource); err != nil {
			return err
		}

	}
	return nil
}

func (s *DeleteResource) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
