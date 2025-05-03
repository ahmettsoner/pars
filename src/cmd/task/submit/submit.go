package submit

import (
	"fmt"
	"log"
	"os"

	"parsdevkit.net/core/utils"
	"parsdevkit.net/engines/group"

	"github.com/spf13/cobra"
)

var (
	// declarationFile bool = false
	name      string
	noInit    bool = true
	filePaths []string
)
var SubmitCmd = &cobra.Command{
	Use:     "submit",
	Aliases: []string{"r"},
	Short:   "Group Information",
	Long:    `Group Information`,
	Run:     executeFunc,
}

func executeFunc(cmd *cobra.Command, args []string) {

	// workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	if len(filePaths) > 0 {
		groupService := group.GroupEngine{}
		if err := groupService.CreateGroupsFromFile(!noInit, filePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		if utils.IsEmpty(name) {
			if len(args) == 0 {
				fmt.Println("Please provide a name for the submit group")
				os.Exit(1)
			} else if len(args) > 0 {
				name = args[0]
			}
		}

		if utils.IsEmpty(name) {
			cmd.Help()
			os.Exit(0)
		}

		var structData = struct {
			Name string
		}{
			Name: name,
		}

		var templateFilePath = "/group/group.yaml.templ"

		groupService := group.GroupEngine{}
		if err := groupService.CreateGroupsFromTemplate(!noInit, structData, templateFilePath); err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {

	SubmitCmd.Flags().StringVarP(&name, "name", "n", "", "Group name")
	SubmitCmd.Flags().BoolVarP(&noInit, "no-init", "", false, "Create group but do not initialize")

	// SubmitCmd.Flags().BoolVarP(&declarationFile, "from-file", "", false, "Create from declaration file")
	SubmitCmd.Flags().StringSliceVarP(&filePaths, "files", "f", nil, "Comma-separated list of declaration files")
}
