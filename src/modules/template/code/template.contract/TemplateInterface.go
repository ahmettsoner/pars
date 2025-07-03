package code_template_contract

import (
	"parsdevkit.net/code/contracts"
	// template_payload "parsdevkit.net/modules/template/code_template_payload"
	template_payload "parsdevkit.net/structs/template/code-template"
)

type TemplateInterface contracts.TemplateServiceInterface[template_payload.TemplateBaseStruct]
