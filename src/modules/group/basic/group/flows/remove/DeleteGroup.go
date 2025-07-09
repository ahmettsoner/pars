package remove

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/group/basic_group_contract"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

type DeleteGroup struct{ flowx.BaseStep }

func (s *DeleteGroup) Name() string { return "DeleteGroup" }

func (s *DeleteGroup) Run(ctx context.Context, fc *flowx.FlowContext) error {

	permanent, ok := flowx.Get[bool](fc, "permanent")
	if !ok {
		panic(fmt.Errorf("xxx: permanent parametresi hatalı tipte"))
	}

	if permanent {
		service := ioc.Get[basic_group_contract.GroupInterface]()

		group, ok := flowx.Get[basic_group_payload_structs.GroupBaseStruct](fc, "group")
		if !ok {
			panic(fmt.Errorf("xxx: group parametresi hatalı tipte"))
		}

		logrus.Debugf("trying to create %v", group.Header.Name)
		if _, err := service.DeleteGroup(group); err != nil {
			return err
		}

	}
	return nil
}

func (s *DeleteGroup) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
