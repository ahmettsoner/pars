package create

import (
	"context"
	"fmt"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type AddProjectReference struct {
	flowx.BaseStep
	reference application_project_payload_structs.ProjectBaseStruct
}

func NewAddProjectReference(reference application_project_payload_structs.ProjectBaseStruct) *AddProjectReference {
	return &AddProjectReference{
		reference: reference,
	}
}

func (s *AddProjectReference) Name() string {
	return fmt.Sprintf("AddProjectReference: %s", &s.reference.Header.Name)
}

func (s *AddProjectReference) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		err := service.AddReferenceToProject(project, s.reference)
		if err != nil {
			return fmt.Errorf("xxx: Application Project oluştururken, projelerin paketi ekleme sırasında hata meydana geldi: '%s'\n%w", project.Header.Name, err)
		}

	}
	return nil
}

func (s *AddProjectReference) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
