package application_project_contract

import (
	"parsdevkit.net/application/contracts"
	// project_payload "parsdevkit.net/modules/project/application_project_payload"
	project_payload "parsdevkit.net/structs/project/application-project"
)

type ProjectInterface contracts.ProjectServiceInterface[project_payload.ProjectBaseStruct]
