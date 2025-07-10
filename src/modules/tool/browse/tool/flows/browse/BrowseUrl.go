package browse

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/tool/browse_tool_contract"
	browse_tool_payload_structs "parsdevkit.net/modules/tool/browse_tool_payload/structs"
)

type OpenBrowser struct{ flowx.BaseStep }

func (s *OpenBrowser) Name() string { return "OpenBrowser" }

func (s *OpenBrowser) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[browse_tool_contract.ToolInterface]()

	tool, ok := flowx.Get[browse_tool_payload_structs.ToolBaseStruct](fc, "tool")
	if !ok {
		panic(fmt.Errorf("xxx: tool parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to browse %v", tool.Header.Name)
	if _, err := service.Browse(tool); err != nil {
		return err
	}

	return nil
}

func (s *OpenBrowser) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	return nil
}
