package file_template_contract

import (
	"parsdevkit.net/application/contracts"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
)

type TemplateInterface contracts.TemplateServiceInterface[file_template_payload_structs.TemplateBaseStruct]
