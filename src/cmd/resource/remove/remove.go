package remove

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/application/engines"
	"parsdevkit.net/modules/resource/data_resource_contract"
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"
	"parsdevkit.net/modules/resource/object_resource_contract"
	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/application/ioc"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
	"parsdevkit.net/pkg/utilities/array"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove name [name]...",
	Aliases:           []string{"r"},
	Short:             "Resource Information",
	Long:              `Resource Information`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: resource name is required.")
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

	var result []string = []string{object_resource_payload_structs.MODULE_KEY, data_resource_payload_structs.MODULE_KEY}

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
	suggestions := listResourceNameSuggestions(args, toComplete)

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

func listResourceNameSuggestions(args []string, toComplete string) []string {

	// workspaceName = workspace.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	objectResourceService := ioc.Get[object_resource_contract.ResourceInterface]()
	dataResourceService := ioc.Get[data_resource_contract.ResourceInterface]()
	objectResourceList, err := objectResourceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *objectResourceList {
		if !array.ContainsSlice(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}

	dataResourceList, err := dataResourceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *dataResourceList {
		if !array.ContainsSlice(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}
	return suggestions
}
