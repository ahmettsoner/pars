package execute

import (
	"fmt"
	"os"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/application"
	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
	platformsCommon "parsdevkit.net/platforms/common"
)

type ExecuteOptions struct {
	Name      string
	Workspace string
}

var commandOptions ExecuteOptions
var maxArgumentCount int = 1

var ExecuteCmd = &cobra.Command{
	Use:     "execute",
	Aliases: []string{"x"},
	Short:   "Execute project(s)",
	Long:    `Execute project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: project name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	var workspaceName, err = parsCMDCommon.GetActiveWorkspaceNameV2(appCtx, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("failed to find active workspace '%s'\n%w", commandOptions.Name, err)
	}
	commandOptions.Workspace = workspaceName

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	projectService := services.NewApplicationProjectService(utils.GetEnvironment(), platformsCommon.Registry)
	project, err := projectService.Clean(commandOptions.Name, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to execute project '%s'\n%w", commandOptions.Name, err)
	}

	fmt.Fprintf(os.Stdout, "✔ Project '%s' executed successfully\n", project.Header.Name)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ExecuteOptions{}
}

func init() {
	ExecuteCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	ExecuteCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
