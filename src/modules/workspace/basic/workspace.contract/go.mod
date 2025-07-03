module parsdevkit.net/modules/workspace/basic_workspace_contract

replace parsdevkit.net/modules/workspace/basic_workspace_payload => ../workspace.payload

replace parsdevkit.net/structs => ../../../../structs

replace parsdevkit.net/application => ../../../../application

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/workspace/basic_workspace_payload v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
