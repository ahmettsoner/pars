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
	Short:             "Template Information",
	Long:              `Template Information`,
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
		codeTemplateService := services.NewCodeTemplateService(utils.GetEnvironment())
		fileTemplateService := services.NewFileTemplateService(utils.GetEnvironment())
		sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())

		for _, name := range commandOptions.Names {

			if checkGlobals {

				commandOptions.Workspace = "None"

				if codeTemplateService.IsExists(name, commandOptions.Workspace) {
					codeTemplate, err := codeTemplateService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Template (" + codeTemplate.Name + ") deleted permanently")
				}

				if fileTemplateService.IsExists(name, commandOptions.Workspace) {
					fileTemplate, err := fileTemplateService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Template (" + fileTemplate.Name + ") deleted permanently")
				}

				if sharedTemplateService.IsExists(name, commandOptions.Workspace) {
					sharedTemplate, err := sharedTemplateService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Template (" + sharedTemplate.Name + ") deleted permanently")
				}

				commandOptions.Workspace = ""
			}

			commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName(commandOptions.Workspace)

			if codeTemplateService.IsExists(name, commandOptions.Workspace) {
				codeTemplate, err := codeTemplateService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Template (" + codeTemplate.Name + ") deleted permanently")
			}

			if fileTemplateService.IsExists(name, commandOptions.Workspace) {
				fileTemplate, err := fileTemplateService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Template (" + fileTemplate.Name + ") deleted permanently")
			}

			if sharedTemplateService.IsExists(name, commandOptions.Workspace) {
				sharedTemplate, err := sharedTemplateService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Template (" + sharedTemplate.Name + ") deleted permanently")
			}
		}
	} else if len(commandOptions.FilePaths) > 0 {
		sharedTemplateService := sharedTemplate.SharedTemplateEngine{}
		if err := sharedTemplateService.RemoveTemplatesFromFile(true, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}

		codeTemplateService := codeTemplate.CodeTemplateEngine{}
		if err := codeTemplateService.RemoveTemplatesFromFile(true, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}

		fileTemplateService := fileTemplate.FileTemplateEngine{}
		if err := fileTemplateService.RemoveTemplatesFromFile(true, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please provide a name for the resource")
		os.Exit(1)
	}

	return nil
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
	RemoveCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "file", "f", nil, "Comma-separated list of declaration files")
	RemoveCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
}

func listTemplateNameSuggestions(args []string, toComplete string) []string {

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
