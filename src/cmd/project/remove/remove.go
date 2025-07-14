package remove

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"

	"parsdevkit.net/pkg/utilities/array"

	"parsdevkit.net/application/engines"

	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
}

var commandOptions RemoveOptions
var maxArgumentCount int = 1

var RemoveCmd = &cobra.Command{
	Use:               "remove name [name]...",
	Aliases:           []string{"r"},
	Short:             "Project Removing",
	Long:              `Project Removing`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: project name is required.")
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

	var result []string = []string{application_project_payload_structs.MODULE_KEY}

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

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
	RemoveCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listProjectNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func listProjectNameSuggestions(args []string, toComplete string) []string {

	var suggestions = make([]string, 0)
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	projectList, err := projectService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		log.Fatal(err)
	}

	for _, project := range *projectList {
		if !array.ContainsSlice(args, project.GetFullName()) && strings.HasPrefix(project.GetFullName(), toComplete) {
			suggestions = append(suggestions, project.GetFullName())
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

func workspaceFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := []string{}

	suggestions := listWorkspaceNameSuggestions(args, toComplete)
	completions = append(completions, suggestions...)

	return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
}
