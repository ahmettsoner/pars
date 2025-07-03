module parsdevkit.net/modules/group/basic_group

replace parsdevkit.net/structs => ../../../structs

replace parsdevkit.net/application => ../../../../application

replace parsdevkit.net/pkg => ../../../../pkg

replace parsdevkit.net/persistence/contexts => ../../../../data/contexts

replace parsdevkit.net/persistence/entities => ../../../../data/entities

replace parsdevkit.net/persistence/repositories => ../../../../data/repositories

replace parsdevkit.net/modules/group/basic_group_contract => ../group.contract

replace parsdevkit.net/modules/group/basic_group_payload => ../group.payload

go 1.23.7

require (
	github.com/sirupsen/logrus v1.9.3
	parsdevkit.net/application v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/group/basic_group_payload v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/entities v0.0.0-00010101000000-000000000000
	parsdevkit.net/persistence/repositories v0.0.0-00010101000000-000000000000
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000
)

require (
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
	parsdevkit.net/persistence/contexts v0.0.0-00010101000000-000000000000 // indirect
	parsdevkit.net/structs v0.0.0-00010101000000-000000000000 // indirect
)
