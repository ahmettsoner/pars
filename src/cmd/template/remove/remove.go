package remove

import (
	"fmt"
	"log"
	"os"
	"strings"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utilities"
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

	if len(commandOptions.Names) > 0 {

		checkGlobals := utilities.IsEmpty(commandOptions.Workspace)
		codeTemplateService := services.NewCodeTemplateService(utils.GetEnvironment())
		fileTemplateService := services.NewFileTemplateService(utils.GetEnvironment())
		sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())

		for _, name := range commandOptions.Names {

			if checkGlobals {

				commandOptions.Workspace = "None"

				ok, err := codeTemplateService.IsExists(name, commandOptions.Workspace)
				if err != nil {
					return fmt.Errorf("xxx: Code Template ('%s') kontrolünde hata oluştu\n%w", name, err)
				}
				if ok {
					codeTemplate, err := codeTemplateService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						return fmt.Errorf("Failed to remove Gblobal Code template(s) '%s'\n%w", name, err)
					}
					fmt.Println("Template (" + codeTemplate.Header.Name + ") deleted permanently")
				}

				ok, err = fileTemplateService.IsExists(name, commandOptions.Workspace)
				if err != nil {
					return fmt.Errorf("xxx: File Template ('%s') kontrolünde hata oluştu\n%w", name, err)
				}
				if ok {
					fileTemplate, err := fileTemplateService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						return fmt.Errorf("Failed to remove Global File template(s) '%s'\n%w", name, err)
					}
					fmt.Println("Template (" + fileTemplate.Header.Name + ") deleted permanently")
				}

				ok, err = sharedTemplateService.IsExists(name, commandOptions.Workspace)
				if err != nil {
					return fmt.Errorf("xxx: Shared Template ('%s') kontrolünde hata oluştu\n%w", name, err)
				}
				if ok {
					sharedTemplate, err := sharedTemplateService.Remove(name, commandOptions.Workspace, true)
					if err != nil {
						return fmt.Errorf("Failed to remove Global Shared template(s) '%s'\n%w", name, err)
					}
					fmt.Println("Template (" + sharedTemplate.Header.Name + ") deleted permanently")
				}

				commandOptions.Workspace = ""
			}
			appCtx := application.GetContext()
			if appCtx == nil {
				return fmt.Errorf("xxx: Current workspace bulunamadı")
			}

			commandOptions.Workspace = parsCMDCommon.GetActiveWorkspaceName(appCtx, commandOptions.Workspace)

			ok, err := codeTemplateService.IsExists(name, commandOptions.Workspace)
			if err != nil {
				return fmt.Errorf("xxx: Code Template ('%s') kontrolünde hata oluştu\n%w", name, err)
			}
			if ok {
				codeTemplate, err := codeTemplateService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					return fmt.Errorf("Failed to remove Active Workspace Code template(s) '%s'\n%w", name, err)
				}
				fmt.Println("Template (" + codeTemplate.Header.Name + ") deleted permanently")
			}

			ok, err = fileTemplateService.IsExists(name, commandOptions.Workspace)
			if err != nil {
				return fmt.Errorf("xxx: File Template ('%s') kontrolünde hata oluştu\n%w", name, err)
			}
			if ok {
				fileTemplate, err := fileTemplateService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					return fmt.Errorf("Failed to remove Active Workspace File template(s) '%s'\n%w", name, err)
				}
				fmt.Println("Template (" + fileTemplate.Header.Name + ") deleted permanently")
			}

			ok, err = sharedTemplateService.IsExists(name, commandOptions.Workspace)
			if err != nil {
				return fmt.Errorf("xxx: Shared Template ('%s') kontrolünde hata oluştu\n%w", name, err)
			}
			if ok {
				sharedTemplate, err := sharedTemplateService.Remove(name, commandOptions.Workspace, true)
				if err != nil {
					return fmt.Errorf("Failed to remove Active Workspace Shared template(s) '%s'\n%w", name, err)
				}
				fmt.Println("Template (" + sharedTemplate.Header.Name + ") deleted permanently")
			}
		}
		fmt.Fprintf(os.Stdout, "✔ template(s) '%v' removed successfully\n", commandOptions.Names)
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
	sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())
	sharedTemplateList, err := sharedTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *sharedTemplateList {
		if !utilities.Contains(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}

	fileTemplateService := services.NewFileTemplateService(utils.GetEnvironment())
	fileTemplateList, err := fileTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *fileTemplateList {
		if !utilities.Contains(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}

	codeTemplateService := services.NewCodeTemplateService(utils.GetEnvironment())
	codeTemplateList, err := codeTemplateService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, resource := range *codeTemplateList {
		if !utilities.Contains(args, resource.Header.Name) && strings.HasPrefix(resource.Header.Name, toComplete) {
			suggestions = append(suggestions, resource.Header.Name)
		}
	}

	return suggestions
}
