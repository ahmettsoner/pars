package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/template"
	"parsdevkit.net/pkg/utilities/json"

	"parsdevkit.net/application"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"
)

type TemplateListOptions struct {
}

var commandOptions TemplateListOptions
var maxArgumentCount int = 0

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List template(s)",
	Long:    `List template(s)`,
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

	err := engines.DispatchEngineList(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = TemplateListOptions{}
}
