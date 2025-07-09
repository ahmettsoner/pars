package update

import (
	"context"
	"fmt"
	"reflect"

	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/diffx"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type UpdateProjectReferences struct{ flowx.BaseStep }

func (s *UpdateProjectReferences) Name() string { return "UpdateProjectReferences" }

func (s *UpdateProjectReferences) Run(ctx context.Context, fc *flowx.FlowContext) error {

	init, ok := flowx.Get[bool](fc, "init")
	if !ok {
		panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
	}

	if init {
		service := ioc.Get[application_project_contract.ProjectInterface]()
		project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
		if !ok {
			panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
		}

		existingProject, err := service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}

		if !reflect.DeepEqual(project.Specifications.References, existingProject.Specifications.References) {
			result := diffx.DiffSlice(existingProject.Specifications.References, project.Specifications.References)

			if len(result.Created) > 0 {
				err := service.AddReferenceToProject(project, result.Created...)
				if err != nil {
					return err
				}
			}
			if len(result.Updated) > 0 {
				for _, item := range result.Updated {
					err := service.RemoveReferenceFromProject(project, item.Old)
					if err != nil {
						return err
					}
					err = service.AddReferenceToProject(project, item.New)
					if err != nil {
						return err
					}
				}
			}
			if len(result.Deleted) > 0 {
				err := service.RemoveReferenceFromProject(project, result.Deleted...)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}

func (s *UpdateProjectReferences) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
