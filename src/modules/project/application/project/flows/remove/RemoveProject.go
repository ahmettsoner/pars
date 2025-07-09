package remove

import (
	"context"
	"fmt"

	"parsdevkit.net/internal/flowx"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type RemoveProject struct{}

func (s *RemoveProject) Name() string { return "RemoveProject" }

func (s *RemoveProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

	return fmt.Errorf("dddd")
}

func (s *RemoveProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
	if !ok {
		panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
	}
	fc.Log("Rollback: Updating project %s", project.Header.Name)
	return nil
}
