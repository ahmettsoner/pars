package data_resource_contract

import (
	"parsdevkit.net/application/contracts"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"
)

type ResourceInterface contracts.ResourceServiceInterface[data_resource_payload_structs.ResourceBaseStruct]
