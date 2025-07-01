package workspace

import (
	"fmt"
	"log"

	"parsdevkit.net/core"
	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utilities"
	"parsdevkit.net/core/utils"
)

func GetActiveWorkspaceNameV2(ctx *core.ApplicationContext, workspaceName string) (string, error) {
	if !utilities.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		ok, err := workspaceService.IsExists(workspaceName)
		if err != nil {
			return "", fmt.Errorf("xxx: workspace ('%s') kontrolünde hata oluştu\n%w", workspaceName, err)
		}
		if !ok {
			return "", fmt.Errorf("workspace '%s' does not exist", workspaceName)
		}
		return workspaceName, nil
	}

	name := ctx.CurrentWorkspace.Name

	if utilities.IsEmpty(name) {
		return "", fmt.Errorf("no active workspace found; please initialize or switch to one")
	}

	return name, nil
}

func GetActiveWorkspaceName(ctx *core.ApplicationContext, workspaceName string) string {

	if !utilities.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			log.Fatal(err)
		}
		if workspace == nil {
			log.Fatal("Workspace name (" + workspaceName + ") is not correct")
		}
	} else {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	if utilities.IsEmpty(workspaceName) {
		log.Fatal("There are no active workspace, please initialize or switch to available workspace")
	}

	return workspaceName
}

func GetActiveWorkspacePath(ctx *core.ApplicationContext, workspaceName string) string {

	if utilities.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspace, err := workspaceService.GetByName(workspaceName)
	if err != nil {
		log.Fatal(err)
	}
	if workspace == nil {
		log.Fatal("Workspace name (" + workspaceName + ") is not correct")
	}
	return workspace.Specifications.GetAbsolutePath()

}
