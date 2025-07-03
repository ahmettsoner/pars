module parsdevkit.net/engines

go 1.23.7

replace parsdevkit.net/persistence/entities => ../../data/entities

replace parsdevkit.net/persistence/repositories => ../../data/repositories

replace parsdevkit.net/platforms/core => ../../platforms/core

replace parsdevkit.net/platforms/common => ../../platforms/common

replace parsdevkit.net/platforms/angular => ../../platforms/angular

replace parsdevkit.net/platforms/nodejs => ../../platforms/nodejs

replace parsdevkit.net/platforms/dotnet => ../../platforms/dotnet

replace parsdevkit.net/platforms/go => ../../platforms/go

replace parsdevkit.net/platforms/pars => ../../platforms/pars

replace parsdevkit.net/structs => ../structs

replace parsdevkit.net/modules/project/application => ../project/application

replace parsdevkit.net/modules/template/code => ../template/code

replace parsdevkit.net/modules/template/shared => ../template/shared

replace parsdevkit.net/modules/template/file => ../template/file

replace parsdevkit.net/modules/task/common => ../task/common

replace parsdevkit.net/modules/resource/data => ../resource/data

replace parsdevkit.net/modules/group/group => ../group/group

replace parsdevkit.net/modules/workspace/workspace => ../workspace/workspace

replace parsdevkit.net/modules/resource/object => ../resource/object

replace parsdevkit.net/providers => ../../providers

replace parsdevkit.net/pkg => ../../pkg

replace parsdevkit.net/persistence/contexts => ../../data/contexts

replace parsdevkit.net/context => ../context

replace parsdevkit.net/application => ../../application

replace parsdevkit.net/components => ../../components

replace parsdevkit.net/models => ../../models

require (	
	github.com/sirupsen/logrus v1.9.3
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/components v0.0.0-00010101000000-000000000000
	parsdevkit.net/context v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/contexts v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/entities v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/repositories v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000
)

require (
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/glebarez/sqlite v1.10.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.25.7 // indirect
	modernc.org/libc v1.22.5 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.23.1 // indirect
	parsdevkit.net/models v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/angular v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/common v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/dotnet v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/go v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/nodejs v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/pars v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000 // indirect
)
