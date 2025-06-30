package clean

import (
	"fmt"
	"os"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

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

var CleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   "Clean project(s)",
	Long:    `Clean project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("Please provide project name using the --url flag or as argument.")
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

	project, err := projectService.CleanV2(commandOptions.Name, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to clean project '%s'\n%w", commandOptions.Name, err)
	}

	fmt.Fprintf(os.Stdout, "✔ Project '%s' cleaned successfully\n", project.Header.Name)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = CleanOptions{}
}

func init() {
	CleanCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	CleanCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
