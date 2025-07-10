module pars/tests

go 1.23.7

replace parsdevkit.net/modules/project/application_project => ../src/modules/project/application/project

replace parsdevkit.net/modules/project/application_project_contract => ../src/modules/project/application/project.contract

replace parsdevkit.net/modules/project/application_project_payload => ../src/modules/project/application/project.payload

replace parsdevkit.net/modules/tool/browse_tool => ../src/modules/tool/browse/tool

replace parsdevkit.net/modules/tool/browse_tool_contract => ../src/modules/tool/browse/tool.contract

replace parsdevkit.net/modules/tool/browse_tool_payload => ../src/modules/tool/browse/tool.payload

replace parsdevkit.net/modules/template/code_template => ../src/modules/template/code/template

replace parsdevkit.net/modules/template/code_template_payload => ../src/modules/template/code/template.payload

replace parsdevkit.net/modules/template/code_template_contract => ../src/modules/template/code/template.contract

replace parsdevkit.net/modules/template/shared_template => ../src/modules/template/shared/template

replace parsdevkit.net/modules/template/shared_template_payload => ../src/modules/template/shared/template.payload

replace parsdevkit.net/modules/template/shared_template_contract => ../src/modules/template/shared/template.contract

replace parsdevkit.net/modules/template/file_template => ../src/modules/template/file/template

replace parsdevkit.net/modules/template/file_template_payload => ../src/modules/template/file/template.payload

replace parsdevkit.net/modules/template/file_template_contract => ../src/modules/template/file/template.contract

replace parsdevkit.net/modules/task/basic_task => ../src/modules/task/basic/task

replace parsdevkit.net/modules/task/basic_task_contract => ../src/modules/task/basic/task.contract

replace parsdevkit.net/modules/task/basic_task_payload => ../src/modules/task/basic/task.payload

replace parsdevkit.net/modules/resource/data_resource => ../src/modules/resource/data/resource

replace parsdevkit.net/modules/resource/data_resource_contract => ../src/modules/resource/data/resource.contract

replace parsdevkit.net/modules/resource/data_resource_payload => ../src/modules/resource/data/resource.payload

replace parsdevkit.net/modules/group/basic_group => ../src/modules/group/basic/group

replace parsdevkit.net/modules/group/basic_group_contract => ../src/modules/group/basic/group.contract

replace parsdevkit.net/modules/group/basic_group_payload => ../src/modules/group/basic/group.payload

replace parsdevkit.net/modules/workspace/basic_workspace => ../src/modules/workspace/basic/workspace

replace parsdevkit.net/modules/workspace/basic_workspace_contract => ../src/modules/workspace/basic/workspace.contract

replace parsdevkit.net/modules/workspace/basic_workspace_payload => ../src/modules/workspace/basic/workspace.payload

replace parsdevkit.net/modules/environment/basic_environment => ../src/modules/environment/basic/environment

replace parsdevkit.net/modules/environment/basic_environment_contract => ../src/modules/environment/basic/environment.contract

replace parsdevkit.net/modules/environment/basic_environment_payload => ../src/modules/environment/basic/environment.payload

replace parsdevkit.net/modules/resource/object_resource => ../src/modules/resource/object/resource

replace parsdevkit.net/modules/resource/object_resource_contract => ../src/modules/resource/object/resource.contract

replace parsdevkit.net/modules/resource/object_resource_payload => ../src/modules/resource/object/resource.payload

replace parsdevkit.net/application => ../src/application

replace parsdevkit.net/internal => ../src/internal

replace parsdevkit.net/persistence/entities => ../src/data/entities

replace parsdevkit.net/persistence/repositories => ../src/data/repositories

replace parsdevkit.net/persistence/contexts => ../src/data/contexts

replace parsdevkit.net/platforms/core => ../src/platforms/core

replace parsdevkit.net/platforms/angular => ../src/platforms/angular

replace parsdevkit.net/platforms/nodejs => ../src/platforms/nodejs

replace parsdevkit.net/platforms/dotnet => ../src/platforms/dotnet

replace parsdevkit.net/platforms/go => ../src/platforms/go

replace parsdevkit.net/platforms/pars => ../src/platforms/pars

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
	parsdevkit.net/modules/group/basic_group_contract v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/group/basic_group_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/project/application_project v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/project/application_project_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/resource/object_resource v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/resource/object_resource_contract v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/resource/object_resource_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/code_template v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/code_template_contract v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/code_template_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/workspace/basic_workspace v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/workspace/basic_workspace_contract v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/workspace/basic_workspace_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/environment/basic_environment v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/environment/basic_environment_contract v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/environment/basic_environment_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/entities v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/repositories v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/angular v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/core v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/dotnet v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/go v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/pars v0.0.0-00010101000000-000000000000
)

require (
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/spf13/viper v1.19.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	parsdevkit.net/components v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/internal v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/project/application_project_contract v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/resource/data_resource v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/resource/data_resource_contract v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/resource/data_resource_payload v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/task/basic_task v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/task/basic_task_contract v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/task/basic_task_payload v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/file_template v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/file_template_contract v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/file_template_payload v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/shared_template v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/shared_template_contract v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/template/shared_template_payload v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/tool/browse_tool v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/modules/tool/browse_tool_contract v0.0.0-00010101000000-000000000000 // indirect
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
	parsdevkit.net/modules/tool/browse_tool_payload v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/persistence/contexts v0.0.0-00010101000000-000000000000
	parsdevkit.net/platforms/nodejs v0.0.0-00010101000000-000000000000
	parsdevkit.net/providers v0.0.0-00010101000000-000000000000 // indirect
)

replace parsdevkit.net/pkg => ../src/pkg
