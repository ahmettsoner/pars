package remove

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	parsCMDCommon "parsdevkit.net/core/cmd"
	"parsdevkit.net/engines/codeTemplate"
	"parsdevkit.net/engines/fileTemplate"
	"parsdevkit.net/engines/sharedTemplate"

	"parsdevkit.net/operation/services"

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
	Short:   "Template Information",
	Long:    `Template Information`,
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
		codeTemplateService := services.NewCodeTemplateService(utils.GetEnvironment())
		fileTemplateService := services.NewFileTemplateService(utils.GetEnvironment())
		sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())

		for _, name := range names {

			if checkGlobals {

				workspaceName = "None"

				if codeTemplateService.IsExists(name, workspaceName) {
					codeTemplate, err := codeTemplateService.Remove(name, workspaceName, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Template (" + codeTemplate.Name + ") deleted permanently")
				}

				if fileTemplateService.IsExists(name, workspaceName) {
					fileTemplate, err := fileTemplateService.Remove(name, workspaceName, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Template (" + fileTemplate.Name + ") deleted permanently")
				}

				if sharedTemplateService.IsExists(name, workspaceName) {
					sharedTemplate, err := sharedTemplateService.Remove(name, workspaceName, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Template (" + sharedTemplate.Name + ") deleted permanently")
				}

				workspaceName = ""
			}

			workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

			if codeTemplateService.IsExists(name, workspaceName) {
				codeTemplate, err := codeTemplateService.Remove(name, workspaceName, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Template (" + codeTemplate.Name + ") deleted permanently")
			}

			if fileTemplateService.IsExists(name, workspaceName) {
				fileTemplate, err := fileTemplateService.Remove(name, workspaceName, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Template (" + fileTemplate.Name + ") deleted permanently")
			}

			if sharedTemplateService.IsExists(name, workspaceName) {
				sharedTemplate, err := sharedTemplateService.Remove(name, workspaceName, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Template (" + sharedTemplate.Name + ") deleted permanently")
			}
		}
	} else if len(filePaths) > 0 {
		sharedTemplateService := sharedTemplate.SharedTemplateEngine{}
		if err := sharedTemplateService.RemoveTemplatesFromFile(true, filePaths...); err != nil {
			log.Fatal(err)
		}

		codeTemplateService := codeTemplate.CodeTemplateEngine{}
		if err := codeTemplateService.RemoveTemplatesFromFile(true, filePaths...); err != nil {
			log.Fatal(err)
		}

		fileTemplateService := fileTemplate.FileTemplateEngine{}
		if err := fileTemplateService.RemoveTemplatesFromFile(true, filePaths...); err != nil {
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
	suggestions := listTemplateNameSuggestions(args, toComplete)

	if len(suggestions) == 1 {
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}

func listTemplateNameSuggestions(args []string, toComplete string) []string {

	// workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	var suggestions = make([]string, 0)
	sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())
	sharedTemplateList, err := sharedTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *sharedTemplateList {
		if !utils.Contains(args, resource.Name) && strings.HasPrefix(resource.Name, toComplete) {
			suggestions = append(suggestions, resource.Name)
		}
	}

	fileTemplateService := services.NewFileTemplateService(utils.GetEnvironment())
	fileTemplateList, err := fileTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *fileTemplateList {
		if !utils.Contains(args, resource.Name) && strings.HasPrefix(resource.Name, toComplete) {
			suggestions = append(suggestions, resource.Name)
		}
	}

	codeTemplateService := services.NewCodeTemplateService(utils.GetEnvironment())
	codeTemplateList, err := codeTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *codeTemplateList {
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
