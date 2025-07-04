package data_resource_contract

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/modules/resource/data_resource_payload"
)

type ResourceInterface contracts.ResourceServiceInterface[data_resource_payload.ResourceBaseStruct]
