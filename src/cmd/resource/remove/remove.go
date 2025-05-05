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

var (
	names         []string
	workspaceName string
	force         string
	filePaths     []string
)

var RemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Short:   "Resource Information",
	Long:    `Resource Information`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {

	if len(args) > 0 {
		names = args

		checkGlobals := utils.IsEmpty(workspaceName)

		objectResourceService := services.NewObjectResourceService(utils.GetEnvironment())
		dataResourceService := services.NewDataResourceService(utils.GetEnvironment())

		for _, name := range names {
			if checkGlobals {
				workspaceName = "None"

				if objectResourceService.IsExists(name, workspaceName) {
					objectResource, err := objectResourceService.Remove(name, workspaceName, true, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Resource (" + objectResource.Name + ") deleted permanently")
				}

				if dataResourceService.IsExists(name, workspaceName) {
					dataResource, err := dataResourceService.Remove(name, workspaceName, true, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Resource (" + dataResource.Name + ") deleted permanently")
				}

				workspaceName = ""
			}

			workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

			if objectResourceService.IsExists(name, workspaceName) {
				objectResource, err := objectResourceService.Remove(name, workspaceName, true, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Resource (" + objectResource.Name + ") deleted permanently")
			}

			if dataResourceService.IsExists(name, workspaceName) {
				dataResource, err := dataResourceService.Remove(name, workspaceName, true, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Resource (" + dataResource.Name + ") deleted permanently")
			}
		}
	} else if len(filePaths) > 0 {
		objectResourceService := objectResource.ObjectResourceEngine{}
		if err := objectResourceService.RemoveResourcesFromFile(true, filePaths...); err != nil {
			log.Fatal(err)
		}

		dataResourceService := dataResource.DataResourceEngine{}
		if err := dataResourceService.RemoveResourcesFromFile(true, filePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please provide a name for the resource")
		os.Exit(1)
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {

	RemoveCmd.Flags().StringSliceVarP(&filePaths, "file", "f", nil, "Comma-separated list of declaration files")
	RemoveCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	suggestions := listResourceNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
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
