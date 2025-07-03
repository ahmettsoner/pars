package basic_group_contract

import (
	"parsdevkit.net/application/contracts"
	group_payload "parsdevkit.net/modules/group/basic_group_payload"
)

type GroupInterface contracts.GroupServiceInterface[group_payload.GroupBaseStruct]
