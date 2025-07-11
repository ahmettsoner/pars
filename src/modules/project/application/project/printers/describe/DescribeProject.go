package list

import (
	"os"

	"parsdevkit.net/internal/printerx"
	"parsdevkit.net/models"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
)

type ViewModel struct {
	Name         string
	Set          string
	Group        string
	Platform     string
	ProjectType  models.ProjectType
	Runtime      application_project_payload_structs.Runtime
	Language     application_project_payload_structs.Language
	Path         []string
	Package      []string
	Layers       []string
	Tags         []string
	Labels       []string
	Dependencies []string
	References   []ReferenceViewModel
}

type ReferenceViewModel struct {
	Name   string
	Group  string
	Set    string
	Tags   []string
	Labels []string
}

type DescribeProject struct {
	Project ViewModel
}

func (s *DescribeProject) Print() error {

	printer := &printerx.DetailPrinter{ShowEmpty: true}
	printer.Print(os.Stdout, s.Project)

	return nil
}
