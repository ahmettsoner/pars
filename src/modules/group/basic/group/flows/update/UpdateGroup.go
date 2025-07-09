package update

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/group/basic_group_contract"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

type UpdateGroup struct{ flowx.BaseStep }

func (s *UpdateGroup) Name() string { return "UpdateGroup" }

func (s *UpdateGroup) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[basic_group_contract.GroupInterface]()

	group, ok := flowx.Get[basic_group_payload_structs.GroupBaseStruct](fc, "group")
	if !ok {
		panic(fmt.Errorf("xxx: group parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", group.Header.Name)
	if _, err := service.SaveGroup(group); err != nil {
		return err
	}

	return nil
}

func (s *UpdateGroup) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
