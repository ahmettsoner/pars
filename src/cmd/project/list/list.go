package list

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/components/workspace"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/pkg/utilities/array"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	Workspace string
}

var commandOptions ListOptions
var maxArgumentCount int = 0

var ListCmd = &cobra.Command{
	Use:               "list",
	Aliases:           []string{"l"},
	Short:             "List project(s)",
	Long:              `List project(s)`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("There is no argument supported")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Workspace) {
		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		commandOptions.Workspace = workspace.GetActiveWorkspaceName(appCtx, "")
	}
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	var result []string = []string{application_project_payload_structs.MODULE_KEY}

	appCtx := application.GetContext()

	err := engines.DispatchEngineList(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ListOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ListCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
	ListCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := make([]string, 0)

	return suggestions, cobra.ShellCompDirectiveNoFileComp
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

func workspaceFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := []string{}

	suggestions := listWorkspaceNameSuggestions(args, toComplete)
	completions = append(completions, suggestions...)

	return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
}
