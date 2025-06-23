package submit

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/core/utils"
	"parsdevkit.net/engines/group"

	"github.com/spf13/cobra"
)

type SubmitOptions struct {
	Name      string
	NoInit    bool
	FilePaths []string
}

var commandOptions = SubmitOptions{
	NoInit: true,
}
var maxArgumentCount int = 0

var SubmitCmd = &cobra.Command{
	Use:     "submit",
	Aliases: []string{"r"},
	Short:   "Group Information",
	Long:    `Group Information`,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	if len(commandOptions.FilePaths) > 0 {
		groupService := group.GroupEngine{}
		if err := groupService.CreateGroupsFromFile(!commandOptions.NoInit, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		if utils.IsEmpty(commandOptions.Name) {
			if len(args) == 0 {
				fmt.Println("Please provide a name for the submit group")
				os.Exit(1)
			} else if len(args) > 0 {
				commandOptions.Name = args[0]
			}
		}

		if utils.IsEmpty(commandOptions.Name) {
			cmd.Help()
			os.Exit(0)
		}

		var structData = struct {
			Name string
		}{
			Name: commandOptions.Name,
		}

		var templateFilePath = "/group/group.yaml.templ"

		groupService := group.GroupEngine{}
		if err := groupService.CreateGroupsFromTemplate(!commandOptions.NoInit, structData, templateFilePath); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func init() {
	addSubCommands()
}

func addSubCommands() {

	SubmitCmd.Flags().StringVarP(&commandOptions.Name, "name", "n", "", "Group name")
	SubmitCmd.Flags().BoolVarP(&commandOptions.NoInit, "no-init", "", false, "Create group but do not initialize")

	SubmitCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "files", "f", nil, "Comma-separated list of declaration files")
}
