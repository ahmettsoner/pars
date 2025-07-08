module parsdevkit.net/platforms/core

go 1.23.7

replace parsdevkit.net/pkg => ../../pkg

replace parsdevkit.net/models => ../../models

replace parsdevkit.net/application => ../../application

require (
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
)

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/models v0.0.0-00010101000000-000000000000 // indirect
)
