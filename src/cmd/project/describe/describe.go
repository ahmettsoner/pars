package describe

import (
	"fmt"
	"log"
	"os"
	"strings"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	name          string
	workspaceName string
	force         string
)

var DescribeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "Information about project",
	Long:    `Information about project`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("Required at least 1 project name")
		}
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if utils.IsEmpty(name) {
		if len(args) == 0 {
			fmt.Println("Please provide a name for the new project")
			os.Exit(1)
		} else if len(args) > 0 {
			name = args[0]
		}
	}

	if utils.IsEmpty(name) {
		cmd.Help()
		os.Exit(0)
	}

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	projectList, err := projectService.ListByFullNameWorkspace(name, workspaceName)
	if err != nil {
		log.Fatal(err)
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
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	DescribeCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	DescribeCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	DescribeCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}

func listProjectNameSuggestions(args []string, toComplete string) []string {

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	projectList, err := projectService.ListByWorkspace(workspaceName)
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
