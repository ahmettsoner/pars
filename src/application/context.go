package application

import (
	"log"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
)

func GetContext() *ApplicationContext {
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

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
