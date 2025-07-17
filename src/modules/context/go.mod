module parsdevkit.net/context

go 1.23.7

replace parsdevkit.net/components => ../../components

replace parsdevkit.net/application => ../../application

replace parsdevkit.net/internal => ../../internal

replace parsdevkit.net/models => ../../models

replace parsdevkit.net/pkg => ../../pkg

require (
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/components v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/internal v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/models v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
