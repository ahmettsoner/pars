package create

import (
	"context"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type SetProjectReferences struct{ flowx.BaseStep }

func (s *SetProjectReferences) Name() string { return "SetProjectReferences" }

func (s *SetProjectReferences) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		if project.Specifications.References != nil {
			err := service.AddReferenceToProject(project, project.Specifications.References...)
			if err != nil {
				return fmt.Errorf("xxx: Application Project oluştururken, projelerin paketi ekleme sırasında hata meydana geldi: '%s'\n%w", project.Header.Name, err)
			}
		}

		fmt.Printf("%v Project created\n", project.Header.Name)
	}
	return nil
}

func (s *SetProjectReferences) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
