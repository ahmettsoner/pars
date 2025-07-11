package describe

import (
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/template/code_template_contract"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"log"
	"strings"

	"parsdevkit.net/application"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/template"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"
	"parsdevkit.net/pkg/utilities/json"

	"parsdevkit.net/pkg/utilities/array"
	_string "parsdevkit.net/pkg/utilities/string"

	"github.com/spf13/cobra"
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
	Short:             "Information about template",
	Long:              `Information about template`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: template name is required. Provide as an argument.")
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

	result = append(result, &code_template_payload_structs.TemplateBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Template,
			code_template_payload_structs.TEMPLATE_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: code_template_payload_structs.TemplateSpecification{
			TemplateIdentifier: template.TemplateIdentifier{},
		},
	})

	result = append(result, &file_template_payload_structs.TemplateBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Template,
			file_template_payload_structs.TEMPLATE_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: file_template_payload_structs.TemplateSpecification{
			TemplateIdentifier: template.TemplateIdentifier{},
		},
	})

	result = append(result, &shared_template_payload_structs.TemplateBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Template,
			shared_template_payload_structs.TEMPLATE_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: shared_template_payload_structs.TemplateSpecification{
			TemplateIdentifier: template.TemplateIdentifier{},
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

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) < maxArgumentCount {

		if len(args) == 0 {
			suggestions := listTemplateNameSuggestions(args, toComplete)

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
func listTemplateNameSuggestions(args []string, toComplete string) []string {

	// workspaceName = workspace.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	templateService := ioc.Get[code_template_contract.TemplateInterface]()
	templateList, err := templateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, template := range *templateList {
		if !array.ContainsSlice(args, template.Header.Name) && strings.HasPrefix(template.Header.Name, toComplete) {
			suggestions = append(suggestions, template.Header.Name)
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
