package list

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/operation/services"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var (
	workspaceName string
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List project(s)",
	Long:    `List project(s)`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("There is no argument supported")
		}
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	applicationProjectService := services.NewApplicationProjectService(utils.GetEnvironment())
	applicationProjectList, err := applicationProjectService.ListByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("(%d) application project available\n", len(*applicationProjectList))

	applicationProjectListBasic, err := applicationProjectService.ListIndividualByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	if len(*applicationProjectListBasic) > 0 {
		fmt.Println()
		for _, project := range *applicationProjectList {
			fmt.Printf("- %v\n", project.GetFullInformation())
		}
	}

	groupService := services.NewGroupService(utils.GetEnvironment())
	groupList, err := groupService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range *groupList {

		applicationProjectList, err := applicationProjectService.ListByFullNameWorkspace(fmt.Sprintf("%v/", group.Name), workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		if len(*applicationProjectList) > 0 {
			fmt.Println()
			fmt.Printf("%v/", group.Name)
			fmt.Println()

			for _, project := range *applicationProjectList {
				fmt.Printf("- %v\n", project.GetFullInformation())
			}
		}
	}

}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ListCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	ListCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := make([]string, 0)

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
