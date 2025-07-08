package basic_task_contract

import (
	"parsdevkit.net/application/contracts"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"
)

type TaskInterface contracts.TaskServiceInterface[basic_task_payload_structs.TaskBaseStruct]
