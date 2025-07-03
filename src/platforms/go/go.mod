module parsdevkit.net/platforms/go

go 1.22

replace parsdevkit.net/platforms/core => ../core

replace parsdevkit.net/pkg => ../../pkg

replace parsdevkit.net/structs => ../../modules/structs

replace parsdevkit.net/models => ../../models

replace parsdevkit.net/providers => ../../providers

require (
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/models v0.0.0-00010101000000-000000000000 // indirect
)
