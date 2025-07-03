module parsdevkit.net/cmd

go 1.23.7

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.19.0
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/components v0.0.0-00010101000000-000000000000
	parsdevkit.net/models v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/group/basic_group v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/group/basic_group_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/project/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/resource/data v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/resource/object v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/task/common v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/code v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/file v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/shared v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/workspace/basic_workspace v0.0.0-00010101000000-000000000000
	parsdevkit.net/orchestrator v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/contexts v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/repositories v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/angular v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/common v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/dotnet v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/nodejs v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/pars v0.0.0-00010101000000-000000000000
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000
)

require (
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/glebarez/sqlite v1.10.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.25.7 // indirect
	modernc.org/libc v1.22.5 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.23.1 // indirect
	parsdevkit.net/context v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/engines v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/persistence/entities v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/platforms/go v0.0.0-00010101000000-000000000000 // indirect
)

replace parsdevkit.net/modules/workspace/basic_workspace => ../modules/workspace/basic/workspace.logic

replace parsdevkit.net/modules/group/basic_group => ../modules/group/basic/group.logic

replace parsdevkit.net/modules/group/basic_group_payload => ../modules/group/basic/group.payload

replace parsdevkit.net/modules/project/application => ../modules/project/application

replace parsdevkit.net/modules/resource/data => ../modules/resource/data

replace parsdevkit.net/modules/resource/object => ../modules/resource/object

replace parsdevkit.net/modules/template/code => ../modules/template/code

replace parsdevkit.net/modules/template/file => ../modules/template/file

replace parsdevkit.net/modules/template/shared => ../modules/template/shared

replace parsdevkit.net/modules/task/common => ../modules/task/common

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

replace parsdevkit.net/components => ../components

replace parsdevkit.net/platforms/angular => ../platforms/angular

replace parsdevkit.net/platforms/dotnet => ../platforms/dotnet

replace parsdevkit.net/platforms/go => ../platforms/go

replace parsdevkit.net/platforms/pars => ../platforms/pars

replace parsdevkit.net/structs => ../modules/structs

replace parsdevkit.net/models => ../models

replace parsdevkit.net/application => ../application

replace parsdevkit.net/orchestrator => ../orchestrator
