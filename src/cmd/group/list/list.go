package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/group"
	"parsdevkit.net/pkg/utilities/json"

	"parsdevkit.net/application"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
)

type GroupListOptions struct {
}

var commandOptions GroupListOptions

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List group(s)",
	Long:    `List group(s)`,
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
	var result []schemas.SchemaInterface = make([]schemas.SchemaInterface, 0)

	result = append(result, &basic_group_payload_structs.GroupBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Group,
			basic_group_payload_structs.GROUP_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: basic_group_payload_structs.GroupSpecification{
			GroupIdentifier: group.GroupIdentifier{},
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
	commandOptions = GroupListOptions{}
}
