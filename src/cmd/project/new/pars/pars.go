package pars

import (
	"parsdevkit.net/cmd/project/new/pars/project"

	"github.com/spf13/cobra"
)

var ParsCmd = &cobra.Command{
	Use:     "pars",
	Aliases: []string{"p"},
	Short:   "Create new Pars project",
	Long:    `Create new Pars project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ParsCmd.AddCommand(project.ProjectCmd)
}
