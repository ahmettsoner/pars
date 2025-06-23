package new

import (
	"fmt"

	"github.com/spf13/cobra"
)

var maxArgumentCount int = 0

var NewCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Initialize project",
	Long:    `Initialize project`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > maxArgumentCount {
			return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
		}
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	// NewCmd.AddCommand(pars.ParsCmd)
	// NewCmd.AddCommand(dotnet.DotnetCmd)
	// NewCmd.AddCommand(angular.AngularCmd)
	// NewCmd.AddCommand(goPkg.GoCmd)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// if len(args) == 1 {
	// 	completions := []string{}

	// 	return completions, cobra.ShellCompDirectiveNoSpace
	// }

	return nil, cobra.ShellCompDirectiveNoFileComp
}
