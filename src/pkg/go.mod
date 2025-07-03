module parsdevkit.net/pkg

replace parsdevkit.net/cmd => ../cmd

replace parsdevkit.net/engines => ../modules/engines

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

go 1.23.7

require gopkg.in/yaml.v3 v3.0.1

require (
	github.com/kr/pretty v0.3.1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)
