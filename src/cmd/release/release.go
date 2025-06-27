package release

import (
	"fmt"
	"os"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

type ReleaseOptions struct {
	Name      string
	Workspace string
}

var commandOptions ReleaseOptions
var maxArgumentCount int = 1

var ReleaseCmd = &cobra.Command{
	Use:     "release",
	Aliases: []string{"r"},
	Short:   "Release project(s)",
	Long:    `Release project(s)`,
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
		return fmt.Errorf("error: too many arguments. Only project name is expected.")
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
	_, err := projectService.Release(commandOptions.Name, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to execute project '%s'\n%w", commandOptions.Name, err)
	}

	fmt.Fprintf(os.Stdout, "✔ Project '%v' released successfully\n", commandOptions.Name)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ReleaseOptions{}
}

func init() {
	ReleaseCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	ReleaseCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
