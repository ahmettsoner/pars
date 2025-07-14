package remove

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/application/engines"

	"parsdevkit.net/modules/template/code_template_contract"
	code_template_payload_structs "parsdevkit.net/modules/template/code_template_payload/structs"
	"parsdevkit.net/modules/template/file_template_contract"
	file_template_payload_structs "parsdevkit.net/modules/template/file_template_payload/structs"
	"parsdevkit.net/modules/template/shared_template_contract"
	shared_template_payload_structs "parsdevkit.net/modules/template/shared_template_payload/structs"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/pkg/utilities/array"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
)

type RemoveOptions struct {
	Names []string
	Force string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove name [name]...",
	Aliases:           []string{"r"},
	Short:             "Template Information",
	Long:              `Template Information`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: template name is required.")
	}

	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) > 0 {
		commandOptions.Names = args
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	var result []string = []string{code_template_payload_structs.MODULE_KEY, file_template_payload_structs.MODULE_KEY, shared_template_payload_structs.MODULE_KEY}

	for _, name := range commandOptions.Names {
		appCtx := application.GetContext()

		err := engines.DispatchEngineRemove(appCtx, result, name)
		if err != nil {
			return fmt.Errorf("Engine processing failed: %v", err)
		}
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = RemoveOptions{}
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listTemplateNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	addSubCommands()
}

func addSubCommands() {
}

func listTemplateNameSuggestions(args []string, toComplete string) []string {

	var suggestions = make([]string, 0)
	codeTemplateService := ioc.Get[code_template_contract.TemplateInterface]()
	fileTemplateService := ioc.Get[file_template_contract.TemplateInterface]()
	sharedTemplateService := ioc.Get[shared_template_contract.TemplateInterface]()

	sharedTemplateList, err := sharedTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *sharedTemplateList {
		if !array.ContainsSlice(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}

	fileTemplateList, err := fileTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *fileTemplateList {
		if !array.ContainsSlice(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}

	codeTemplateList, err := codeTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *codeTemplateList {
		if !array.ContainsSlice(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}

	return suggestions
}
