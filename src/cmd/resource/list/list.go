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

var ListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List resource(s)",
	Long:    `List resource(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {

	checkGlobals := utils.IsEmpty(workspaceName)
	objectResourceService := services.NewObjectResourceService(utils.GetEnvironment())
	dataResourceService := services.NewDataResourceService(utils.GetEnvironment())

	if checkGlobals {
		fmt.Println("*** Global Resources ***")
		fmt.Println()

		workspaceName = "None"

		objectResourceList, err := objectResourceService.ListByWorkspace(workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("(%d) object resource available\n\n", len(*objectResourceList))
		for _, resource := range *objectResourceList {
			fmt.Printf("- %v\n", resource.GetFullInformation())
		}

		fmt.Println()
		fmt.Println("--------------------------")
		fmt.Println()

		dataResourceList, err := dataResourceService.ListByWorkspace(workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("(%d) data resource available\n\n", len(*dataResourceList))
		for _, resource := range *dataResourceList {
			fmt.Printf("- %v\n", resource.GetFullInformation())
		}

		workspaceName = ""
		fmt.Println()
	}

	fmt.Println("*** Workspace Specific Resources ***")
	fmt.Println()

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	objectResourceList, err := objectResourceService.ListByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) object resource available\n\n", len(*objectResourceList))
	for _, resource := range *objectResourceList {
		fmt.Printf("- %v\n", resource.GetFullInformation())
	}

	fmt.Println()
	fmt.Println("--------------------------")
	fmt.Println()

	dataResourceList, err := dataResourceService.ListByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) data resource available\n\n", len(*dataResourceList))
	for _, resource := range *dataResourceList {
		fmt.Printf("- %v\n", resource.GetFullInformation())
	}

}


