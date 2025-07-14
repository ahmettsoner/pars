package remove

import (
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/group/basic_group_contract"

	"log"
	"strings"

	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"

	"parsdevkit.net/application"

	"parsdevkit.net/application/engines"

	"parsdevkit.net/pkg/utilities/array"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names []string
	Force string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove name [name]...",
	Aliases:           []string{"r"},
	Short:             "Remove group",
	Long:              `Remove group`,
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

	var result []string = []string{basic_group_payload_structs.MODULE_KEY}

	for _, name := range commandOptions.Names {

		appCtx := application.GetContext()

		err := engines.DispatchEngineRemove(appCtx, result, name)
		if err != nil {
			return fmt.Errorf("Engine processing failed: %v", err)
		}
	}

	return nil
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	if len(args) == 0 {
		suggestions := listGroupNameSuggestions(args, toComplete)

		return suggestions, cobra.ShellCompDirectiveNoSpace
	}

	return make([]string, 0), cobra.ShellCompDirectiveNoFileComp
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = RemoveOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
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
