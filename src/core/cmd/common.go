package cmd

import (
	"fmt"
	"log"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	manifestServices "parsdevkit.net/engines"
)

func GetActiveWorkspaceNameV2(workspaceName string) (string, error) {
	if !utils.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		if !workspaceService.IsExists(workspaceName) {
			return "", fmt.Errorf("workspace '%s' does not exist", workspaceName)
		}
		return workspaceName, nil
	}

	// workspaceName boşsa context'ten al
	appContext := manifestServices.GetContext()
	name := appContext.CurrentWorkspace.Name

	if utils.IsEmpty(name) {
		return "", fmt.Errorf("no active workspace found; please initialize or switch to one")
	}

	return name, nil
}

func GetActiveWorkspaceName(workspaceName string) string {

	appContext := manifestServices.GetContext()

	if !utils.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			log.Fatal(err)
		}
		if workspace == nil {
			log.Fatal("Workspace name (" + workspaceName + ") is not correct")
		}
	} else {
		workspaceName = appContext.CurrentWorkspace.Name
	}

	if utils.IsEmpty(workspaceName) {
		log.Fatal("There are no active workspace, please initialize or switch to available workspace")
	}

	return workspaceName
}

func GetActiveWorkspacePath(workspaceName string) string {

	appContext := manifestServices.GetContext()

	if !utils.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			log.Fatal(err)
		}
		if workspace == nil {
			log.Fatal("Workspace name (" + workspaceName + ") is not correct")
		}
		return workspace.Specifications.GetAbsolutePath()
	} else {
		return appContext.CurrentWorkspace.Specifications.GetAbsolutePath()
	}

}
