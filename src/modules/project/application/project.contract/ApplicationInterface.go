package application_project_contract

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/modules/project/application_project_payload"
)

type ProjectInterface contracts.ProjectServiceInterface[application_project_payload.ProjectBaseStruct]
