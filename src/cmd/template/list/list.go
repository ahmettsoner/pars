package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"parsdevkit.net/application/engines"

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

	var result []string = []string{
		code_template_payload_structs.MODULE_KEY,
		file_template_payload_structs.MODULE_KEY,
		shared_template_payload_structs.MODULE_KEY,
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
