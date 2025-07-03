module parsdevkit.net/application

go 1.23.7

replace parsdevkit.net/engines => ../modules/engines

replace parsdevkit.net/pkg => ../pkg

replace parsdevkit.net/context => ../modules/context

replace parsdevkit.net/persistence/entities => ../data/entities

replace parsdevkit.net/persistence/repositories => ../data/repositories

replace parsdevkit.net/persistence/contexts => ../data/contexts

replace parsdevkit.net/providers => ../providers

replace parsdevkit.net/platforms/core => ../platforms/core

replace parsdevkit.net/platforms/common => ../platforms/common

replace parsdevkit.net/platforms/nodejs => ../platforms/nodejs

replace parsdevkit.net/platforms/angular => ../platforms/angular

replace parsdevkit.net/platforms/dotnet => ../platforms/dotnet

replace parsdevkit.net/platforms/go => ../platforms/go

replace parsdevkit.net/platforms/pars => ../platforms/pars

replace parsdevkit.net/structs => ../modules/structs

replace parsdevkit.net/models => ../models

require (
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v3 v3.0.1
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000
)

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/sys v0.26.0 // indirect
)
