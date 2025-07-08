module parsdevkit.net/components

go 1.23.7

replace parsdevkit.net/application => ../application

replace parsdevkit.net/pkg => ../pkg

replace parsdevkit.net/platforms/core => ../platforms/core

replace parsdevkit.net/platforms/angular => ../platforms/angular

replace parsdevkit.net/platforms/nodejs => ../platforms/nodejs

replace parsdevkit.net/platforms/dotnet => ../platforms/dotnet

replace parsdevkit.net/platforms/go => ../platforms/go

replace parsdevkit.net/platforms/pars => ../platforms/pars

replace parsdevkit.net/models => ../models

replace parsdevkit.net/providers => ../providers

replace parsdevkit.net/modules/workspace/basic_workspace_payload => ../modules/workspace/basic/workspace.payload


require (
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000
)

require (
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	parsdevkit.net/modules/workspace/basic_workspace_payload v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/models v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/angular v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/dotnet v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/go v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/nodejs v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/pars v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000 // indirect
)
