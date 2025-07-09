package update

import (
	"context"
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/diffx"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type GenerateProject struct{ flowx.BaseStep }

func (s *GenerateProject) Name() string { return "GenerateProject" }

func (s *GenerateProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	init, ok := flowx.Get[bool](fc, "init")
	if !ok {
		panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
	}

	if init {
		project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
		if !ok {
			panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
		}

		existingProject, err := service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}

		if !reflect.DeepEqual(project.Specifications.Layers, existingProject.Specifications.Layers) {

			result := diffx.DiffSlice(existingProject.Specifications.Layers, project.Specifications.Layers)

			if len(result.Created) > 0 {
				err := service.CreateLayerFolder(project, result.Created...)
				if err != nil {
					return err
				}
			}
			if len(result.Updated) > 0 {
				//TODO Burda değişiklik tespit edilerek eğer move ve rename yapılabilir dosya ve klasörlere, silmekten daha güvenli
				for _, item := range result.Updated {
					err := service.DeleteLayerFolder(project, item.Old)
					if err != nil {
						return err
					}
					err = service.CreateLayerFolder(project, item.New)
					if err != nil {
						return err
					}
				}
			}
			if len(result.Deleted) > 0 {
				err := service.DeleteLayerFolder(project, result.Deleted...)
				if err != nil {
					return err
				}
			}
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

		logrus.Debugf("trying to create %v", project.Header.Name)
		if _, err := service.GenerateProject(project); err != nil {
			return err
		}

		fmt.Printf("%v Project created\n", project.Header.Name)
	}
	return nil
}

func (s *GenerateProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	service := ioc.Get[application_project_contract.ProjectInterface]()

	init, ok := flowx.Get[bool](fc, "init")
	if !ok {
		panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
	}

	if init {
		project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
		if !ok {
			panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
		}

		logrus.Debugf("trying to create %v", project.Header.Name)
		if _, err := service.UndoGenerateProject(project); err != nil {
			return err
		}

		fmt.Printf("%v Project created\n", project.Header.Name)
		fc.Log("Rollback: Creating project %s", project.Header.Name)
	}
	return nil
}
