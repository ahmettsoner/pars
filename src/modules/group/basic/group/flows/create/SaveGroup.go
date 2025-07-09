package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/group/basic_group_contract"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

type SaveGroup struct{ flowx.BaseStep }

func (s *SaveGroup) Name() string { return "SaveGroup" }

func (s *SaveGroup) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

func (s *SaveGroup) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[basic_group_contract.GroupInterface]()

	group, ok := flowx.Get[basic_group_payload_structs.GroupBaseStruct](fc, "group")
	if !ok {
		panic(fmt.Errorf("xxx: group parametresi hatalı tipte"))
	}

	logrus.Warnf("yyy: rolling back the creating %v!", group.Header.Name)
	if _, err := service.UndoSaveGroup(group); err != nil {
		return fmt.Errorf("xxx: Code Group Information kayıt sırasında meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", &group.Header.Name, &err)
	}

	fc.Log("Rollback: Persist group %s", group.Header.Name)
	return nil
}
