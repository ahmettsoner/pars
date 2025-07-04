module parsdevkit.net/modules/project/application_project_contract

replace parsdevkit.net/modules/project/application_project_payload => ../project.payload

replace parsdevkit.net/structs => ../../../../structs

replace parsdevkit.net/application => ../../../../application

replace parsdevkit.net/pkg => ../../../../pkg

replace parsdevkit.net/models => ../../../../models

go 1.23.7

require (
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/project/application_project_payload v0.0.0-00010101000000-000000000000
)

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/models v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
