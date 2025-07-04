package remove

import (
	"fmt"
	"log"
	"os"
	"strings"

	"parsdevkit.net/pkg/utilities/array"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/group/basic_group_contract"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove [name]...",
	Aliases:           []string{"r"},
	Short:             "Group Information",
	Long:              `Group Information`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: group name is required.")
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

	if len(commandOptions.Names) > 0 {

		groupService := ioc.Get[basic_group_contract.GroupInterface]()
		for _, name := range commandOptions.Names {
			_, err := groupService.Remove(name, true)
			if err != nil {
				return fmt.Errorf("Failed to remove group(s) '%s'\n%w", name, err)
			}
		}
		fmt.Fprintf(os.Stdout, "✔ Group(s) '%v' removed successfully\n", commandOptions.Names)
	}
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = RemoveOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().StringSliceVarP(&commandOptions.Names, "names", "n", nil, "Comma-separated list of names")
	// RemoveCmd.RegisterFlagCompletionFunc("name", nameFlagCompletion)

}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listGroupNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func listGroupNameSuggestions(args []string, toComplete string) []string {

	// workspaceName = workspace.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	groupService := ioc.Get[basic_group_contract.GroupInterface]()
	groupList, err := groupService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range *groupList {
		if !array.ContainsSlice(args, group.Header.Name) && strings.HasPrefix(group.Header.Name, toComplete) {
			suggestions = append(suggestions, group.Header.Name)
		}
	}
	return suggestions
}

func listWorkspaceNameSuggestions(args []string, toComplete string) []string {
	var suggestions = make([]string, 0)
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, workspace := range *workspaceList {
		if !array.ContainsSlice(args, workspace.Header.Name) && strings.HasPrefix(workspace.Header.Name, toComplete) {
			suggestions = append(suggestions, workspace.Header.Name)
		}
	}
	return suggestions
}
