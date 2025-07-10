package update

import (
	"context"
	"fmt"

	"parsdevkit.net/application/ioc"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type DeleteProjectDependency struct {
	flowx.BaseStep
	dependency applicationProject.Dependency
}

func NewDeleteProjectDependency(dependency applicationProject.Dependency) *DeleteProjectDependency {
	return &DeleteProjectDependency{
		dependency: dependency,
	}
}

func (s *DeleteProjectDependency) Name() string {
	return fmt.Sprintf("DeleteProjectDependency: %s", s.dependency.Name)
}

func (s *DeleteProjectDependency) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	init, ok := flowx.Get[bool](fc, "init")
	if !ok {
		panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
	}

	if init {
		project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
		if !ok {
			panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
		}

		if err := service.RemoveDependencyFromProject(project, s.dependency); err != nil {
			return err
		}
	}
	return nil
}

func (s *DeleteProjectDependency) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
