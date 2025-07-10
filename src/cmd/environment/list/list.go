package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/environment"
	"parsdevkit.net/pkg/utilities/json"

	"parsdevkit.net/application"
	basic_environment_payload_structs "parsdevkit.net/modules/environment/basic_environment_payload/structs"
)

type EnvironmentListOptions struct {
}

var commandOptions EnvironmentListOptions

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List environment(s)",
	Long:    `List environment(s)`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	fmt.Printf("════════════════════════════════════\n")
	fmt.Printf("⏳ Listing Environments \n")
	fmt.Printf("═══════════════════════════════════\n\n")

	var result []schemas.SchemaInterface = make([]schemas.SchemaInterface, 0)

	result = append(result, &basic_environment_payload_structs.EnvironmentBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Environment,
			basic_environment_payload_structs.ENVIRONMENT_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: basic_environment_payload_structs.EnvironmentSpecification{
			EnvironmentIdentifier: environment.EnvironmentIdentifier{},
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
	commandOptions = EnvironmentListOptions{}
}
