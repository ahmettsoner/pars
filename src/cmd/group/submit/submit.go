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

var (
	name      string
	noInit    bool = true
	filePaths []string
)

var maxArgumentCount int = 0

var SubmitCmd = &cobra.Command{
	Use:     "submit",
	Aliases: []string{"s"},
	Short:   "Group Information",
	Long:    `Group Information`,
	Run:     executeFunc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > maxArgumentCount {
			return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
		}
		return nil
	},
	ValidArgsFunction: validArguments,
}

func executeFunc(cmd *cobra.Command, args []string) {
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
	if len(filePaths) > 0 {

		allFiles, err := utils.WalkDir(filePaths...)
		if err != nil {
			fmt.Println("Error processing file paths:", err)
			return
		}

		groupService := group.GroupEngine{}
		if err := groupService.CreateGroupsFromFile(!noInit, allFiles...); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Please provide a file location for the submit group(s)")
		os.Exit(1)
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {

	SubmitCmd.Flags().BoolVarP(&noInit, "no-init", "", false, "Create group but do not initialize")

	SubmitCmd.Flags().StringSliceVarP(&filePaths, "file", "f", nil, "Comma-separated list of declaration files")
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
