package basic_workspace_contract

import (
	workspace_payload "parsdevkit.net/modules/workspace/basic_workspace_payload"
	"parsdevkit.net/workspace/contracts"
)

type WorkspaceInterface contracts.WorkspaceServiceInterface[workspace_payload.WorkspaceBaseStruct]
