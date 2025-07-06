package engines

import (
	"fmt"

	applicationproject "parsdevkit.net/modules/project/application_project_payload"
)

func PrintRefInfo(projects []applicationproject.ProjectSpecification) {
	for _, project := range projects {
		fmt.Printf("%v - %v (%v)\n", project.Group, project.Name, project.Workspace)
		for _, ref := range project.References {
			fmt.Printf("\t%v - %v (%v)\n", ref.Specifications.Group, ref.Header.Name, ref.Specifications.Workspace)
		}
	}
}
