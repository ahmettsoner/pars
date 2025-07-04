module parsdevkit.net/modules/task/basic_task_contract

replace parsdevkit.net/modules/task/basic_task_payload => ../task.payload

replace parsdevkit.net/structs => ../../../structs

replace parsdevkit.net/basic => ../../../../basic

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	parsdevkit.net/basic v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/task/basic_task_payload v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
