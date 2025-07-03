package basic_task_contract

import (
	"parsdevkit.net/basic/contracts"
	task_payload "parsdevkit.net/modules/task/basic_task_payload"
)

type TaskInterface contracts.TaskServiceInterface[task_payload.TaskBaseStruct]
