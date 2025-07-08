package object_resource_contract

import (
	"parsdevkit.net/application/contracts"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
)

type ResourceInterface contracts.ResourceServiceInterface[object_resource_payload_structs.ResourceBaseStruct]
