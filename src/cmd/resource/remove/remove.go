package remove

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"parsdevkit.net/engines/dataResource"
	"parsdevkit.net/engines/objectResource"
	"parsdevkit.net/operation/services"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type RemoveOptions struct {
	Names     []string
	Workspace string
	Force     string
	FilePaths []string
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
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(commandOptions.Names) == 0 && len(args) == 0 {
		return fmt.Errorf("error: project name is required. Provide it with '--name' or as an argument.")
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

				if objectResourceService.IsExists(name, commandOptions.Workspace) {
					objectResource, err := objectResourceService.Remove(name, commandOptions.Workspace, true, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Resource (" + objectResource.Name + ") deleted permanently")
				}

				if dataResourceService.IsExists(name, commandOptions.Workspace) {
					dataResource, err := dataResourceService.Remove(name, commandOptions.Workspace, true, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Resource (" + dataResource.Name + ") deleted permanently")
				}

				commandOptions.Workspace = ""
			}

			commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName(commandOptions.Workspace)

			if objectResourceService.IsExists(name, commandOptions.Workspace) {
				objectResource, err := objectResourceService.Remove(name, commandOptions.Workspace, true, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Resource (" + objectResource.Name + ") deleted permanently")
			}

			if dataResourceService.IsExists(name, commandOptions.Workspace) {
				dataResource, err := dataResourceService.Remove(name, commandOptions.Workspace, true, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Resource (" + dataResource.Name + ") deleted permanently")
			}
		}
	} else if len(commandOptions.FilePaths) > 0 {
		objectResourceService := objectResource.ObjectResourceEngine{}
		if err := objectResourceService.RemoveResourcesFromFile(true, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}

		dataResourceService := dataResource.DataResourceEngine{}
		if err := dataResourceService.RemoveResourcesFromFile(true, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please provide a name for the resource")
		os.Exit(1)
	}

	return nil
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

	RemoveCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "file", "f", nil, "Comma-separated list of declaration files")
	RemoveCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
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

func fileFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	files, _ := filepath.Glob(filepath.Join(toComplete, "*"))
	completions := []string{}
	for _, file := range files {
		if info, err := os.Stat(file); err == nil && !info.IsDir() {
			completions = append(completions, file)
		}
	}
	return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
}
