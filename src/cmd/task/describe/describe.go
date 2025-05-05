package describe

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	parsCMDCommon "parsdevkit.net/core/cmd"

	"github.com/spf13/cobra"
)

var (
	name          string
	workspaceName string
	force         string
)

var DescribeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "Information about project",
	Long:    `Information about project`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {
	if utils.IsEmpty(name) {
		if len(args) == 0 {
			fmt.Println("Please provide a name for the new project")
			os.Exit(1)
		} else if len(args) > 0 {
			name = args[0]
		}
	}

	if utils.IsEmpty(name) {
		cmd.Help()
		os.Exit(0)
	}

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	groupService := services.NewGroupService(utils.GetEnvironment())
	group, err := groupService.GetByName(name)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("%v", group.Name)
	fmt.Println(name)

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	projectList, err := projectService.ListByGroupName(group.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tProjects:\n")
	for _, e := range *projectList {
		name := fmt.Sprintf("\t\t - %v", e.GetFullInformation())
		fmt.Println(name)
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	DescribeCmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	DescribeCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
}
