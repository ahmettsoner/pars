package object_resource_contract

import (
	// resource_payload "parsdevkit.net/modules/resource/object_resource_payload"
	"parsdevkit.net/application/contracts"
	resource_payload "parsdevkit.net/structs/resource/object-resource"
)

type ResourceInterface contracts.ResourceServiceInterface[resource_payload.ResourceBaseStruct]
