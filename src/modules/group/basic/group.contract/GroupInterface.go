package basic_group_contract

import (
	"parsdevkit.net/application/contracts"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

type GroupInterface contracts.GroupServiceInterface[basic_group_payload_structs.GroupBaseStruct]
