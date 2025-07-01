package destroy

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"parsdevkit.net/application"
	parsCMDCommon "parsdevkit.net/core/cmd"
	"parsdevkit.net/core/utilities"
	"parsdevkit.net/core/utilities/json"
	"parsdevkit.net/core/utils"
	"parsdevkit.net/operation/services"

	"github.com/spf13/cobra"
)

type DestroyOptions struct {
	Name      string
	Workspace string
	NoInit    bool
	FilePaths []string
}

var commandOptions = DestroyOptions{
	NoInit: true,
}
var maxArgumentCount int = 0

var DestroyCmd = &cobra.Command{
	Use:               "destroy",
	Aliases:           []string{"d"},
	Short:             "Destroy Schema(s)",
	Long:              `Destroy Schema(s)`,
	Args:              validateArgs,
	PreRunE:           prepareFunc,
	RunE:              executeFunc,
	PostRun:           afterFunc,
	ValidArgsFunction: validArguments,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	if len(commandOptions.FilePaths) == 0 {
		return fmt.Errorf("Please provide a file location for the destroy schema(s)")
	}

	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if utilities.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	if utilities.IsEmpty(commandOptions.Workspace) {
		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		var workspaceName, err = parsCMDCommon.GetActiveWorkspaceNameV2(appCtx, commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("failed to find active workspace '%s'\n%w", commandOptions.Name, err)
		}
		commandOptions.Workspace = workspaceName
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	if len(commandOptions.FilePaths) > 0 {

		result, err := application.GetAllManifestFilesInPath(commandOptions.FilePaths...)

		if err != nil {
			log.Printf("❌ Error: %v", err)
		} else {
			for _, data := range result {

				if err := data.Validate(); err != nil {
					jsonObject, _ := json.ToJson(data)
					return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
				}

				fmt.Printf("✅ Loaded: %#v\n", data.GetHeader().Name)
			}

			appCtx := application.GetContext()
			if appCtx == nil {
				return fmt.Errorf("xxx: Current workspace bulunamadı")
			}
			err = application.DispatchEngineDestroy(appCtx, result)
			if err != nil {
				log.Fatalf("Engine processing failed: %v", err)
			}
		}

		fmt.Fprintf(os.Stdout, "✔ Schema(s) '%v' destroyed successfully\n", commandOptions.FilePaths)
	} else {
		cmd.Help()
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = DestroyOptions{
		NoInit: true,
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	DestroyCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
	DestroyCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)

	DestroyCmd.Flags().BoolVarP(&commandOptions.NoInit, "no-init", "", false, "Create project but do not initialize")

	DestroyCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "file", "f", nil, "Comma-separated list of declaration files")
	DestroyCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
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
		if !utilities.Contains(args, workspace.Header.Name) && strings.HasPrefix(workspace.Header.Name, toComplete) {
			suggestions = append(suggestions, workspace.Header.Name)
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
