package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
)

type SaveWorkspace struct{ flowx.BaseStep }

func (s *SaveWorkspace) Name() string { return "SaveWorkspace" }

func (s *SaveWorkspace) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	workspace, ok := flowx.Get[basic_workspace_payload_structs.WorkspaceBaseStruct](fc, "workspace")
	if !ok {
		panic(fmt.Errorf("xxx: workspace parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", workspace.Header.Name)
	if _, err := service.SaveWorkspace(workspace); err != nil {
		return err
	}

	return nil
}

func (s *SaveWorkspace) Compensate(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[basic_workspace_contract.WorkspaceInterface]()

	workspace, ok := flowx.Get[basic_workspace_payload_structs.WorkspaceBaseStruct](fc, "workspace")
	if !ok {
		panic(fmt.Errorf("xxx: workspace parametresi hatalı tipte"))
	}

	logrus.Warnf("yyy: rolling back the creating %v!", workspace.Header.Name)
	if _, err := service.UndoSaveWorkspace(workspace); err != nil {
		return fmt.Errorf("xxx: Code Workspace Information kayıt sırasında meydana gelen hata için uygulanan rollback hatası oluştu: '%s'\n%w", &workspace.Header.Name, &err)
	}

	fc.Log("Rollback: Persist workspace %s", workspace.Header.Name)
	return nil
}
