package open

import (
	"parsdevkit.net/providers"

	"parsdevkit.net/cmd/open/project"
	"parsdevkit.net/cmd/open/workspace"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	name string
)

var OpenCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "Open in editor",
	Long:    `Open in editor`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {

	path := parsCMDCommon.GetActiveWorkspacePath(name)

	providers.VSCodeExecute("", path)
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	OpenCmd.AddCommand(workspace.WorkspaceCommand)
	OpenCmd.AddCommand(project.ProjectCmd)
}
