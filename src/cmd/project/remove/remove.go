package remove

import (
	"fmt"
	"log"
	"os"
	"strings"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"

	"parsdevkit.net/pkg/utilities/array"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/modules/project/application_project_contract"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"parsdevkit.net/components/workspace"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
	FilePaths []string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove [name]...",
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

	if len(commandOptions.Names) > 0 {

		projectService := ioc.Get[application_project_contract.ProjectInterface]()
		for _, name := range commandOptions.Names {
			_, err := projectService.Remove(name, commandOptions.Workspace, false, true)
			if err != nil {
				return fmt.Errorf("Failed to remove project(s) '%s'\n%w", name, err)
			}
		}
		fmt.Fprintf(os.Stdout, "✔ Project(s) '%v' removed successfully\n", commandOptions.Names)
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
