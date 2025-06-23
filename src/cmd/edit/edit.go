package edit

import (
	"parsdevkit.net/cmd/edit/project"

	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "Edit in editor",
	Long:    `Edit in editor`,
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	EditCmd.AddCommand(project.ProjectCommand)
}
