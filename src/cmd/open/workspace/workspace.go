package workspace

import (
	"fmt"

	"parsdevkit.net/components/workspace"
	"parsdevkit.net/providers"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
	_string "parsdevkit.net/pkg/utilities/string"
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
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: workspace name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only workspace name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	path := workspace.GetActiveWorkspacePath(appCtx, commandOptions.Name)

	providers.VSCodeExecute("", path)

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = WorkspaceOptions{}
}

func init() {
	WorkspaceCommand.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Workspace name")
}
