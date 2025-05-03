package engines

import (
	"log"

	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"
)

type Context struct {
	CurrentWorkspace *workspace.WorkspaceBaseStruct
}

func GetContext() *Context {
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

	context := Context{
		CurrentWorkspace: currentWorkspace,
	}

	return &context
}
