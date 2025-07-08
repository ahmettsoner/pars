module parsdevkit.net/modules/template/shared_template_contract

replace parsdevkit.net/modules/template/shared_template_payload => ../template.payload

replace parsdevkit.net/shared => ../../../../shared

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	parsdevkit.net/shared v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/shared_template_payload v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
