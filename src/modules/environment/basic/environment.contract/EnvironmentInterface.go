package basic_environment_contract

import (
	"parsdevkit.net/application/contracts"
	basic_environment_payload_structs "parsdevkit.net/modules/environment/basic_environment_payload/structs"
)

type EnvironmentInterface contracts.EnvironmentServiceInterface[basic_environment_payload_structs.EnvironmentBaseStruct]
