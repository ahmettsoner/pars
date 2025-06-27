module parsdevkit.net/providers

go 1.22

replace parsdevkit.net/models => ../../modules/models

replace parsdevkit.net/core => ../../../core

replace parsdevkit.net/core/utils => ../../modules/utils

require github.com/sirupsen/logrus v1.9.3

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
