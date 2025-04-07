package work

import (
	"github.com/spf13/cobra"
)

var (
	name string
)

var WorkCmd = &cobra.Command{
	Use:     "work",
	Aliases: []string{"o"},
	Short:   "Change directory",
	Long:    `Change directory`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	// WorkCmd.AddCommand(workspace.WorkspaceCommand)
}
