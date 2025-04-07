package remove

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"parsdevkit.net/engines/applicationProject"
	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	names         []string
	workspaceName string
	force         string
	filePaths     []string
)

var RemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Short:   "Project Removing",
	Long:    `Project Removing`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	if len(filePaths) > 0 {
		applicationProjectService := applicationProject.ApplicationProjectEngine{}
		if err := applicationProjectService.RemoveProjectsFromFile(workspaceName, true, filePaths...); err != nil {
			log.Fatal(err)
		}
	}

	if len(args) > 0 {
		names = args

		applicationProjectService := services.NewApplicationProjectService(utils.GetEnvironment())
		for _, name := range names {
			applicationProject, err := applicationProjectService.Remove(name, workspaceName, false, true)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Project (" + applicationProject.Name + ") deleted permanently")
		}
	} else if len(filePaths) > 0 {
		applicationProjectService := applicationProject.ApplicationProjectEngine{}
		if err := applicationProjectService.RemoveProjectsFromFile(workspaceName, true, filePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please provide a name for the project")
		os.Exit(1)
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	RemoveCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)

	RemoveCmd.Flags().StringSliceVarP(&filePaths, "files", "f", nil, "Comma-separated list of declaration files")
	RemoveCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listProjectNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
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

func fileFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	files, _ := filepath.Glob(filepath.Join(toComplete, "*"))
	completions := []string{}
	for _, file := range files {
		if info, err := os.Stat(file); err == nil && !info.IsDir() {
			completions = append(completions, file)
		}
	}
	return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
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
