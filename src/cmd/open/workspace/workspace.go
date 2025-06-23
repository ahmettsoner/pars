package workspace

import (
	"fmt"

	"parsdevkit.net/providers"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

type WorkspaceOptions struct {
	Name string
}

var commandOptions WorkspaceOptions
var maxArgumentCount int = 1

var WorkspaceCommand = &cobra.Command{
	Use:     "workspace",
	Aliases: []string{"w"},
	Short:   "Workspace workspace",
	Long:    `Workspace workspace`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: workspace name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only workspace name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	path := parsCMDCommon.GetActiveWorkspacePath(commandOptions.Name)

	providers.VSCodeExecute("", path)

	return nil
}

func init() {
	WorkspaceCommand.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Workspace name")
}
