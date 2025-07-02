package application

import (
	"log"

	"parsdevkit.net/structs/workspace"

	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"
)

func GetContext() *ApplicationContext {
	workspaceService := ioc.Get[contracts.WorkspaceServiceInterface[workspace.WorkspaceBaseStruct]]()

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

	if currentWorkspace != nil {
		context := ApplicationContext{
			CurrentWorkspace: currentWorkspace.GetHeader(),
		}
		return &context
	}
	return nil
}
