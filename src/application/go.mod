module parsdevkit.net/application

go 1.23.7

replace parsdevkit.net/engines => ../modules/engines

replace parsdevkit.net/core => ../core

replace parsdevkit.net/context => ../modules/context


replace parsdevkit.net/persistence/entities => ../data/entities

replace parsdevkit.net/persistence/repositories => ../data/repositories

replace parsdevkit.net/persistence/contexts => ../data/contexts

replace parsdevkit.net/templates => ../core/templates

replace parsdevkit.net/core/utils => ../modules/utils

replace parsdevkit.net/providers => ../components/providers

replace parsdevkit.net/platforms/core => ../platforms/core

replace parsdevkit.net/platforms/common => ../platforms/common

replace parsdevkit.net/platforms/nodejs => ../platforms/nodejs

replace parsdevkit.net/platforms/angular => ../platforms/angular

replace parsdevkit.net/platforms/dotnet => ../platforms/dotnet

replace parsdevkit.net/platforms/go => ../platforms/go

replace parsdevkit.net/platforms/pars => ../platforms/pars

replace parsdevkit.net/structs => ../modules/structs

replace parsdevkit.net/models => ../modules/models

replace parsdevkit.net/operation/services => ../modules/services

require (
	parsdevkit.net/core v0.0.0-00010101000000-000000000000
	parsdevkit.net/core/utils v0.0.0-00010101000000-000000000000
	parsdevkit.net/operation/services v0.0.0-00010101000000-000000000000
)

require (
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/glebarez/sqlite v1.10.0 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.25.7 // indirect
	modernc.org/libc v1.22.5 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.23.1 // indirect
	parsdevkit.net/models v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/persistence/contexts v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/persistence/entities v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/persistence/repositories v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/angular v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/common v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/dotnet v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/go v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/nodejs v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/pars v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000 // indirect
)
