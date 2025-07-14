package remove

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/pkg/utilities/array"

	"parsdevkit.net/application/ioc"

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
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: workspace name is required.")
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

	var result []string = []string{basic_workspace_payload_structs.MODULE_KEY}

	for _, name := range commandOptions.Names {

		appCtx := application.GetContext()

		err := engines.DispatchEngineRemove(appCtx, result, name)
		if err != nil {
			return fmt.Errorf("Engine processing failed: %v", err)
		}
	}
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = RemoveOptions{}
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
