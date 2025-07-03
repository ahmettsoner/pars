package describe

import (
	"parsdevkit.net/application/contracts"
	"parsdevkit.net/application/ioc"
	applicationproject "parsdevkit.net/structs/project/application-project"

	"fmt"
	"log"
	"strings"

	"parsdevkit.net/pkg/utilities/array"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/components/workspace"

	workspaceStruct "parsdevkit.net/structs/workspace"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
)

type DescribeOptions struct {
	Name      string
	Workspace string
	Force     string
}

var commandOptions DescribeOptions
var maxArgumentCount int = 1

var DescribeCmd = &cobra.Command{
	Use:               "describe",
	Aliases:           []string{"d"},
	Short:             "Information about project",
	Long:              `Information about project`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: group name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only group name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
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

	projectService := ioc.Get[contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct]]()
	projectList, err := projectService.ListByFullNameWorkspace(commandOptions.Name, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to describe project '%s'\n%w", commandOptions.Name, err)
	}

	for _, e := range *projectList {
		name := fmt.Sprintf(" - %v", e.GetFullInformation())
		fmt.Println(name)
		labels := fmt.Sprintf("\t Labels: %v", e.Specifications.Labels)
		fmt.Println(labels)
		projectPlatform := fmt.Sprintf("\t Platform: %v", e.Specifications.Platform.Type.String())
		fmt.Println(projectPlatform)
		projectType := fmt.Sprintf("\t Type: %v", e.Specifications.ProjectType)
		fmt.Println(projectType)
		projectRuntime := fmt.Sprintf("\t Runtime: %v", e.Specifications.Runtime.Type.String())
		fmt.Println(projectRuntime)
		layers := fmt.Sprintf("\t Layers: %v", e.Specifications.Configuration.Layers)
		fmt.Println(layers)
	}
	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = DescribeOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	DescribeCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	DescribeCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
	DescribeCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}

func listProjectNameSuggestions(args []string, toComplete string) []string {

	var suggestions = make([]string, 0)
	projectService := ioc.Get[contracts.ProjectServiceInterface[applicationproject.ProjectBaseStruct]]()

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

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listProjectNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func listWorkspaceNameSuggestions(args []string, toComplete string) []string {
	var suggestions = make([]string, 0)
	workspaceService := ioc.Get[contracts.WorkspaceServiceInterface[workspaceStruct.WorkspaceBaseStruct]]()
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
