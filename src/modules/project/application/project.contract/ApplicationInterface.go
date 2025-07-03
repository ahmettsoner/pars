package application_project_contract

import (
	"parsdevkit.net/application/contracts"
	project_payload "parsdevkit.net/modules/project/application_project_payload"
)

type ProjectInterface contracts.ProjectServiceInterface[project_payload.ProjectBaseStruct]
