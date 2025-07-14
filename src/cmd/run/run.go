package run

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

type RunOptions struct {
	Names     []string
	Workspace string
}

var commandOptions RunOptions

var RunCmd = &cobra.Command{
	Use:     "run [name]...",
	Aliases: []string{"c"},
	Short:   "Run project(s)",
	Long:    `Run project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    runFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: project name is required.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) > 0 {
		commandOptions.Names = args
	}

	if _string.IsEmpty(commandOptions.Workspace) {
		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		commandOptions.Workspace = workspace.GetActiveWorkspaceName(appCtx, "")
	}

	return nil
}

func runFunc(cmd *cobra.Command, args []string) error {
	var result []schemas.SchemaInterface = make([]schemas.SchemaInterface, 0)

	if len(commandOptions.Names) > 0 {
		for _, name := range commandOptions.Names {
			result = append(result, &application_project_payload_structs.ProjectBaseStruct{
				Header: schemas.NewSchemaHeader(
					schemas.StructTypes.Project,
					application_project_payload_structs.PROJECT_KIND,
					name,
					schemas.Metadata{},
				),
				Specifications: application_project_payload_structs.ProjectSpecification{
					ProjectIdentifier: project.ProjectIdentifier{
						Workspace: commandOptions.Workspace,
					},
				},
			})
		}
	}

	var loadedSchemas []string = make([]string, 0)
	for _, data := range result {

		if err := data.Validate(); err != nil {
			jsonObject, _ := json.ToJson(data)
			return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
		}

		loadedSchemas = append(loadedSchemas, fmt.Sprintf("%s.%s", data.GetHeader().Name, data.GetKey()))
	}
	fmt.Printf("Loaded Schemas: %s\n", _string.Concat(", ", loadedSchemas...))
	fmt.Printf("────────────────────────────────────\n")

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	err := engines.DispatchEngineRun(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = RunOptions{}
}

func init() {
	RunCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
