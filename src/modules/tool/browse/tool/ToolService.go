package browse_tool

import (
	"fmt"
	"os/exec"
	"runtime"

	"parsdevkit.net/modules/tool/browse_tool_contract"
	browse_tool_payload_structs "parsdevkit.net/modules/tool/browse_tool_payload/structs"
	"parsdevkit.net/pkg/utilities/url"
)

type ToolService struct {
}

func NewToolService(environment string) browse_tool_contract.ToolInterface {

	return &ToolService{}
}

func (s ToolService) Browse(model browse_tool_payload_structs.ToolBaseStruct) (*browse_tool_payload_structs.ToolBaseStruct, error) {
	err := openBrowser(url.EnsureProtocol(model.Browse.Url))
	if err != nil {
		return nil, fmt.Errorf("Failed to open browser:", err)
	}
	return &model, nil
}
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin": // macOS
		cmd = "open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}

	return exec.Command(cmd, args...).Start()
}
