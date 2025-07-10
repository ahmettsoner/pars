package clean

import (
	"fmt"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/project"
	"parsdevkit.net/components/workspace"
	"parsdevkit.net/pkg/utilities/json"
	_string "parsdevkit.net/pkg/utilities/string"

	"github.com/spf13/cobra"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"

	"parsdevkit.net/application"
)

type CleanOptions struct {
	Name      string
	Workspace string
}

var commandOptions CleanOptions
var maxArgumentCount int = 1

var CleanCmd = &cobra.Command{
	Use:     "clean [name]",
	Aliases: []string{"c"},
	Short:   "Clean project(s)",
	Long:    `Clean project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("Please provide project name using the --url flag or as argument.")
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

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	var workspaceName, err = workspace.GetActiveWorkspaceNameV2(appCtx, commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("failed to find active workspace '%s'\n%w", commandOptions.Name, err)
	}
	commandOptions.Workspace = workspaceName

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	fmt.Printf("\n\n════════════════════════════════════\n")
	fmt.Printf("🔍 Cleaning: %s\n", _string.Concat(", ", commandOptions.Name))
	fmt.Printf("═══════════════════════════════════\n\n")

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

	var loadedSchemas []string = make([]string, 0)
	for _, data := range result {

		if err := data.Validate(); err != nil {
			jsonObject, _ := json.ToJson(data)
			return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
		}

		loadedSchemas = append(loadedSchemas, fmt.Sprintf("%s.%s", data.GetHeader().Name, data.GetKey()))
	}
	fmt.Printf("✅ Loaded Schemas: %s\n", _string.Concat(", ", loadedSchemas...))
	fmt.Printf("────────────────────────────────────\n")

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	err := engines.DispatchEngineClean(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = CleanOptions{}
}

func init() {
	CleanCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	CleanCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
