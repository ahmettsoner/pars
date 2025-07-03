module pars/tests

go 1.23.7

replace parsdevkit.net/modules/project/application_project => ../src/modules/project/application/project

replace parsdevkit.net/modules/template/code => ../src/modules/template/code

replace parsdevkit.net/modules/template/shared => ../src/modules/template/shared

replace parsdevkit.net/modules/template/file => ../src/modules/template/file

replace parsdevkit.net/modules/task/common => ../src/modules/task/common

replace parsdevkit.net/modules/resource/data => ../src/modules/resource/data

replace parsdevkit.net/modules/group/basic_group => ../src/modules/group/basic/group

replace parsdevkit.net/modules/group/basic_group_payload => ../src/modules/group/basic/group.payload

replace parsdevkit.net/modules/workspace/basic_workspace => ../src/modules/workspace/basic/workspace

replace parsdevkit.net/modules/resource/object => ../src/modules/resource/object

replace parsdevkit.net/application => ../src/application

replace parsdevkit.net/persistence/entities => ../src/data/entities

replace parsdevkit.net/persistence/repositories => ../src/data/repositories

replace parsdevkit.net/persistence/contexts => ../src/data/contexts

replace parsdevkit.net/platforms/core => ../src/platforms/core

replace parsdevkit.net/platforms/common => ../src/platforms/common

replace parsdevkit.net/platforms/angular => ../src/platforms/angular

replace parsdevkit.net/platforms/nodejs => ../src/platforms/nodejs

replace parsdevkit.net/platforms/dotnet => ../src/platforms/dotnet

replace parsdevkit.net/platforms/go => ../src/platforms/go

replace parsdevkit.net/platforms/pars => ../src/platforms/pars

replace parsdevkit.net/structs => ../src/modules/structs

replace parsdevkit.net/context => ../src/modules/context

replace parsdevkit.net/providers => ../src/providers

replace parsdevkit.net/models => ../src/models

replace parsdevkit.net/engines => ../src/modules/engines

replace parsdevkit.net/cmd => ../src/cmd

replace parsdevkit.net/components => ../src/components

require (
	github.com/magiconair/properties v1.8.7
	github.com/stretchr/testify v1.9.0
	gopkg.in/yaml.v3 v3.0.1
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/models v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/group/basic_group v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/group/basic_group_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/project/application_project v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/resource/object v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/code v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/workspace/basic_workspace v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/entities v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/repositories v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/angular v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/common v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/dotnet v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/go v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/pars v0.0.0-00010101000000-000000000000
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000
)

require (
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/spf13/viper v1.19.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	parsdevkit.net/components v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/resource/data v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/task/common v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/file v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/shared v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/orchestrator v0.0.0-00010101000000-000000000000 // indirect
)

require (
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/glebarez/go-sqlite v1.21.2 // indirect
	github.com/glebarez/sqlite v1.10.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/sirupsen/logrus v1.9.3
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gorm.io/gorm v1.25.7 // indirect
	modernc.org/libc v1.22.5 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.23.1 // indirect
	parsdevkit.net/cmd v0.0.0-00010101000000-000000000000
	parsdevkit.net/context v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/engines v0.0.0-00010101000000-000000000000 // indirect; indirec
	parsdevkit.net/persistence/contexts v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/nodejs v0.0.0-00010101000000-000000000000
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000 // indirect
)

replace parsdevkit.net/pkg => ../src/pkg

replace parsdevkit.net/orchestrator => ../src/orchestrator
