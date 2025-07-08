package code_template_contract

import (
	"parsdevkit.net/application/contracts"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
)

type TemplateInterface contracts.TemplateServiceInterface[code_template_payload_structs.TemplateBaseStruct]
