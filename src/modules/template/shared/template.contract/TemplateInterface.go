package shared_template_contract

import (
	// template_payload "parsdevkit.net/modules/template/shared_template_payload"
	"parsdevkit.net/shared/contracts"
	template_payload "parsdevkit.net/structs/template/shared-template"
)

type TemplateInterface contracts.TemplateServiceInterface[template_payload.TemplateBaseStruct]
