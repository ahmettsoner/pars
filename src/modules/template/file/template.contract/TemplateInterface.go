package code_template_contract

import (
	"parsdevkit.net/application/contracts"
	// template_payload "parsdevkit.net/modules/template/file_template_payload"
	template_payload "parsdevkit.net/structs/template/file-template"
)

type TemplateInterface contracts.TemplateServiceInterface[template_payload.TemplateBaseStruct]
