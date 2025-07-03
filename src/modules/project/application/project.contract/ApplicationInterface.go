package application_project_contract

import (
	"parsdevkit.net/application/contracts"
	// project_payload "parsdevkit.net/modules/project/application_project_payload"
	applicationProjectSchema "parsdevkit.net/structs/project/application-project"
)

type ProjectInterface contracts.ProjectServiceInterface[applicationProjectSchema.ProjectBaseStruct]
