package basic_task_contract

import (
	"parsdevkit.net/basic/contracts"
	// task_payload "parsdevkit.net/modules/task/basic_task_payload"
	task_payload "parsdevkit.net/structs/task/common-task"
)

type TaskInterface contracts.TaskServiceInterface[task_payload.TaskBaseStruct]
