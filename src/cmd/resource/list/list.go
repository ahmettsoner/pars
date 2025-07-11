package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/resource"
	"parsdevkit.net/pkg/utilities/json"

	"parsdevkit.net/application"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"
)

type ResourceListOptions struct {
}

var commandOptions ResourceListOptions
var maxArgumentCount int = 0

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List resource(s)",
	Long:    `List resource(s)`,
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

	result = append(result, &data_resource_payload_structs.ResourceBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			data_resource_payload_structs.RESOURCE_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: data_resource_payload_structs.ResourceSpecification{
			ResourceIdentifier: resource.ResourceIdentifier{},
		},
	})

	result = append(result, &object_resource_payload_structs.ResourceBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Resource,
			object_resource_payload_structs.RESOURCE_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: object_resource_payload_structs.ResourceSpecification{
			ResourceIdentifier: resource.ResourceIdentifier{},
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
	commandOptions = ResourceListOptions{}
}
