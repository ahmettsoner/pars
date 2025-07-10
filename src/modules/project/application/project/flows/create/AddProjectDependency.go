package create

import (
	"context"
	"fmt"

	"parsdevkit.net/application/ioc"
	applicationProject "parsdevkit.net/application/structs/project"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type AddProjectDependency struct {
	flowx.BaseStep
	dependency applicationProject.Dependency
}

func NewAddProjectDependency(dependency applicationProject.Dependency) *AddProjectDependency {
	return &AddProjectDependency{
		dependency: dependency,
	}
}

func (s *AddProjectDependency) Name() string {
	return fmt.Sprintf("AddProjectDependency: %s", &s.dependency.Name)
}

func (s *AddProjectDependency) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		err := service.AddDependenciesToProject(project, s.dependency)
		if err != nil {
			return fmt.Errorf("xxx: Application Project oluştururken, projelerin paketi ekleme sırasında hata meydana geldi: '%s'\n%w", project.Header.Name, err)
		}
	}
	return nil
}

func (s *AddProjectDependency) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}

func (s *AddProjectDependency) IgnoreError() bool {
	return true
}
