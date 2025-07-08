package core

import (
	"parsdevkit.net/application/platforms"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type ApplicationPlatformManagerInterface platforms.PlatformInterface[application_project_payload_structs.ProjectBaseStruct]
