package remote

import (
	"fmt"
	"os"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	name          string
	workspaceName string
)

var RemoteCmd = &cobra.Command{
	Use:     "remote",
	Aliases: []string{""},
	Short:   "Remote project(s)",
	Long:    `Remote project(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if utils.IsEmpty(name) {
		if len(args) == 0 {
			fmt.Println("Please provide a project name")
			os.Exit(1)
		} else if len(args) > 0 {
			name = args[0]
		}
	}

	if utils.IsEmpty(name) {
		cmd.Help()
		os.Exit(0)
	}

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	// projectService := services.NewApplicationProjectService()
	// project, err := projectService.Remote(name, workspaceID)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Project (" + (*project).Definition.Name + ") remoteed")
}

func init() {
	RemoteCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	RemoteCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	// RemoveCommand.Flags().StringVarP(&force, "force", "", "", "Force to delete")

	// if err := RemoveCommand.MarkFlagRequired("force"); err != nil {
	// 	fmt.Println(err)
	// }
}
