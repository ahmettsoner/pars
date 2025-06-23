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

type SubmitOptions struct {
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
	Short:             "Resource Information",
	Long:              `Resource Information`,
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

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	if len(commandOptions.FilePaths) > 0 {
		objectResourceService := objectResource.ObjectResourceEngine{}
		if err := objectResourceService.CreateResourcesFromFile(!commandOptions.NoInit, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}
		dataResourceService := dataResource.DataResourceEngine{}
		if err := dataResourceService.CreateResourcesFromFile(!commandOptions.NoInit, commandOptions.FilePaths...); err != nil {
			log.Fatal(err)
		}
	} else {
		cmd.Help()
	}

	return nil
}

func validArguments(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	return nil, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	SubmitCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "file", "f", nil, "Comma-separated list of declaration files")
	SubmitCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
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
