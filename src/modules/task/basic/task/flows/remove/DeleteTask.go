package remove

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/task/basic_task_contract"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"
)

type DeleteTask struct{ flowx.BaseStep }

func (s *DeleteTask) Name() string { return "DeleteTask" }

func (s *DeleteTask) Run(ctx context.Context, fc *flowx.FlowContext) error {

	permanent, ok := flowx.Get[bool](fc, "permanent")
	if !ok {
		panic(fmt.Errorf("xxx: permanent parametresi hatalı tipte"))
	}

	if permanent {
		service := ioc.Get[basic_task_contract.TaskInterface]()

		task, ok := flowx.Get[basic_task_payload_structs.TaskBaseStruct](fc, "task")
		if !ok {
			panic(fmt.Errorf("xxx: task parametresi hatalı tipte"))
		}

		logrus.Debugf("trying to create %v", task.Header.Name)
		if _, err := service.DeleteTask(task); err != nil {
			return err
		}

	}
	return nil
}

func (s *DeleteTask) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
