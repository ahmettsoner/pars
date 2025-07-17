package browse

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/application/schemas"
	"parsdevkit.net/application/structs/tool"
	applicationTool "parsdevkit.net/application/structs/tool"
	"parsdevkit.net/pkg/utilities/json"

	browse_tool_payload_structs "parsdevkit.net/modules/tool/browse_tool_payload/structs"
)

type BrowseOptions struct {
	URL string
}

var commandOptions BrowseOptions
var maxArgumentCount int = 1

var BrowseCmd = &cobra.Command{
	Use:     "browse",
	Aliases: []string{""},
	Short:   "Browse Url",
	Long:    `Browse Url`,
	Args:    validateArgs,
	PreRunE: prepareFunc,
	RunE:    executeFunc,
	PostRun: afterFunc,
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.URL) && len(args) == 0 {
		return fmt.Errorf("Please provide an access URL using the --url flag or as argument.")
	}
	if len(args) > maxArgumentCount {
		return fmt.Errorf("Undefined argument(s) found: %v", args[maxArgumentCount:])
	}
	return nil
}

func prepareFunc(cmd *cobra.Command, args []string) error {
	if _string.IsEmpty(commandOptions.URL) && len(args) > 0 {
		commandOptions.URL = args[0]
	}

	return nil
}

func executeFunc(cmd *cobra.Command, args []string) error {
	var result []schemas.SchemaInterface = make([]schemas.SchemaInterface, 0)
	result = append(result, &browse_tool_payload_structs.ToolBaseStruct{
		Header: schemas.NewSchemaHeader(
			schemas.StructTypes.Tool,
			browse_tool_payload_structs.TOOL_KIND,
			commandOptions.URL,
			schemas.Metadata{},
		),
		Specifications: applicationTool.ToolSpecification{
			ToolIdentifier: tool.ToolIdentifier{},
		},
		Browse: browse_tool_payload_structs.ToolBrowse{
			Url: commandOptions.URL,
		},
	})

	var loadedSchemas []string = make([]string, 0)
	for _, data := range result {

		if err := data.Validate(); err != nil {
			jsonObject, _ := json.ToJson(data)
			return fmt.Errorf("invalid data: '%s'\n%w", jsonObject, err)
		}

		loadedSchemas = append(loadedSchemas, fmt.Sprintf("%s.%s", data.GetHeader().Name, data.GetKey()))
	}
	fmt.Printf("Loaded Schemas: %s\n", _string.Concat(", ", loadedSchemas...))
	fmt.Printf("────────────────────────────────────\n")

	appCtx := application.GetContext()
	if appCtx == nil {
		return fmt.Errorf("xxx: Current workspace bulunamadı")
	}
	err := engines.DispatchEngineBrowse(appCtx, result)
	if err != nil {
		return fmt.Errorf("Engine processing failed: %v", err)
	}

	return nil
}
func afterFunc(cmd *cobra.Command, args []string) {
	commandOptions = BrowseOptions{}
}

func init() {
	BrowseCmd.Flags().StringVarP(&commandOptions.URL, "url", "u", "", "Url")
}

func openBrowser(url string) error {
	if url == "" {
		return fmt.Errorf("URL is empty")
	}

	var openCmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		openCmd = exec.Command("open", url)
	case "windows":
		openCmd = exec.Command("cmd", "/c", "start", "", url)
	default:
		openCmd = exec.Command("xdg-open", url)
	}

	return openCmd.Start()
}
