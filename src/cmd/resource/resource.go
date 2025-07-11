package resource

import (
	"parsdevkit.net/cmd/resource/describe"
	"parsdevkit.net/cmd/resource/list"
	"parsdevkit.net/cmd/resource/remove"

	"github.com/spf13/cobra"
)

var ResourceCmd = &cobra.Command{
	Use:     "resource",
	Aliases: []string{"r"},
	Short:   "Resource Information",
	Long:    `Resource Information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ResourceCmd.AddCommand(remove.RemoveCmd)
	ResourceCmd.AddCommand(list.ListCommand)
	ResourceCmd.AddCommand(describe.DescribeCmd)
}
