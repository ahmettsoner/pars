module parsdevkit.net/structs

go 1.22

replace parsdevkit.net/models => ../models

replace parsdevkit.net/core => ../../core

replace parsdevkit.net/core/utils => ../utils

require (
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v3 v3.0.1
	parsdevkit.net/core v0.0.0-00010101000000-000000000000
	parsdevkit.net/models v0.0.0-00010101000000-000000000000
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
