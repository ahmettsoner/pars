package test

import (
	"fmt"
	"os"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/components/workspace"
	_string "parsdevkit.net/core/utilities/string"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
	platformsCommon "parsdevkit.net/platforms/common"
)

type CleanOptions struct {
	Name      string
	Workspace string
}

var commandOptions CleanOptions
var maxArgumentCount int = 1

var TestCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Test project(s)",
	Long:    `Test project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: project name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only project name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	if _string.IsEmpty(commandOptions.Workspace) {
		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		commandOptions.Workspace = workspace.GetActiveWorkspaceName(appCtx, "")
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	projectService := services.NewApplicationProjectService(utils.GetEnvironment(), platformsCommon.Registry)
	_, err := projectService.Test(commandOptions.Name, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to test project '%s'\n%w", commandOptions.Name, err)
	}

	fmt.Fprintf(os.Stdout, "✔ Project '%v' tested successfully\n", commandOptions.Name)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = CleanOptions{}
}

func init() {
	TestCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	TestCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
