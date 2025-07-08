package create

import (
	"context"
	"fmt"

	"parsdevkit.net/internal/flowx"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type GenerateProject struct{}

func (s *GenerateProject) Name() string { return "GenerateProject" }

func (s *GenerateProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

	return fmt.Errorf("dddd")
}

func (s *GenerateProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
	if !ok {
		panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
	}
	fc.Log("Rollback: Generating project %s", project.Header.Name)
	return nil
}
