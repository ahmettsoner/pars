package engines

import (
	"log"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core"
	"parsdevkit.net/core/utils"
)

func GetContext() *core.Context {
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())

	currentWorkspace, err := workspaceService.GetActiveWorkspace()
	if err != nil {
		log.Fatal(err)
	}

	if currentWorkspace == nil {
		currentWorkspace, err = workspaceService.GetSelectedWorkspace()
		if err != nil {
			log.Fatal(err)
		}
	}

	context := core.Context{
		CurrentWorkspace: currentWorkspace.GetHeader(),
	}

	return &context
}
