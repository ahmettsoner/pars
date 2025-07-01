package open

import (
	"fmt"
	"os"

	"parsdevkit.net/providers"

	"parsdevkit.net/application"
	"parsdevkit.net/cmd/open/project"
	workspaceCommand "parsdevkit.net/cmd/open/workspace"

	"parsdevkit.net/components/workspace"
	"parsdevkit.net/core/utilities"

	"github.com/spf13/cobra"
)

type OpenOptions struct {
	Name string
}

var commandOptions OpenOptions
var maxArgumentCount int = 1

var OpenCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "Open in editor",
	Long:    `Open in editor`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utilities.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: workspace name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only workspace name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utilities.IsEmpty(commandOptions.Name) && len(args) > 0 {
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

	fmt.Fprintf(os.Stdout, "✔ Project '%v' opend successfully\n", commandOptions.Name)
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = OpenOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	OpenCmd.AddCommand(workspaceCommand.WorkspaceCommand)
	OpenCmd.AddCommand(project.ProjectCmd)
}
