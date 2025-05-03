package submit

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	parsCMDCommon "parsdevkit.net/core/cmd"
	"parsdevkit.net/core/utils"
	"parsdevkit.net/engines/applicationProject"
	"parsdevkit.net/operation/services"

	"github.com/spf13/cobra"
)

var (
	// declarationFile bool = false
	workspaceName string
	noInit        bool = true
	filePaths     []string
)

var maxArgumentCount int = 0

var SubmitCmd = &cobra.Command{
	Use:     "submit",
	Aliases: []string{"s"},
	Short:   "Initialize project",
	Long:    `Initialize project`,
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

	workspaceName = parsCMDCommon.GetActiveWorkspaceName(workspaceName)

	if len(filePaths) > 0 {

		allFiles, err := utils.WalkDir(filePaths...)
		if err != nil {
			fmt.Println("Error processing file paths:", err)
			return
		}

		applicationProjectService := applicationProject.ApplicationProjectEngine{}
		if err := applicationProjectService.CreateProjectsFromFile(workspaceName, !noInit, allFiles...); err != nil {
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
	SubmitCmd.Flags().StringVarP(&workspaceName, "workspace", "w", "", "Workspace name")
	SubmitCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)

	SubmitCmd.Flags().BoolVarP(&noInit, "no-init", "", false, "Create project but do not initialize")

	// SubmitCmd.AddCommand(pars.ParsCmd)
	// SubmitCmd.AddCommand(dotnet.DotnetCmd)
	// SubmitCmd.AddCommand(angular.AngularCmd)
	// SubmitCmd.AddCommand(goPkg.GoCmd)
	// SubmitCmd.Flags().BoolVarP(&declarationFile, "from-file", "", false, "Create from declaration file")
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

func listWorkspaceNameSuggestions(args []string, toComplete string) []string {
	var suggestions = make([]string, 0)
	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, workspace := range *workspaceList {
		if !utils.Contains(args, workspace.Name) && strings.HasPrefix(workspace.Name, toComplete) {
			suggestions = append(suggestions, workspace.Name)
		}
	}
	return suggestions
}

func workspaceFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	completions := []string{}

	suggestions := listWorkspaceNameSuggestions(args, toComplete)
	completions = append(completions, suggestions...)

	return completions, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
}
