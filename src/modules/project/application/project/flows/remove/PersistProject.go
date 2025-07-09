package create

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/internal/flowx"
	"parsdevkit.net/modules/project/application_project_contract"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type SaveProject struct{}

func (s *SaveProject) Name() string { return "SaveProject" }

func (s *SaveProject) Run(ctx context.Context, fc *flowx.FlowContext) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	init, ok := flowx.Get[bool](fc, "init")
	if !ok {
		panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
	}

	project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
	if !ok {
		panic(fmt.Errorf("xxx: project parametresi hatalı tipte"))
	}

	logrus.Debugf("trying to create %v", project.Header.Name)
	if _, err := service.Create(project, init); err != nil {
		return err
	}

	fmt.Printf("%v Project created\n", project.Header.Name)
	return nil
}

func (s *SaveProject) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	project, ok := flowx.Get[application_project_payload_structs.ProjectBaseStruct](fc, "project")
	if !ok {
		panic(fmt.Errorf("xxx: init parametresi hatalı tipte"))
	}
	fc.Log("Rollback: Persist project %s", project.Header.Name)
	return nil
}
