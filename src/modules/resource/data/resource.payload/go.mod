module parsdevkit.net/modules/data/resource_data_payload

replace parsdevkit.net/structs => ../../../../structs

replace parsdevkit.net/resource => ../../../../resource

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	gopkg.in/yaml.v3 v3.0.1
	parsdevkit.net/resource v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
)
