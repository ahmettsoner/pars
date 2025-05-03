package project

import (
	"parsdevkit.net/cmd/project/describe"
	"parsdevkit.net/cmd/project/list"
	"parsdevkit.net/cmd/project/new"
	"parsdevkit.net/cmd/project/remove"
	"parsdevkit.net/cmd/project/submit"

	"github.com/spf13/cobra"
)

var ProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Project Information",
	Long:    `Project Information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ProjectCmd.AddCommand(new.NewCmd)
	ProjectCmd.AddCommand(submit.SubmitCmd)
	ProjectCmd.AddCommand(remove.RemoveCmd)
	ProjectCmd.AddCommand(describe.DescribeCmd)
	ProjectCmd.AddCommand(list.ListCmd)
}
