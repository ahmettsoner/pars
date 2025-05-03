package release

import (
	"fmt"
	"log"
	"os"
	"strings"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	name          string
	workspaceName string
)

var ReleaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"r"},
	Short:   "Release project(s)",
	Long:    `Release project(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if strings.TrimSpace(name) == "" {
		if len(args) == 0 {
			fmt.Println("Please provide a project name")
			os.Exit(1)
		} else if len(args) > 0 {
			name = args[0]
		}
	}

	if strings.TrimSpace(name) == "" {
		cmd.Help()
		os.Exit(0)
	}

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	project, err := projectService.Release(name, workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project (" + project.Name + ") releaseed")
}

func init() {
	ReleaseCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	ReleaseCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	// RemoveCommand.Flags().StringVarP(&force, "force", "", "", "Force to delete")

	// if err := RemoveCommand.MarkFlagRequired("force"); err != nil {
	// 	fmt.Println(err)
	// }
}
