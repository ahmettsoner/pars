package submit

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"parsdevkit.net/core/utils"
	"parsdevkit.net/engines/group"

	"github.com/spf13/cobra"
)

type SubmitOptions struct {
	Name      string
	Workspace string
	NoInit    bool
	FilePaths []string
}

var commandOptions = SubmitOptions{
	NoInit: true,
}
var maxArgumentCount int = 0

var SubmitCmd = &cobra.Command{
	Use:               "submit",
	Aliases:           []string{"s"},
	Short:             "Group Information",
	Long:              `Group Information`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utils.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	// if len(args) == 1 {
	// 	name = args[0]

	// 	var structData = struct {
	// 		Name string
	// 	}{
	// 		Name: name,
	// 	}

	// 	var templateFilePath = "/group/group.yaml.templ"

	// 	groupService := group.GroupEngine{}
	// 	if err := groupService.CreateGroupsFromTemplate(!noInit, structData, templateFilePath); err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else
	if len(commandOptions.FilePaths) > 0 {

		allFiles, err := utils.WalkDir(commandOptions.FilePaths...)
		if err != nil {
			return fmt.Errorf("Error processing file paths: %v", allFiles)
		}

		groupService := group.GroupEngine{}
		if err := groupService.CreateGroupsFromFile(!commandOptions.NoInit, allFiles...); err != nil {
			log.Fatal(err)
		}
	} else {
		return fmt.Errorf("Please provide a file location for the submit group(s)")
	}

	return nil
}

func init() {
	addSubCommands()
}

func addSubCommands() {

	SubmitCmd.Flags().BoolVarP(&commandOptions.NoInit, "no-init", "", false, "Create group but do not initialize")

	SubmitCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "file", "f", nil, "Comma-separated list of declaration files")
	SubmitCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// if len(args) == 1 {
	// 	completions := []string{}

	// 	return completions, cobra.ShellCompDirectiveNoSpace
	// }

	return nil, cobra.ShellCompDirectiveNoFileComp
}
func fileFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	items, _ := filepath.Glob(filepath.Join(toComplete, "*"))
	completions := []string{}
	for _, item := range items {
		if _, err := os.Stat(item); err == nil {
			item, _ := GetLastComponent(item)
			completions = append(completions, item)
		}
	}
	return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
}
func GetLastComponent(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		// Get the last directory name
		return filepath.Base(filepath.Clean(path)), nil
	} else {
		// Get the filename
		return filepath.Base(path), nil
	}
}
