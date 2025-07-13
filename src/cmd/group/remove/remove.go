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
	_string "parsdevkit.net/pkg/utilities/string"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Name  string
	Force string
}

var commandOptions RemoveOptions
var maxArgumentCount int = 1

var RemoveCmd = &cobra.Command{
	Use:               "remove [name]",
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
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: group name is required. Provide as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	var result []string = []string{basic_group_payload_structs.MODULE_KEY}

	appCtx := application.GetContext()

	err := engines.DispatchEngineRemove(appCtx, result, commandOptions.Name)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) < maxArgumentCount {

		if len(args) == 0 {
			suggestions := listGroupNameSuggestions(args, toComplete)

			return suggestions, cobra.ShellCompDirectiveNoSpace
		}
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
