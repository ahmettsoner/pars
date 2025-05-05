package environment

import (
	"parsdevkit.net/cmd/environment/list"

	"github.com/spf13/cobra"
)

var (
	switchTo string
)

var EnvironmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"env"},
	Short:   "Environment information",
	Long:    `Environment information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	EnvironmentCmd.AddCommand(list.ListCommand)
}
