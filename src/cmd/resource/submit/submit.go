package submit

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"parsdevkit.net/engines/dataResource"
	"parsdevkit.net/engines/objectResource"
)

var (
	noInit    bool = false
	filePaths []string
)

var maxArgumentCount int = 0

var SubmitCmd = &cobra.Command{
	Use:     "submit",
	Aliases: []string{"s"},
	Short:   "Resource Information",
	Long:    `Resource Information`,
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
	if len(filePaths) > 0 {
		objectResourceService := objectResource.ObjectResourceEngine{}
		if err := objectResourceService.CreateResourcesFromFile(!noInit, filePaths...); err != nil {
			log.Fatal(err)
		}
		dataResourceService := dataResource.DataResourceEngine{}
		if err := dataResourceService.CreateResourcesFromFile(!noInit, filePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		cmd.Help()
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
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

	dirs, _ := filepath.Glob(filepath.Join(toComplete, "*"))
	completions := []string{}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err == nil {
			dir, _ := GetLastComponent(dir)
			completions = append(completions, dir)
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
