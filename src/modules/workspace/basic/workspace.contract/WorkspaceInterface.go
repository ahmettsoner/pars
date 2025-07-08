package basic_workspace_contract

import (
	"parsdevkit.net/application/contracts"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type WorkspaceInterface contracts.WorkspaceServiceInterface[basic_workspace_payload_structs.WorkspaceBaseStruct]
