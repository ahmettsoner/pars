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
	Short:   "List task(s)",
	Long:    `List task(s)`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	checkGlobals := utils.IsEmpty(workspaceName)
	taskService := services.NewCommonTaskService(utils.GetEnvironment())

	if checkGlobals {
		fmt.Println()
		fmt.Println("*** Global Tasks ***")
		fmt.Println()

		workspaceName = "None"

		taskList, err := taskService.ListByWorkspace(workspaceName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("(%d) task available\n\n", len(*taskList))
		for _, task := range *taskList {
			fmt.Printf("- %v\n", task.Name)
		}

		workspaceName = ""
	}

	fmt.Println()
	fmt.Println("*** Workspace Specific Tasks ***")
	fmt.Println()

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	taskList, err := taskService.ListByWorkspace(workspaceName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("(%d) task available\n\n", len(*taskList))
	for _, task := range *taskList {
		fmt.Printf("- %v\n", task.Name)
	}

}


