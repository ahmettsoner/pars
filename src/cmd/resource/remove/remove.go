package remove

import (
	"fmt"
	"log"
	"strings"

	"parsdevkit.net/operation/services"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
}

var commandOptions RemoveOptions

var RemoveCmd = &cobra.Command{
	Use:               "remove",
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

	if len(commandOptions.Names) > 0 {

		checkGlobals := utils.IsEmpty(commandOptions.Workspace)

		objectResourceService := services.NewObjectResourceService(utils.GetEnvironment())
		dataResourceService := services.NewDataResourceService(utils.GetEnvironment())

		for _, name := range commandOptions.Names {
			if checkGlobals {
				commandOptions.Workspace = "None"

				ok, err := objectResourceService.IsExists(name, commandOptions.Workspace)
				if err != nil {
					return fmt.Errorf("xxx: Object Resource ('%s') kontrolünde hata oluştu\n%w", name, err)
				}
				if ok {
					objectResource, err := objectResourceService.Remove(name, commandOptions.Workspace, true, true)
					if err != nil {
						return fmt.Errorf("Failed to remove object resource(s) '%s'\n%w", commandOptions.Names, err)
					}
					fmt.Println("Resource (" + objectResource.Name + ") deleted permanently")
				}

				ok, err = dataResourceService.IsExists(name, commandOptions.Workspace)
				if err != nil {
					return fmt.Errorf("xxx: Data Resource ('%s') kontrolünde hata oluştu\n%w", name, err)
				}
				if ok {
					dataResource, err := dataResourceService.Remove(name, commandOptions.Workspace, true, true)
					if err != nil {
						return fmt.Errorf("Failed to remove object resource(s) '%s'\n%w", name, err)
					}
					fmt.Println("Resource (" + dataResource.Name + ") deleted permanently")
				}

				commandOptions.Workspace = ""
			}

			commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName(application.GetContext(), commandOptions.Workspace)

			ok, err := objectResourceService.IsExists(name, commandOptions.Workspace)
			if err != nil {
				return fmt.Errorf("xxx: Object Resource ('%s') kontrolünde hata oluştu\n%w", name, err)
			}
			if ok {
				objectResource, err := objectResourceService.Remove(name, commandOptions.Workspace, true, true)
				if err != nil {
					return fmt.Errorf("Failed to remove data resource(s) '%s'\n%w", commandOptions.Names, err)
				}
				fmt.Println("Resource (" + objectResource.Name + ") deleted permanently")
			}

			ok, err = dataResourceService.IsExists(name, commandOptions.Workspace)
			if err != nil {
				return fmt.Errorf("xxx: Data Resource ('%s') kontrolünde hata oluştu\n%w", name, err)
			}
			if ok {
				dataResource, err := dataResourceService.Remove(name, commandOptions.Workspace, true, true)
				if err != nil {
					return fmt.Errorf("Failed to remove data resource(s) '%s'\n%w", name, err)
				}
				fmt.Println("Resource (" + dataResource.Name + ") deleted permanently")
			}
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

	// workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	objectResourceService := services.NewObjectResourceService(utils.GetEnvironment())
	objectResourceList, err := objectResourceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *objectResourceList {
		if !utils.Contains(args, resource.Name) && strings.HasPrefix(resource.Name, toComplete) {
			suggestions = append(suggestions, resource.Name)
		}
	}

	dataResourceService := services.NewDataResourceService(utils.GetEnvironment())
	dataResourceList, err := dataResourceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *dataResourceList {
		if !utils.Contains(args, resource.Name) && strings.HasPrefix(resource.Name, toComplete) {
			suggestions = append(suggestions, resource.Name)
		}
	}
	return suggestions
}
