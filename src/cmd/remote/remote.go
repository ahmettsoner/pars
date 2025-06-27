package remote

import (
	"fmt"

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

var RemoteCmd = &cobra.Command{
	Use:     "remote",
	Aliases: []string{""},
	Short:   "Remote project(s)",
	Long:    `Remote project(s)`,
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
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ReleaseOptions{}
}

func init() {
	RemoteCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	RemoteCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
