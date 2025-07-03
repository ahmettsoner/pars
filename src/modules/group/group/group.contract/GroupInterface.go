package groupcontract

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/modules/group/group_payload"
)

type GroupInterface contracts.GroupServiceInterface[group_payload.GroupBaseStruct]
