package describe

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"parsdevkit.net/core"
	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/engines"

	"github.com/spf13/cobra"
)

var (
	name                              string
	workspaceDescribeViewTypeEnumFlag core.WorkspaceDescribeViewTypeEnumFlag
	pathOnly                          bool
)
var maxArgumentCount int = 1

var DescribeCmd = &cobra.Command{
	Use:     "describe [name]",
	Aliases: []string{"d"},
	Short:   "Information about workspace",
	Long:    `Information about workspace`,
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

	appContext := engines.GetContext()

	if appContext.CurrentWorkspace == nil {
		fmt.Println("* You have to set current workspace")
	} else {
		if utils.IsEmpty(name) {
			if appContext != nil {
				name = appContext.CurrentWorkspace.Name
			} else {
				log.Fatal("Workspace cannot be accessable")
			}
		}
	}

	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspace, err := workspaceService.GetByName(name)
	if err != nil {
		log.Fatal(err)
	}

	if workspace == nil {
		fmt.Println("There are no workspace yet...")
	} else {
		projectService := services.NewApplicationProjectService(utils.GetEnvironment())
		projectList, err := projectService.ListByWorkspace(workspace.Specifications.Name)
		if err != nil {
			log.Fatal(err)
		}

		if pathOnly {
			fmt.Print(workspace.Specifications.Path)
			return
		}

		fmt.Printf("Workspace (%v) has %d project\n", workspace.Name, len(*projectList))
		fmt.Printf("Path : %v \n", workspace.Specifications.Path)

		fmt.Printf("\nProjects:\n")
		if workspaceDescribeViewTypeEnumFlag.Value == "flat" {
			for _, e := range *projectList {
				name := fmt.Sprintf(" - %v", e.GetFullInformation())
				fmt.Println(name)
			}
		} else if workspaceDescribeViewTypeEnumFlag.Value == "hierarchical" {
			groups := make(map[string][]string)
			keys := []string{}

			for _, e := range *projectList {
				name := e.GetInformation()
				if !utils.IsEmpty(e.Specifications.GroupObject.Name) {
					groups[e.Specifications.Group] = append(groups[e.Specifications.Group], name)
				} else {
					groups[name] = []string{}
				}
			}

			for key := range groups {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			for _, key := range keys {
				groupItems := groups[key]

				if len(groupItems) > 0 {
					fmt.Printf("%s\n", key)
					for _, value := range groupItems {
						fmt.Printf("  - %s\n", value)
					}
				} else {
					fmt.Printf("- %s\n", key)
				}
			}
		}
	}
}

func init() {
	workspaceDescribeViewTypeValues := core.WorkspaceDescribeViewTypeToArray()
	workspaceDescribeViewTypeEnumFlag.Value = core.WorkspaceDescribeViewTypes.Hierarchical
	DescribeCmd.Flags().VarP(&workspaceDescribeViewTypeEnumFlag, "view", "v", fmt.Sprintf("Select view type %v", workspaceDescribeViewTypeValues))
	DescribeCmd.RegisterFlagCompletionFunc("view", viewTypeFlagCompletion)

	DescribeCmd.Flags().BoolVarP(&pathOnly, "path", "p", false, "Show path only")
}
func viewTypeFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var suggestions = make([]string, 0)

	workspaceDescribeViewTypeValues := core.WorkspaceDescribeViewTypeToArray()

	for _, _type := range workspaceDescribeViewTypeValues {
		suggestions = append(suggestions, string(_type))
	}

	return suggestions, cobra.ShellCompDirectiveNoSpace
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) < maxArgumentCount {

		if len(args) == 0 {
			suggestions := listWorkspaceNameSuggestions(args, toComplete)

			return suggestions, cobra.ShellCompDirectiveNoSpace
		}
	}

	return make([]string, 0), cobra.ShellCompDirectiveNoFileComp
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
