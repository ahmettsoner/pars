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

type UpdateProjectLayers struct{ flowx.BaseStep }

func (s *UpdateProjectLayers) Name() string { return "UpdateProjectLayers" }

func (s *UpdateProjectLayers) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

	}
	return nil
}

func (s *UpdateProjectLayers) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}
