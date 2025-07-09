package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/task/basic_task_contract"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"
)

type SaveTask struct{ flowx.BaseStep }

func (s *SaveTask) Name() string { return "SaveTask" }

func (s *SaveTask) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[basic_task_contract.TaskInterface]()

	task, ok := flowx.Get[basic_task_payload_structs.TaskBaseStruct](fc, "task")
	if !ok {
		panic(fmt.Errorf("xxx: task parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", task.Header.Name)
	if _, err := service.SaveTask(task); err != nil {
		return err
	}

	return nil
}

func (s *SaveTask) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[basic_task_contract.TaskInterface]()

	task, ok := flowx.Get[basic_task_payload_structs.TaskBaseStruct](fc, "task")
	if !ok {
		panic(fmt.Errorf("xxx: task parametresi hatalı tipte"))
	}

	logrus.Warnf("yyy: rolling back the creating %v!", task.Header.Name)
	if _, err := service.UndoSaveTask(task); err != nil {
		return fmt.Errorf("xxx: Code Task Information kayıt sırasında meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", &task.Header.Name, &err)
	}

	fc.Log("Rollback: Persist task %s", task.Header.Name)
	return nil
}
