package edit

import (
	"parsdevkit.net/cmd/edit/project"

	"github.com/spf13/cobra"
)

var (
	name string
)

var EditCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "Edit in editor",
	Long:    `Edit in editor`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	EditCmd.AddCommand(project.WorkspaceCommand)
}
