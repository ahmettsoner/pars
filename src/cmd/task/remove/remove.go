package remove

import (
	"fmt"
	"log"
	"os"

	parsCMDCommon "parsdevkit.net/core/cmd"
	"parsdevkit.net/engines/commonTask"

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
	Short:   "Task Information",
	Long:    `Task Information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if len(filePaths) > 0 {
		taskService := commonTask.CommonTaskEngine{}
		if err := taskService.RemoveTasksFromFile(true, filePaths...); err != nil {
			log.Fatal(err)
		}
	} else {

		if len(names) == 0 {
			if len(args) == 0 {
				fmt.Println("Please provide a name for the remove the remove")
				os.Exit(1)
			} else if len(args) > 0 {
				name = args[0]
			}
		}

		if len(names) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		checkGlobals := utils.IsEmpty(workspaceName)
		taskService := services.NewCommonTaskService(utils.GetEnvironment())

		for _, name := range names {
			if checkGlobals {
				workspaceName = "None"

				if taskService.IsExists(name, workspaceName) {
					task, err := taskService.Remove(name, workspaceName, true)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Task (" + task.Name + ") deleted permanently")
				}

				workspaceName = ""
			}

			workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

			if taskService.IsExists(name, workspaceName) {
				task, err := taskService.Remove(name, workspaceName, true)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Task (" + task.Name + ") deleted permanently")
			}
		}
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	RemoveCmd.Flags().StringSliceVarP(&names, "name", "n", nil, "Template names")

	RemoveCmd.Flags().StringSliceVarP(&filePaths, "files", "f", nil, "Comma-separated list of declaration files")
}
