package update

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/task/basic_task_contract"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"
)

type UpdateTask struct{ flowx.BaseStep }

func (s *UpdateTask) Name() string { return "UpdateTask" }

func (s *UpdateTask) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

func (s *UpdateTask) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
