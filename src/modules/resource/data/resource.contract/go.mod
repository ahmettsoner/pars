module parsdevkit.net/modules/resource/data_resource_contract

replace parsdevkit.net/modules/resource/data_resource_payload => ../resource.payload

replace parsdevkit.net/data => ../../../../data

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	parsdevkit.net/data v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/resource/data_resource_payload v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
