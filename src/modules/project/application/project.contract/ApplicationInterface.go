package application_project_contract

import (
	"parsdevkit.net/application/contracts"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type ProjectInterface contracts.ProjectServiceInterface[application_project_payload_structs.ProjectBaseStruct]
