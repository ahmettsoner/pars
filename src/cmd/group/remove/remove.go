package remove

import (
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/application/structs/project"
	"parsdevkit.net/modules/group/basic_group_contract"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"log"
	"strings"

	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	"parsdevkit.net/application"

	"parsdevkit.net/application/engines"

	"parsdevkit.net/pkg/utilities/array"
	_string "parsdevkit.net/pkg/utilities/string"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/schemas"
)

type RemoveOptions struct {
	Name      string
	Workspace string
	Force     string
}

var commandOptions RemoveOptions
var maxArgumentCount int = 1

var RemoveCmd = &cobra.Command{
	Use:               "remove [name]",
	Aliases:           []string{"r"},
	Short:             "Remove group",
	Long:              `Remove group`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: group name is required. Provide as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	var result []schemas.SchemaInterface = make([]schemas.SchemaInterface, 0)

	result = append(result, &application_project_payload_structs.ProjectBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Project,
			application_project_payload_structs.PROJECT_KIND,
			commandOptions.Name,
			schemas.Metadata{},
		),
		Specifications: application_project_payload_structs.ProjectSpecification{
			ProjectIdentifier: project.ProjectIdentifier{
				Workspace: commandOptions.Workspace,
			},
		},
	})

	appCtx := application.GetContext()

	err := engines.DispatchEngineDestroy(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
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
	commandOptions = RemoveOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	// RemoveCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	// RemoveCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)
}
func listGroupNameSuggestions(args []string, toComplete string) []string {

	// workspaceName = workspace.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	groupService := ioc.Get[basic_group_contract.GroupInterface]()
	groupList, err := groupService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range *groupList {
		if !array.ContainsSlice(args, group.Header.Name) && strings.HasPrefix(group.Header.Name, toComplete) {
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
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, workspace := range *workspaceList {
		if !array.ContainsSlice(args, workspace.Header.Name) && strings.HasPrefix(workspace.Header.Name, toComplete) {
			suggestions = append(suggestions, workspace.Header.Name)
		}
	}
	return suggestions
}
