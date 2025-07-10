module parsdevkit.net/modules/environment/basic_environment_contract

replace parsdevkit.net/modules/environment/basic_environment_payload => ../environment.payload

replace parsdevkit.net/application => ../../../../application

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/environment/basic_environment_payload v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
