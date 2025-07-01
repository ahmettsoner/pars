package describe

import (
	"fmt"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/spf13/cobra"
)

type DescribeOptions struct {
	Name      string
	Workspace string
	Force     string
}

var commandOptions DescribeOptions
var maxArgumentCount int = 1

var DescribeCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "Information about project",
	Long:    `Information about project`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if utilities.IsEmpty(commandOptions.Name) && len(args) == 0 {
		return fmt.Errorf("error: group name is required. Provide it with '--name' or as an argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("error: too many arguments. Only group name is expected.")
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utilities.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}
	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	groupService := services.NewGroupService(utils.GetEnvironment())
	group, err := groupService.GetByName(commandOptions.Name)
	if err != nil {
		return fmt.Errorf("Failed to describe task '%s'\n%w", commandOptions.Name, err)
	}

	name := fmt.Sprintf("%v", group.Name)
	fmt.Println(name)

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	projectList, err := projectService.ListByGroupName(group.Name)
	if err != nil {
		return fmt.Errorf("Failed to retrieve group tasks '%s'\n%w", commandOptions.Name, err)
	}

	fmt.Printf("\tProjects:\n")
	for _, e := range *projectList {
		name := fmt.Sprintf("\t\t - %v", e.GetFullInformation())
		fmt.Println(name)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = DescribeOptions{}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	DescribeCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Project name")

	DescribeCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
}
