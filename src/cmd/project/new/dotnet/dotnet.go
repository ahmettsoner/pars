package dotnet

import (
	"parsdevkit.net/cmd/project/new/dotnet/console"
	"parsdevkit.net/cmd/project/new/dotnet/library"
	"parsdevkit.net/cmd/project/new/dotnet/webapi"
	"parsdevkit.net/cmd/project/new/dotnet/webapp"

	"github.com/spf13/cobra"
)

var DotnetCmd = &cobra.Command{
	Use:     "dotnet",
	Aliases: []string{"d"},
	Short:   "Create new Dotnet project",
	Long:    `Create new Dotnet project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	DotnetCmd.AddCommand(webapi.WebApiCmd)
	DotnetCmd.AddCommand(webapp.WebAppCmd)
	DotnetCmd.AddCommand(library.LibraryCmd)
	DotnetCmd.AddCommand(console.ConsoleCmd)
}
