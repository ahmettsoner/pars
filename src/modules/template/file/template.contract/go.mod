module parsdevkit.net/modules/template/file_template_contract

replace parsdevkit.net/modules/template/code_template_payload => ../template.payload

replace parsdevkit.net/code => ../../../../code

replace parsdevkit.net/pkg => ../../../../pkg

go 1.23.7

require (
	parsdevkit.net/code v0.0.0-00010101000000-000000000000
	parsdevkit.net/modules/template/code_template_payload v0.0.0-00010101000000-000000000000
)

require (
	gopkg.in/yaml.v3 v3.0.1 // indirect
	parsdevkit.net/pkg v0.0.0-00010101000000-000000000000 // indirect
)
