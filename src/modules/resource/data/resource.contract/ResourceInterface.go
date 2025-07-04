package data_resource_contract

import (
	"parsdevkit.net/application/contracts"
	// resource_payload "parsdevkit.net/modules/resource/data_resource_payload"
	resource_payload "parsdevkit.net/structs/resource/data-resource"
)

type ResourceInterface contracts.ResourceServiceInterface[resource_payload.ResourceBaseStruct]
