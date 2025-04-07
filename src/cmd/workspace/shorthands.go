package workspace

import (
	"parsdevkit.net/cmd/workspace/list"

	"github.com/spf13/cobra"
)

var WorkspaceListShorthandsCmd = &cobra.Command{
	Use:    "wl",
	Short:  "Alias for 'workspace list'",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		list.ListCommand.Run(cmd, args)
	},
}
