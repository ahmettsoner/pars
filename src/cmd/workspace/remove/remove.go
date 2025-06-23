package remove

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names []string
	Force bool
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove name [name]...",
	Aliases:           []string{"r"},
	Short:             "Workspace removing",
	Long:              `Workspace removing`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: workspace name is required. Provide it with '--name' or as an argument.")
	}

	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) > 0 {
		commandOptions.Names = args
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	for _, name := range commandOptions.Names {
		workspace, err := workspaceService.Remove(name, commandOptions.Force, true)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Workspace (" + workspace.Name + ") deleted permanently")
	}

	return nil
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listWorkspaceNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().BoolVarP(&commandOptions.Force, "force", "f", false, "Workspace name")
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
