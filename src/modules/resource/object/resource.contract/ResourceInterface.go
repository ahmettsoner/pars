package object_resource_contract

import (
	resource_payload "parsdevkit.net/modules/resource/object_resource_payload"
	"parsdevkit.net/object/contracts"
)

type ResourceInterface contracts.ResourceServiceInterface[resource_payload.ResourceBaseStruct]
