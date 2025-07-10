package browse_tool_contract

import (
	"parsdevkit.net/application/contracts"
	browse_tool_payload_structs "parsdevkit.net/modules/tool/browse_tool_payload/structs"
)

type ToolInterface contracts.ToolServiceInterface[browse_tool_payload_structs.ToolBaseStruct]
