package list

import (
	"fmt"

	"parsdevkit.net/components/workspace"
	"parsdevkit.net/modules/template/code_template_contract"
	"parsdevkit.net/modules/template/file_template_contract"
	"parsdevkit.net/modules/template/shared_template_contract"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application/ioc"

	"github.com/spf13/cobra"
	"parsdevkit.net/application"
)

type ListOptions struct {
	Workspace string
}

var commandOptions ListOptions
var maxArgumentCount int = 0

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List template(s)",
	Long:    `List template(s)`,
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
	checkGlobals := _string.IsEmpty(commandOptions.Workspace)
	codeTemplateService := ioc.Get[code_template_contract.TemplateInterface]()
	fileTemplateService := ioc.Get[file_template_contract.TemplateInterface]()
	sharedTemplateService := ioc.Get[shared_template_contract.TemplateInterface]()

	if checkGlobals {
		fmt.Println("*** Global Templates ***")
		fmt.Println()

		commandOptions.Workspace = "None"

		sharedTemplateList, err := sharedTemplateService.ListByWorkspace(commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to list global shared templates\n%w", err)
		}

		fmt.Printf("(%d) shared template available\n\n", len(*sharedTemplateList))
		for _, template := range *sharedTemplateList {
			fmt.Printf("- %v\n", template.GetFullInformation())
		}

		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Println()

		codeTemplateList, err := codeTemplateService.ListByWorkspace(commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to list global code templates\n%w", err)
		}

		fmt.Printf("(%d) code template available\n\n", len(*codeTemplateList))
		for _, template := range *codeTemplateList {
			fmt.Printf("- %v\n", template.GetFullInformation())
		}

		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Println()

		fileTemplateList, err := fileTemplateService.ListByWorkspace(commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("Failed to list global file templates\n%w", err)
		}

		fmt.Printf("(%d) file template available\n\n", len(*fileTemplateList))
		for _, template := range *fileTemplateList {
			fmt.Printf("- %v\n", template.GetFullInformation())
		}

		commandOptions.Workspace = ""
		fmt.Println()
	}

	fmt.Println("*** Workspace Specific Templates ***")
	fmt.Println()

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	commandOptions.Workspace = workspace.GetActiveWorkspaceName(appCtx, commandOptions.Workspace)

	sharedTemplateList, err := sharedTemplateService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to list Active Workspace Shared templates\n%w", err)
	}

	fmt.Printf("(%d) shared template available\n\n", len(*sharedTemplateList))
	for _, template := range *sharedTemplateList {
		fmt.Printf("- %v\n", template.GetFullInformation())
	}

	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()

	codeTemplateList, err := codeTemplateService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to list Active Workspace Code templates\n%w", err)
	}

	fmt.Printf("(%d) code template available\n\n", len(*codeTemplateList))
	for _, template := range *codeTemplateList {
		fmt.Printf("- %v\n", template.GetFullInformation())
	}

	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()

	fileTemplateList, err := fileTemplateService.ListByWorkspace(commandOptions.Workspace)
	if err != nil {
		return fmt.Errorf("Failed to list Active Workspace File templates\n%w", err)
	}

	fmt.Printf("(%d) file template available\n\n", len(*fileTemplateList))
	for _, template := range *fileTemplateList {
		fmt.Printf("- %v\n", template.GetFullInformation())
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ListOptions{}
}
