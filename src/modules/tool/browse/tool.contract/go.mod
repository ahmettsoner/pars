module parsdevkit.net/modules/tool/browse_tool_contract

replace parsdevkit.net/modules/tool/browse_tool_payload => ../tool.payload

replace parsdevkit.net/application => ../../../../application

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/tool/browse_tool_payload v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
