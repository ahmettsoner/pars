package basic_workspace_contract

import (
	"parsdevkit.net/application/contracts"
	workspace_payload "parsdevkit.net/modules/workspace/basic_workspace_payload"
)

type GroupInterface contracts.GroupServiceInterface[workspace_payload.GroupBaseStruct]
