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

type UpdateProjectDependencies struct{ flowx.BaseStep }

func (s *UpdateProjectDependencies) Name() string { return "UpdateProjectDependencies" }

func (s *UpdateProjectDependencies) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		existingProject, err := service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}

		if !reflect.DeepEqual(project.Specifications.Dependencies, existingProject.Specifications.Dependencies) {
			result := diffx.DiffSlice(existingProject.Specifications.Dependencies, project.Specifications.Dependencies)

			if len(result.Created) > 0 {
				err := service.AddDependenciesToProject(project, result.Created...)
				if err != nil {
					return err
				}

			}
			if len(result.Updated) > 0 {
				for _, item := range result.Updated {
					err := service.RemoveDependencyFromProject(project, item.Old)
					if err != nil {
						return err
					}
					err = service.AddDependenciesToProject(project, item.New)
					if err != nil {
						return err
					}
				}
			}
			if len(result.Deleted) > 0 {
				err := service.RemoveDependencyFromProject(project, result.Deleted...)
				if err != nil {
					return err
				}
			}
		}

		if project.Specifications.References != nil {
			err := service.AddReferenceToProject(project, project.Specifications.References...)
			if err != nil {
				return fmt.Errorf("xxx: Application Project oluştururken, projelerin paketi ekleme sırasında hata meydana geldi: '%s'\n%w", project.Header.Name, err)
			}
		}
	}
	return nil
}

func (s *UpdateProjectDependencies) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}

func (s *UpdateProjectDependencies) IgnoreError() bool {
	return true
}
