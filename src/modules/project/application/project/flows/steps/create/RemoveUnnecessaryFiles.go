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

type RemoveUnnecessaryFiles struct{ flowx.BaseStep }

func (s *RemoveUnnecessaryFiles) Name() string { return "RemoveUnnecessaryFiles" }

func (s *RemoveUnnecessaryFiles) Run(ctx context.Context, fc *flowx.FlowContext) error {

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

		logrus.Debugf("trying to create %v", project.Header.Name)
		if _, err := service.RemoveUnnecessaryFiles(project); err != nil {
			return err
		}

		fmt.Printf("%v Project created\n", project.Header.Name)
	}
	return nil
}

func (s *RemoveUnnecessaryFiles) Compensate(ctx context.Context, fc *flowx.FlowContext) error {
	return nil
}

func (s *RemoveUnnecessaryFiles) IgnoreError() bool {
	return true
}
