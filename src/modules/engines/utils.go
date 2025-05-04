package engines

import (
	"fmt"

	applicationproject "parsdevkit.net/structs/project/application-project"
)

func PrintRefInfo(projects []applicationproject.ProjectSpecification) {
	for _, project := range projects {
		fmt.Printf("%v - %v (%v)\n", project.Group, project.Name, project.Workspace)
		for _, ref := range project.Configuration.References {
			fmt.Printf("\t%v - %v (%v)\n", ref.Specifications.Group, ref.Name, ref.Specifications.Workspace)
		}
	}
}
