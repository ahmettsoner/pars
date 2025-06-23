package angular

import (
	"parsdevkit.net/cmd/project/new/angular/library"
	"parsdevkit.net/cmd/project/new/angular/spa"

	"github.com/spf13/cobra"
)

var AngularCmd = &cobra.Command{
	Use:     "angular",
	Aliases: []string{"a"},
	Short:   "Create new Angular project",
	Long:    `Create new Angular project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	AngularCmd.AddCommand(spa.SPACmd)
	AngularCmd.AddCommand(library.LibraryCmd)
}
