package describe

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
	group "parsdevkit.net/modules/group/group"
)

type DescribeOptions struct {
	Name      string
	Workspace string
	Force     string
}

var commandOptions DescribeOptions
var maxArgumentCount int = 1

var DescribeCmd = &cobra.Command{
	Use:               "describe [name]",
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
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	groupService := group.NewGroupService(utils.GetEnvironment())
	group, err := groupService.GetByName(commandOptions.Name)
	if err != nil {
		return fmt.Errorf("Failed to retrieve group '%s'\n%w", commandOptions.Name, err)
	}

	name := fmt.Sprintf("Group Name:\t%v", group.Header.Name)
	fmt.Println(name)

	path := fmt.Sprintf("Path:\t\t%v", group.Specifications.Path)
	fmt.Println(path)

	packageName := fmt.Sprintf("Package:\t%v", group.Specifications.GetPackageString())
	fmt.Println(packageName)

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	projectList, err := projectService.ListByGroupName(group.Header.Name)
	if err != nil {
		return fmt.Errorf("Failed to retrieve group projects '%s'\n%w", commandOptions.Name, err)
	}

	fmt.Printf("Projects:\n")
	for _, e := range *projectList {
		name := fmt.Sprintf("\t - %v", e.GetFullInformation())
		fmt.Println(name)
	}

	return nil
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
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = DescribeOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	// DescribeCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	// DescribeCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}
func listGroupNameSuggestions(args []string, toComplete string) []string {

	// workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	groupService := group.NewGroupService(utils.GetEnvironment())
	groupList, err := groupService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range *groupList {
		if !utils.Contains(args, group.Header.Name) && strings.HasPrefix(group.Header.Name, toComplete) {
			suggestions = append(suggestions, group.Header.Name)
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
		if !utils.Contains(args, workspace.Header.Name) && strings.HasPrefix(workspace.Header.Name, toComplete) {
			suggestions = append(suggestions, workspace.Header.Name)
		}
	}
	return suggestions
}
