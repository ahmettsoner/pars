module parsdevkit.net/modules/project/application_project_payload

replace parsdevkit.net/structs => ../../../../structs

replace parsdevkit.net/application => ../../../../application

replace parsdevkit.net/models => ../../../../models

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v3 v3.0.1
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/models v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
)

require golang.org/x/sys v0.26.0 // indirect
