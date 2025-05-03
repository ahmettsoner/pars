package remove

import (
	"fmt"
	"log"
	"os"
	"strings"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var (
	names []string
	force bool
)

var RemoveCmd = &cobra.Command{
	Use:     "remove name [name]...",
	Aliases: []string{"r"},
	Short:   "Workspace removing",
	Long:    `Workspace removing`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("Required at least 1 workspace name")
		}
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		names = args
	}

	if len(names) == 0 {
		var err error
		names, err = cmd.Flags().GetStringArray("name")
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(names) == 0 {
		fmt.Println("Please provide a name for the workspace")
		os.Exit(1)
	}

	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	for _, name := range names {
		workspace, err := workspaceService.Remove(name, force, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Workspace (" + workspace.Name + ") deleted permanently")
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().BoolVarP(&force, "force", "f", false, "Workspace name")
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listWorkspaceNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func listWorkspaceNameSuggestions(args []string, toComplete string) []string {
	var suggestions = make([]string, 0)
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, workspace := range *workspaceList {
		if !utils.Contains(args, workspace.Name) && strings.HasPrefix(workspace.Name, toComplete) {
			suggestions = append(suggestions, workspace.Name)
		}
	}
	return suggestions
}
