package describe

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/application/ioc"

	"parsdevkit.net/application"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/workspace"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
	"parsdevkit.net/pkg/utilities/array"
	"parsdevkit.net/pkg/utilities/json"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"github.com/spf13/cobra"
)

type DescribeOptions struct {
	Name                  string
	WorkspaceDescribeView application.WorkspaceDescribeViewTypeEnumFlag
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
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}

	if &appCtx.CurrentWorkspace == nil {
		fmt.Println("* You have to set current workspace")
	} else {
		if _string.IsEmpty(commandOptions.Name) {
			if appCtx != nil {
				commandOptions.Name = appCtx.CurrentWorkspace.Name
			} else {
				return fmt.Errorf("Workspace cannot be accessable")
			}
		}
	}

	return nil
}
func executeFunc(cmd *cobra.Command, args []string) error {

	var result []schemas.SchemaInterface = make([]schemas.SchemaInterface, 0)

	result = append(result, &basic_workspace_payload_structs.WorkspaceBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Workspace,
			basic_workspace_payload_structs.WORKSPACE_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: basic_workspace_payload_structs.WorkspaceSpecification{
			WorkspaceIdentifier: workspace.WorkspaceIdentifier{},
		},
	})

	var loadedSchemas []string = make([]string, 0)
	for _, data := range result {

		if err := data.Validate(); err != nil {
			jsonObject, _ := json.ToJson(data)
			return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
		}

		loadedSchemas = append(loadedSchemas, fmt.Sprintf("%s.%s", data.GetHeader().Name, data.GetKey()))
	}

	appCtx := application.GetContext()

	err := engines.DispatchEngineDescribe(appCtx, result, commandOptions.Name)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
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
	workspaceDescribeViewTypeValues := application.WorkspaceDescribeViewTypeToArray()
	commandOptions.WorkspaceDescribeView.Value = application.WorkspaceDescribeViewTypes.Hierarchical
	DescribeCmd.Flags().VarP(&commandOptions.WorkspaceDescribeView, "view", "v", fmt.Sprintf("Select view type %v", workspaceDescribeViewTypeValues))
	DescribeCmd.RegisterFlagCompletionFunc("view", viewTypeFlagCompletion)

	DescribeCmd.Flags().BoolVarP(&commandOptions.PathOnly, "path", "p", false, "Show path only")
}
func viewTypeFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var suggestions = make([]string, 0)

	workspaceDescribeViewTypeValues := application.WorkspaceDescribeViewTypeToArray()

	for _, _type := range workspaceDescribeViewTypeValues {
		suggestions = append(suggestions, string(_type))
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
