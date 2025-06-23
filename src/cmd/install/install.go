package install

import (
	"fmt"
	"log"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

type InstallOptions struct {
	Name      string
	Workspace string
}

var commandOptions InstallOptions
var maxArgumentCount int = 1

var InstallCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{""},
	Short:   "Install project(s)",
	Long:    `Install project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: name is required. Provide it with '--name' or as an argument.")
	}

	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	if utils.IsEmpty(commandOptions.Workspace) {
		commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName("")
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	project, err := projectService.Install(commandOptions.Name, commandOptions.Workspace)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project (" + project.Name + ") packages installed")

	return nil
}

func init() {
	InstallCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	InstallCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
	// RemoveCommand.Flags().StringVarP(&force, "force", "", "", "Force to delete")

	// if err := RemoveCommand.MarkFlagRequired("force"); err != nil {
	// 	fmt.Println(err)
	// }
}
