package describe

import (
	"fmt"
	"log"
	"os"
	"strings"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var (
	name string
	// workspaceName string
	force string
)
var maxArgumentCount int = 1

var DescribeCmd = &cobra.Command{
	Use:     "describe [name]",
	Aliases: []string{"d"},
	Short:   "Information about project",
	Long:    `Information about project`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > maxArgumentCount {
			return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
		}
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		name = args[0]
	}

	if utils.IsEmpty(name) {
		cmd.Help()
		os.Exit(0)
	}

	// workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	groupService := services.NewGroupService(utils.GetEnvironment())
	group, err := groupService.GetByName(name)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("Group Name:\t%v", group.Name)
	fmt.Println(name)

	path := fmt.Sprintf("Path:\t\t%v", group.Specifications.Path)
	fmt.Println(path)

	packageName := fmt.Sprintf("Package:\t%v", group.Specifications.GetPackageString())
	fmt.Println(packageName)

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	projectList, err := projectService.ListByGroupName(group.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Projects:\n")
	for _, e := range *projectList {
		name := fmt.Sprintf("\t - %v", e.GetFullInformation())
		fmt.Println(name)
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	// DescribeCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	// DescribeCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) < maxArgumentCount {

		if len(args) == 0 {
			suggestions := listGroupNameSuggestions(args, toComplete)

			return suggestions, cobra.ShellCompDirectiveNoSpace
		}
	}

	return make([]string, 0), cobra.ShellCompDirectiveNoFileComp
}
func listGroupNameSuggestions(args []string, toComplete string) []string {

	// workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	groupService := services.NewGroupService(utils.GetEnvironment())
	groupList, err := groupService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range *groupList {
		if !utils.Contains(args, group.Name) && strings.HasPrefix(group.Name, toComplete) {
			suggestions = append(suggestions, group.Name)
		}
	}
	return suggestions
}

func workspaceFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var suggestions = make([]string, 0)

	workspaceList := listWorkspaceNameSuggestions(args, toComplete)

	for _, workspace := range workspaceList {
		suggestions = append(suggestions, workspace)
	}

	return suggestions, cobra.ShellCompDirectiveNoSpace
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
