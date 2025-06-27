package group

import (
	"parsdevkit.net/cmd/group/describe"
	"parsdevkit.net/cmd/group/list"
	"parsdevkit.net/cmd/group/remove"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Group Information",
	Long:    `Group Information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	GroupCmd.AddCommand(remove.RemoveCmd)
	GroupCmd.AddCommand(list.ListCmd)
	GroupCmd.AddCommand(describe.DescribeCmd)
}
