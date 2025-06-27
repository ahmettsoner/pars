package describe

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
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
	if utils.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: group name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only group name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	if utils.IsEmpty(commandOptions.Workspace) {
		commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName("")
	}
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
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
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	projectList, err := projectService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		log.Fatal(err)
	}

	for _, project := range *projectList {
		if !utils.Contains(args, project.GetFullName()) && strings.HasPrefix(project.GetFullName(), toComplete) {
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

func workspaceFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := []string{}

	suggestions := listWorkspaceNameSuggestions(args, toComplete)
	completions = append(completions, suggestions...)

	return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
}
