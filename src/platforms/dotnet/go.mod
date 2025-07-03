module parsdevkit.net/platforms/dotnet

go 1.23.7

replace parsdevkit.net/platforms/core => ../core

replace parsdevkit.net/pkg => ../../pkg

replace parsdevkit.net/structs => ../../modules/structs

replace parsdevkit.net/models => ../../models

replace parsdevkit.net/providers => ../../providers

replace parsdevkit.net/application => ../../application

require (
	github.com/clbanning/mxj/v2 v2.7.0
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/models v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
