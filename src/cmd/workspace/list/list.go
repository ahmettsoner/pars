package list

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/workspace"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
	"parsdevkit.net/pkg/utilities/json"

	"github.com/spf13/cobra"
)

var maxArgumentCount int = 0

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Example: `  pars workspace list [flags]
  pars wl [flags]`,
	Short:   "List workspace project(s)",
	Long:    `List workspace project(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("There is no argument supported")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {

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

	err := engines.DispatchEngineList(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
}
