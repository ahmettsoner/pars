package _go

import (
	"parsdevkit.net/cmd/project/new/dotnet/console"

	"github.com/spf13/cobra"
)

var GoCmd = &cobra.Command{
	Use:     "go",
	Aliases: []string{"g"},
	Short:   "Create new Go project",
	Long:    `Create new Go project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	GoCmd.AddCommand(console.ConsoleCmd)
}
