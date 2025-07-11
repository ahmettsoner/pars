package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/task"
	basic_task_payload_structs "parsdevkit.net/modules/task/basic_task_payload/structs"
	"parsdevkit.net/pkg/utilities/json"
)

type ListOptions struct {
	Workspace string
}

var commandOptions ListOptions
var maxArgumentCount int = 0

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List task(s)",
	Long:    `List task(s)`,
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

	result = append(result, &basic_task_payload_structs.TaskBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Task,
			basic_task_payload_structs.TASK_KIND,
			"temp-obj",
			schemas.Metadata{},
		),
		Specifications: basic_task_payload_structs.TaskSpecification{
			TaskIdentifier: task.TaskIdentifier{},
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
	commandOptions = ListOptions{}
}
