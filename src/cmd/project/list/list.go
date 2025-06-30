package list

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/operation/services"

	parsCMDCommon "parsdevkit.net/core/cmd"
	group "parsdevkit.net/modules/group/group"

	"parsdevkit.net/application"
	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
	platformsCommon "parsdevkit.net/platforms/common"
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
	if utils.IsEmpty(commandOptions.Workspace) {
		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName(appCtx, "")
	}
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	applicationProjectService := services.NewApplicationProjectService(utils.GetEnvironment(), platformsCommon.Registry)
	applicationProjectList, err := applicationProjectService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to listing projects\n%w", err)
	}
	fmt.Printf("(%d) application project available\n", len(*applicationProjectList))

	applicationProjectListBasic, err := applicationProjectService.ListIndividualByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to list Workspace Application projects\n%w", err)
	}

	if len(*applicationProjectListBasic) > 0 {
		fmt.Println()
		for _, project := range *applicationProjectList {
			fmt.Printf("- %v\n", project.GetFullInformation())
		}
	}

	groupService := group.NewGroupService(utils.GetEnvironment())
	groupList, err := groupService.List()
	if err != nil {
		return fmt.Errorf("Failed to list projects groups\n%w", err)
	}

	for _, group := range *groupList {

		applicationProjectList, err := applicationProjectService.ListByFullNameWorkspace(fmt.Sprintf("%v/", group.Header.Name), commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to list Group '%s' projects\n%w", group.Header.Name, err)
		}

		if len(*applicationProjectList) > 0 {
			fmt.Println()
			fmt.Printf("%v/", group.Header.Name)
			fmt.Println()

			for _, project := range *applicationProjectList {
				fmt.Printf("- %v\n", project.GetFullInformation())
			}
		}
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
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, workspace := range *workspaceList {
		if !utils.Contains(args, workspace.Header.Name) && strings.HasPrefix(workspace.Header.Name, toComplete) {
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
