package task

import (
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:     "task",
	Aliases: []string{},
	Short:   "Task Information",
	Long:    `Task Information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	// TaskCmd.AddCommand(new.SubmitCmd)
	// TaskCmd.AddCommand(remove.RemoveCmd)
	// TaskCmd.AddCommand(list.ListCmd)
	// TaskCmd.AddCommand(describe.DescribeCmd)
}
