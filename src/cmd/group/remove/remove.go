package remove

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"parsdevkit.net/engines/group"
	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var (
	names []string
	// workspaceName string
	force     string
	filePaths []string
)

var RemoveCmd = &cobra.Command{
	Use:     "remove name [name]...",
	Aliases: []string{"r"},
	Short:   "Group Information",
	Long:    `Group Information`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {

	if len(args) > 0 {
		names = args

		groupService := services.NewGroupService(utils.GetEnvironment())
		for _, name := range names {
			group, err := groupService.Remove(name, true)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Group (" + group.Name + ") deleted permanently")
		}
	} else if len(filePaths) > 0 {
		groupService := group.GroupEngine{}
		if err := groupService.RemoveGroupsFromFile(true, filePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please provide a name for the group")
		os.Exit(1)
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().StringSliceVarP(&filePaths, "file", "f", nil, "Comma-separated list of declaration files")
	RemoveCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listGroupNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
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
