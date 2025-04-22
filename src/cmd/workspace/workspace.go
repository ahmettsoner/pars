package workspace

import (
	"fmt"
	"log"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/cmd/workspace/describe"
	"parsdevkit.net/cmd/workspace/list"
	"parsdevkit.net/cmd/workspace/remove"
	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var (
	switchTo string
)

var WorkspaceCmd = &cobra.Command{
	Use:     "workspace",
	Aliases: []string{"w"},
	Short:   "Workspace information",
	Long:    `Workspace information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if !utils.IsEmpty(switchTo) {

		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
		workspace, err := workspaceService.ChangeCurrentWorkspace(switchTo)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Swithched to: " + workspace.Name)
	} else {
		cmd.Help()
	}
}

func init() {
	addSubCommands()

	WorkspaceCmd.Flags().StringVarP(&switchTo, "switch", "s", "", "Switch to workspace")
	WorkspaceCmd.RegisterFlagCompletionFunc("switch", switchFlagCompletion)
}

func addSubCommands() {
	WorkspaceCmd.AddCommand(list.ListCommand)
	WorkspaceCmd.AddCommand(describe.DescribeCmd)
	WorkspaceCmd.AddCommand(remove.RemoveCmd)
}

func switchFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	var workspaces = make([]string, 0)
	for _, workspace := range *workspaceList {
		workspaces = append(workspaces, workspace.Name)
	}

	return workspaces, cobra.ShellCompDirectiveNoFileComp
}
