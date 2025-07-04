package basic_workspace_contract

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/modules/workspace/basic_workspace_payload"
)

type WorkspaceInterface contracts.WorkspaceServiceInterface[basic_workspace_payload.WorkspaceBaseStruct]
