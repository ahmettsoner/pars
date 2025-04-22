package list

import (
	"fmt"
	"log"

	"parsdevkit.net/operation/services"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

var (
	workspaceName string
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List template(s)",
	Long:    `List template(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	checkGlobals := utils.IsEmpty(workspaceName)
	sharedTemplateService := services.NewSharedTemplateService(utils.GetEnvironment())
	codeTemplateService := services.NewCodeTemplateService(utils.GetEnvironment())
	fileTemplateService := services.NewFileTemplateService(utils.GetEnvironment())

	if checkGlobals {
		fmt.Println("*** Global Templates ***")
		fmt.Println()

		workspaceName = "None"

		sharedTemplateList, err := sharedTemplateService.ListByWorkspace(workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("(%d) shared template available\n\n", len(*sharedTemplateList))
		for _, template := range *sharedTemplateList {
			fmt.Printf("- %v\n", template.GetFullInformation())
		}

		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Println()

		codeTemplateList, err := codeTemplateService.ListByWorkspace(workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("(%d) code template available\n\n", len(*codeTemplateList))
		for _, template := range *codeTemplateList {
			fmt.Printf("- %v\n", template.GetFullInformation())
		}

		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Println()

		fileTemplateList, err := fileTemplateService.ListByWorkspace(workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("(%d) file template available\n\n", len(*fileTemplateList))
		for _, template := range *fileTemplateList {
			fmt.Printf("- %v\n", template.GetFullInformation())
		}

		workspaceName = ""
		fmt.Println()
	}

	fmt.Println("*** Workspace Specific Templates ***")
	fmt.Println()

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	sharedTemplateList, err := sharedTemplateService.ListByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) shared template available\n\n", len(*sharedTemplateList))
	for _, template := range *sharedTemplateList {
		fmt.Printf("- %v\n", template.GetFullInformation())
	}

	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()

	codeTemplateList, err := codeTemplateService.ListByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) code template available\n\n", len(*codeTemplateList))
	for _, template := range *codeTemplateList {
		fmt.Printf("- %v\n", template.GetFullInformation())
	}

	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()

	fileTemplateList, err := fileTemplateService.ListByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) file template available\n\n", len(*fileTemplateList))
	for _, template := range *fileTemplateList {
		fmt.Printf("- %v\n", template.GetFullInformation())
	}
}


