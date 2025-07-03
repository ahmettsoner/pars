package data_resource_contract

import (
	"parsdevkit.net/data/contracts"
	resource_payload "parsdevkit.net/modules/resource/data_resource_payload"
)

type ResourceInterface contracts.ResourceServiceInterface[resource_payload.ResourceBaseStruct]
