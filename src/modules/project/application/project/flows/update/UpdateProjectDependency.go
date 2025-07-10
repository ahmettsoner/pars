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

type UpdateProjectDependency struct {
	flowx.BaseStep
	oldDependency applicationProject.Dependency
	newDependency applicationProject.Dependency
}

func NewUpdateProjectDependency(old applicationProject.Dependency, new applicationProject.Dependency) *UpdateProjectDependency {
	return &UpdateProjectDependency{
		oldDependency: old,
		newDependency: new,
	}
}

func (s *UpdateProjectDependency) Name() string {
	return fmt.Sprintf("UpdateProjectDependency: %s", s.oldDependency.Name)
}

func (s *UpdateProjectDependency) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		if err := service.RemoveDependencyFromProject(project, s.oldDependency); err != nil {
			return err
		}

		if err := service.AddDependenciesToProject(project, s.newDependency); err != nil {
			return err
		}
	}
	return nil
}

func (s *UpdateProjectDependency) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
