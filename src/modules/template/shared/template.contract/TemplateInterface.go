package shared_template_contract

import (
	"parsdevkit.net/application/contracts"
	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"
)

type TemplateInterface contracts.TemplateServiceInterface[shared_template_payload_structs.TemplateBaseStruct]
