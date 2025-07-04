package object_resource_contract

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/modules/resource/object_resource_payload"
)

type ResourceInterface contracts.ResourceServiceInterface[object_resource_payload.ResourceBaseStruct]
