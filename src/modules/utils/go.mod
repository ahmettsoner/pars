module parsdevkit.net/core/utils

go 1.22

require gopkg.in/yaml.v3 v3.0.1

require golang.org/x/sys v0.15.0 // indirect

require (
	github.com/sirupsen/logrus v1.9.3
	parsdevkit.net/core v0.0.0-00010101000000-000000000000 // indirect
)

replace parsdevkit.net/core => ../../core
