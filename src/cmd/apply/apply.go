package apply

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/components/schema"
	"parsdevkit.net/components/workspace"
	"parsdevkit.net/pkg/utilities/array"
	"parsdevkit.net/pkg/utilities/json"
	_string "parsdevkit.net/pkg/utilities/string"

	"github.com/spf13/cobra"
)

type ApplyOptions struct {
	Name      string
	Workspace string
	NoInit    bool
	FilePaths []string
}

var commandOptions = ApplyOptions{
	NoInit: true,
}
var maxArgumentCount int = 0

var ApplyCmd = &cobra.Command{
	Use:               "apply",
	Aliases:           []string{"a"},
	Short:             "Apply Schema(s)",
	Long:              `Apply Schema(s)`,
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
		return fmt.Errorf("Please provide a file location for the apply schema(s)")
	}

	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.Name) && len(args) > 0 {
		commandOptions.Name = args[0]
	}

	if _string.IsEmpty(commandOptions.Workspace) {
		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		var workspaceName, err = workspace.GetActiveWorkspaceNameV2(appCtx, commandOptions.Workspace)
		if err != nil {
			return fmt.Errorf("failed to find active workspace '%s'\n%w", commandOptions.Name, err)
		}
		commandOptions.Workspace = workspaceName
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {

	if len(commandOptions.FilePaths) > 0 {
		fmt.Printf("\n\n════════════════════════════════════\n")
		fmt.Printf("🔍 Applying Schemas: %s\n", _string.Concat(", ", commandOptions.FilePaths...))
		fmt.Printf("═══════════════════════════════════\n\n")
		result, err := schema.GetAllManifestFilesInPath(commandOptions.FilePaths...)

		if err != nil {
			return fmt.Errorf("%w", err)
		}

		if len(result) == 0 {
			fmt.Printf("Belirtilen dosya/dizinde uygun schema bulunamadı!\n")
			return nil
		}

		var loadedSchemas []string = make([]string, 0)
		for _, data := range result {

			if err := data.Validate(); err != nil {
				jsonObject, _ := json.ToJson(data)
				return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
			}

			loadedSchemas = append(loadedSchemas, data.GetHeader().Name)
		}
		fmt.Printf("✅ Loaded Schemas: %s\n", _string.Concat(", ", loadedSchemas...))

		appCtx := application.GetContext()
		if appCtx == nil {
			return fmt.Errorf("xxx: Current workspace bulunamadı")
		}
		err = engines.DispatchEngineProcess(appCtx, result)
		if err != nil {
			return fmt.Errorf("Engine processing failed: %v", err)
		}
	} else {
		cmd.Help()
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = ApplyOptions{
		NoInit: true,
	}
}

func init() {
	addSubCommands()
}

func addSubCommands() {
	ApplyCmd.Flags().StringVarP(&commandOptions.Workspace, "workspace", "w", "", "Workspace name")
	ApplyCmd.RegisterFlagCompletionFunc("workspace", workspaceFlagCompletion)

	ApplyCmd.Flags().BoolVarP(&commandOptions.NoInit, "no-init", "", false, "Create project but do not initialize")

	ApplyCmd.Flags().StringSliceVarP(&commandOptions.FilePaths, "file", "f", nil, "Comma-separated list of declaration files")
	ApplyCmd.RegisterFlagCompletionFunc("file", fileFlagCompletion)
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
	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	workspaceList, err := workspaceService.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, workspace := range *workspaceList {
		if !array.ContainsSlice(args, workspace.Header.Name) && strings.HasPrefix(workspace.Header.Name, toComplete) {
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
