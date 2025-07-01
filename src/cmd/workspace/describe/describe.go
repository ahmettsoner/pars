package describe

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"parsdevkit.net/application"
	"parsdevkit.net/core"
	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utilities"
	"parsdevkit.net/core/utils"
	platformsCommon "parsdevkit.net/platforms/common"

	"github.com/spf13/cobra"
)

type DescribeOptions struct {
	Name                  string
	WorkspaceDescribeView core.WorkspaceDescribeViewTypeEnumFlag
	PathOnly              bool
}

var commandOptions DescribeOptions
var maxArgumentCount int = 1

var DescribeCmd = &cobra.Command{
	Use:               "describe [name]",
	Aliases:           []string{"d"},
	Short:             "Information about workspace",
	Long:              `Information about workspace`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only group name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utilities.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}

	if &appCtx.CurrentWorkspace == nil {
		fmt.Println("* You have to set current workspace")
	} else {
		if utilities.IsEmpty(commandOptions.Name) {
			if appCtx != nil {
				commandOptions.Name = appCtx.CurrentWorkspace.Name
			} else {
				return fmt.Errorf("Workspace cannot be accessable")
			}
		}
	}

	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspace, err := workspaceService.GetByName(commandOptions.Name)
	if err != nil {
		return fmt.Errorf("Failed to retrieve workspace '%s'\n%w", commandOptions.Name, err)
	}

	if workspace == nil {
		fmt.Println("There are no workspace yet...")
	} else {
		projectService := services.NewApplicationProjectService(utils.GetEnvironment(), platformsCommon.Registry)
		projectList, err := projectService.ListByWorkspace(workspace.Specifications.Name)
		if err != nil {
			return fmt.Errorf("Failed to retrieve workspace projects '%s'\n%w", commandOptions.Name, err)
		}

		if commandOptions.PathOnly {
			fmt.Print(workspace.Specifications.Path)
		}

		fmt.Printf("Workspace (%v) has %d project\n", workspace.Header.Name, len(*projectList))
		fmt.Printf("Path : %v \n", workspace.Specifications.Path)

		fmt.Printf("\nProjects:\n")
		if commandOptions.WorkspaceDescribeView.Value == "flat" {
			for _, e := range *projectList {
				name := fmt.Sprintf(" - %v", e.GetFullInformation())
				fmt.Println(name)
			}
		} else if commandOptions.WorkspaceDescribeView.Value == "hierarchical" {
			groups := make(map[string][]string)
			keys := []string{}

			for _, e := range *projectList {
				name := e.GetInformation()
				if !utilities.IsEmpty(e.Specifications.GroupObject.Name) {
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

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = DescribeOptions{}
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

func init() {
	workspaceDescribeViewTypeValues := core.WorkspaceDescribeViewTypeToArray()
	commandOptions.WorkspaceDescribeView.Value = core.WorkspaceDescribeViewTypes.Hierarchical
	DescribeCmd.Flags().VarP(&commandOptions.WorkspaceDescribeView, "view", "v", fmt.Sprintf("Select view type %v", workspaceDescribeViewTypeValues))
	DescribeCmd.RegisterFlagCompletionFunc("view", viewTypeFlagCompletion)

	DescribeCmd.Flags().BoolVarP(&commandOptions.PathOnly, "path", "p", false, "Show path only")
}
func viewTypeFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var suggestions = make([]string, 0)

	workspaceDescribeViewTypeValues := core.WorkspaceDescribeViewTypeToArray()

	for _, _type := range workspaceDescribeViewTypeValues {
		suggestions = append(suggestions, string(_type))
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
		if !utilities.Contains(args, workspace.Header.Name) && strings.HasPrefix(workspace.Header.Name, toComplete) {
			suggestions = append(suggestions, workspace.Header.Name)
		}
	}
	return suggestions
}
