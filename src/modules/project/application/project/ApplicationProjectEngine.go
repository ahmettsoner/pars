package application_project

import (
	"errors"
	"reflect"

	"fmt"

	"parsdevkit.net/application/ioc"
	applicationProject "parsdevkit.net/application/structs/project"

	"parsdevkit.net/modules/project/application_project_contract"

	"parsdevkit.net/modules/group/basic_group_contract"
	group_payload "parsdevkit.net/modules/group/basic_group_payload"
	applicationprojectStruct "parsdevkit.net/modules/project/application_project_payload"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	workspaceStruct "parsdevkit.net/modules/workspace/basic_workspace_payload"
	_string "parsdevkit.net/pkg/utilities/string"

	"parsdevkit.net/application"
	"parsdevkit.net/application/schemas"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/application/engines"
	"parsdevkit.net/pkg/utilities/encrypt"
)

type ApplicationProjectEngine struct{}

func (s ApplicationProjectEngine) Validate(data []schemas.SchemaInterface) bool {
	for _, item := range data {
		_, ok := item.(*applicationprojectStruct.ProjectBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s ApplicationProjectEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	applicationprojects := make([]applicationprojectStruct.ProjectBaseStruct, 0, len(data))

	for _, item := range data {
		applicationproject, ok := item.(*applicationprojectStruct.ProjectBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected applicationprojectStruct.ProjectBaseStruct, got %T", item)
		}

		if err := s.completeInformation(ctx, applicationproject); err != nil {
			return err
		}
		applicationprojects = append(applicationprojects, *applicationproject)
	}

	return s.createProjects(applicationprojects, true)
}
func (s ApplicationProjectEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	applicationprojects := make([]applicationprojectStruct.ProjectBaseStruct, 0, len(data))

	for _, item := range data {
		applicationproject, ok := item.(*applicationprojectStruct.ProjectBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected applicationprojectStruct.ProjectBaseStruct, got %T", item)
		}

		if err := s.completeInformation(ctx, applicationproject); err != nil {
			return err
		}
		applicationprojects = append(applicationprojects, *applicationproject)
	}

	return s.removeProjects(applicationprojects, true)
}
func (s ApplicationProjectEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  "Project.Application",
		Order: 2000,
	}
}

func (s ApplicationProjectEngine) createProjects(projects []applicationprojectStruct.ProjectBaseStruct, init bool) error {

	projectsReadyToCreate := make([]applicationprojectStruct.ProjectBaseStruct, 0)
	projectsForUpdate := make([]applicationprojectStruct.ProjectBaseStruct, 0)
	projectService := ioc.Get[application_project_contract.ProjectInterface]()

	for _, project := range projects {
		ok, err := projectService.IsExists(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(project)
			if err != nil {
				return err
			}
			structHash, err := projectService.GetHash(project.GetFullName(), project.Specifications.Workspace)
			if err != nil {
				return err
			}

			if newModelHash != structHash {
				projectsForUpdate = append(projectsForUpdate, project)
			}
		} else {
			projectsReadyToCreate = append(projectsReadyToCreate, project)
			// for _, reference := range applicationprojectStruct.Specifications.References {
			// 	if _, ok := projectReferenceMap[applicationprojectStruct.Specifications.GetUniqueKey()][reference.GetUniqueKey()]; ok {
			// 		continue
			// 	} else {
			// 		if ok := projectService.IsExists(reference.GetFullName(), reference.Workspace.Name); ok {
			// 			projectReferenceMap[applicationprojectStruct.Specifications.GetUniqueKey()][reference.GetUniqueKey()] = reference
			// 		}
			// 	}
			// }
		}
	}
	logrus.Debugf("'%d' project(s) detected that will create", len(projectsReadyToCreate))
	logrus.Debugf("'%d' project(s) detected that will update", len(projectsForUpdate))

	logrus.Debugf("'%d' project(s) creating", len(projectsReadyToCreate))
	logrus.Debugf("updating %v project(s) ", len(projectsForUpdate))
	orderedByReferenceProjects, err := s.sortProjectsByReference(projectsReadyToCreate)
	if err != nil {
		return err
	}

	logrus.Debugf("'%d' ordered project(s) processing", len(orderedByReferenceProjects))
	for index, project := range orderedByReferenceProjects {

		projectReferences, err := s.getProjectReferences(project)
		if err != nil {
			return err
		}

		project.Specifications.References = projectReferences

		logrus.Debugf("trying to create %v", project.Header.Name)
		if _, err := projectService.Create(project, init); err != nil {
			return err
		}

		fmt.Printf("%v (%d) Project created\n", project.Header.Name, index)

	}

	for _, project := range projectsForUpdate {
		existingProject, err := projectService.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}

		if !reflect.DeepEqual(project.Specifications.Layers, existingProject.Specifications.Layers) {
			newItems := make([]applicationProject.Layer, 0)
			updatedItems := make([]struct {
				Old applicationProject.Layer
				New applicationProject.Layer
			}, 0)
			deletedItems := make([]applicationProject.Layer, 0)

			for _, newLayer := range project.Specifications.Layers {
				found := false
				for _, existingLayer := range existingProject.Specifications.Layers {
					if newLayer.Name == existingLayer.Name {
						found = true
						if !reflect.DeepEqual(newLayer, existingLayer) {
							updatedItems = append(updatedItems, struct {
								Old applicationProject.Layer
								New applicationProject.Layer
							}{
								Old: existingLayer,
								New: newLayer,
							})
						}
						break
					}
				}
				if !found {
					newItems = append(newItems, newLayer)
				}
			}

			for _, existingLayer := range existingProject.Specifications.Layers {
				found := false
				for _, newLayer := range project.Specifications.Layers {
					if newLayer.Name == existingLayer.Name {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, existingLayer)
				}
			}

			if len(newItems) > 0 {

				err := projectService.CreateLayerFolder(project, newItems...)
				if err != nil {
					return err
				}
			}
			if len(updatedItems) > 0 {
				for _, item := range updatedItems {
					err := projectService.DeleteLayerFolder(project, item.Old)
					if err != nil {
						return err
					}
					err = projectService.CreateLayerFolder(project, item.New)
					if err != nil {
						return err
					}
				}
			}
			if len(deletedItems) > 0 {
				err := projectService.DeleteLayerFolder(project, deletedItems...)
				if err != nil {
					return err
				}
			}
		}

		if !reflect.DeepEqual(project.Specifications.Dependencies, existingProject.Specifications.Dependencies) {
			newItems := make([]applicationProject.Dependency, 0)
			updatedItems := make([]struct {
				Old applicationProject.Dependency
				New applicationProject.Dependency
			}, 0)
			deletedItems := make([]applicationProject.Dependency, 0)

			for _, newItem := range project.Specifications.Dependencies {
				found := false
				for _, existingItem := range existingProject.Specifications.Dependencies {
					if newItem.Name == existingItem.Name {
						found = true
						if !reflect.DeepEqual(newItem, existingItem) {
							updatedItems = append(updatedItems, struct {
								Old applicationProject.Dependency
								New applicationProject.Dependency
							}{
								Old: existingItem,
								New: newItem,
							})
						}
						break
					}
				}
				if !found {
					newItems = append(newItems, newItem)
				}
			}

			for _, existingItem := range existingProject.Specifications.Dependencies {
				found := false
				for _, newItem := range project.Specifications.Dependencies {
					if newItem.Name == existingItem.Name {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, existingItem)
				}
			}

			if len(newItems) > 0 {
				err := projectService.AddDependenciesToProject(project, newItems...)
				if err != nil {
					return err
				}

			}
			if len(updatedItems) > 0 {
				for _, item := range updatedItems {
					err := projectService.RemoveDependencyFromProject(project, item.Old)
					if err != nil {
						return err
					}
					err = projectService.AddDependenciesToProject(project, item.New)
					if err != nil {
						return err
					}
				}
			}
			if len(deletedItems) > 0 {
				err := projectService.RemoveDependencyFromProject(project, deletedItems...)
				if err != nil {
					return err
				}
			}
		}

		if !reflect.DeepEqual(project.Specifications.References, existingProject.Specifications.References) {
			newItems := make([]applicationprojectStruct.ProjectBaseStruct, 0)
			updatedItems := make([]struct {
				Old applicationprojectStruct.ProjectBaseStruct
				New applicationprojectStruct.ProjectBaseStruct
			}, 0)
			deletedItems := make([]applicationprojectStruct.ProjectBaseStruct, 0)

			for _, newRef := range project.Specifications.References {
				found := false
				for _, existingRef := range existingProject.Specifications.References {
					if newRef.Header.Name == existingRef.Header.Name {
						found = true
						// if !reflect.DeepEqual(newRef, existingRef) {
						if newRef.Specifications.Name != existingRef.Specifications.Name || newRef.Specifications.Group != existingRef.Specifications.Group || newRef.Specifications.Workspace != existingRef.Specifications.Workspace {
							updatedItems = append(updatedItems, struct {
								Old applicationprojectStruct.ProjectBaseStruct
								New applicationprojectStruct.ProjectBaseStruct
							}{
								Old: existingRef,
								New: newRef,
							})
						}
						break
					}
				}
				if !found {
					newItems = append(newItems, newRef)
				}
			}

			for _, existingRef := range existingProject.Specifications.References {
				found := false
				for _, newRef := range project.Specifications.References {
					if newRef.Header.Name == existingRef.Header.Name {
						found = true
						break
					}
				}
				if !found {
					deletedItems = append(deletedItems, existingRef)
				}
			}

			if len(newItems) > 0 {
				err := projectService.AddReferenceToProject(project, newItems...)
				if err != nil {
					return err
				}
			}
			if len(updatedItems) > 0 {
				for _, item := range updatedItems {
					err := projectService.RemoveReferenceFromProject(project, item.Old)
					if err != nil {
						return err
					}
					err = projectService.AddReferenceToProject(project, item.New)
					if err != nil {
						return err
					}
				}
			}
			if len(deletedItems) > 0 {
				err := projectService.RemoveReferenceFromProject(project, deletedItems...)
				if err != nil {
					return err
				}
			}
		}

		if _, err := projectService.Create(project, false); err != nil {
			return err
		}
	}
	return nil
}
func (s ApplicationProjectEngine) removeProjects(projects []applicationprojectStruct.ProjectBaseStruct, permanent bool) error {

	projectsReadyToDelete := make([]applicationprojectStruct.ProjectBaseStruct, 0)
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	for _, project := range projects {

		ok, err := projectService.IsExists(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}
		if ok {
			projectsReadyToDelete = append(projectsReadyToDelete, project)
		}
	}
	logrus.Debugf("'%d' project(s) detected that will delete", len(projectsReadyToDelete))

	for _, project := range projectsReadyToDelete {

		if _, err := projectService.Remove(project.GetFullName(), project.Specifications.Workspace, false, permanent); err != nil {
			return err
		}

		fmt.Printf("%v Project deleted\n", project.GetFullName())

	}

	logrus.Debugf("'%d' project(s) deleting", len(projectsReadyToDelete))

	return nil
}

func (s ApplicationProjectEngine) completeInformation(ctx *application.ApplicationContext, model *applicationprojectStruct.ProjectBaseStruct) error {

	logrus.Debugf("filling project (%v) information", model.Header.Name)

	if _string.IsEmpty(model.Specifications.Name) {
		model.Specifications.Name = model.Header.Name
	}
	if len(model.Specifications.ProjectIdentifier.Path) == 0 {
		model.Specifications.ProjectIdentifier.Path = append(model.Specifications.ProjectIdentifier.Path, model.Specifications.Name)
	}

	activeWorkspace, err := s.getWorkspace(ctx, *model)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	model.Specifications.Workspace = activeWorkspace.Header.Name
	model.Specifications.WorkspaceObject = activeWorkspace.Specifications.WorkspaceIdentifier
	logrus.Debugf("workspace (%v) detected for (%v)", activeWorkspace.Header.Name, model.Header.Name)

	group, err := s.getGroup(*model)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	model.Specifications.Group = group.Header.Name
	model.Specifications.GroupObject = group.Specifications.GroupIdentifier
	logrus.Debugf("group (%v) detected for (%v)", group.Header.Name, model.Header.Name)

	projectReferences, err := s.getProjectReferences(*model)
	if err != nil {
		return err
	}
	model.Specifications.References = projectReferences
	logrus.Debugf("project references (%d) restored for (%v)", len(model.Specifications.References), model.Specifications.Path)

	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	projectType, err := projectService.GetDefaultPlatformProjectType(*model)
	if err != nil {
		return err
	}

	model.Specifications.Layers = append(model.Specifications.Layers, applicationProject.Layer{})
	model.Specifications.ProjectType = projectType
	return nil
}

func (s ApplicationProjectEngine) getWorkspace(ctx *application.ApplicationContext, project applicationprojectStruct.ProjectBaseStruct) (*workspaceStruct.WorkspaceBaseStruct, error) {
	//TODO: Bu şekilde interface'ten tip dönüşümü tamamlanamadı, yapı buna dönüştürülmeli
	// if projectStruct, ok := project.(project.ProjectBaseStruct); !ok {
	// 	return nil, fmt.Errorf("incompatible model type: expected %T, got %T", project, projectStruct)
	// } else {
	// }

	workspaceName := project.Specifications.Workspace
	if _string.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *workspaceStruct.WorkspaceBaseStruct = nil

	if !_string.IsEmpty(workspaceName) {
		workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
		workspace, err := workspaceService.GetByName(workspaceName)
		if err != nil {
			return nil, err
		}
		if workspace == nil {
			return nil, fmt.Errorf("workspace name (%v) is not correct", workspaceName)
		}
		result = workspace
	} else {
	}

	return result, nil
}
func (s ApplicationProjectEngine) getGroup(project applicationprojectStruct.ProjectBaseStruct) (*group_payload.GroupBaseStruct, error) {
	result := group_payload.GroupBaseStruct{}

	groupName := project.Specifications.Group

	if !_string.IsEmpty(groupName) {
		groupService := ioc.Get[basic_group_contract.GroupInterface]()
		group, err := groupService.GetByName(groupName)
		if err != nil {
			return nil, err
		}
		if group == nil {
			return nil, fmt.Errorf("group name (%v) is not correct", groupName)
		}
		result = *group
	}

	return &result, nil
}

func (s ApplicationProjectEngine) getProjectReferences(prj applicationprojectStruct.ProjectBaseStruct) ([]applicationprojectStruct.ProjectBaseStruct, error) {

	projectReferences := make([]applicationprojectStruct.ProjectBaseStruct, 0)

	for _, reference := range prj.Specifications.References {
		logrus.Debugf("reference (%v) processing for (%v)", reference.Header.Name, prj.Header.Name)

		selectedProject, err := s.getProjectReference(prj, reference)
		if err != nil {
			return nil, err
		}

		if selectedProject != nil {
			projectReferences = append(projectReferences, *selectedProject)
		} else {
			projectReferences = append(projectReferences, reference)
		}
	}

	return projectReferences, nil
}

func (s ApplicationProjectEngine) getProjectReference(prj applicationprojectStruct.ProjectBaseStruct, reference applicationprojectStruct.ProjectBaseStruct) (*applicationprojectStruct.ProjectBaseStruct, error) {

	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	var projectReference *applicationprojectStruct.ProjectBaseStruct = nil

	logrus.Debugf("reference (%v) processing for (%v)", reference.Header.Name, prj.Header.Name)

	if reference.Specifications.ID == 0 {

		logrus.Debugf("found new project defination for (%v) referenced by (%v)", reference.Header.Name, prj.Header.Name)

		if reference.Specifications.WorkspaceObject.ID == 0 {
			if _string.IsEmpty(reference.Specifications.Workspace) {
				reference.Specifications.Workspace = prj.Specifications.Workspace
				logrus.Debugf("decided to using same workspace (%v) for reference (%v)", reference.Specifications.Workspace, reference.Header.Name)
			}
			selectedWorkspace, err := workspaceService.GetByName(reference.Specifications.Workspace)
			if err != nil {
				return nil, err
			}
			reference.Specifications.Workspace = selectedWorkspace.Header.Name
			reference.Specifications.WorkspaceObject = selectedWorkspace.Specifications.WorkspaceIdentifier
			logrus.Debugf("different workspace (%v) for reference (%v)", reference.Specifications.Workspace, reference.Header.Name)
		}

		selectedProject, err := projectService.GetByFullNameWorkspace(reference.GetFullName(), reference.Specifications.Workspace)
		if err != nil {
			return nil, err
		}

		if selectedProject != nil {
			projectReference = selectedProject
		} else {
			projectReference = &reference
		}
	}

	return projectReference, nil
}

func (s ApplicationProjectEngine) sortProjectsByReference(projects []applicationprojectStruct.ProjectBaseStruct) ([]applicationprojectStruct.ProjectBaseStruct, error) {

	projectService := ioc.Get[application_project_contract.ProjectInterface]()

	logrus.Debugf("'%d' projects preparing for ordering", len(projects))
	projectMap := make(map[string]applicationprojectStruct.ProjectSpecification)
	for _, project := range projects {
		projectMap[project.GetUniqueKey()] = project.Specifications
	}
	logrus.Debugf("'%d' project(s) mapped", len(projects))

	for _, project := range projects {

		logrus.Debugf("checking references for project (%v)", project.Header.Name)
		for _, reference := range project.Specifications.References {
			logrus.Debugf("validating reference (%v) for project (%v)", reference.Header.Name, project.Header.Name)
			//TODO: kontrol edilecek, id checkler iptal ediliyor
			if reference.Specifications.ID == 0 {
				if _, ok := projectMap[reference.GetUniqueKey()]; !ok {
					referenceInformation, err := s.getProjectReference(project, reference)
					if err != nil {
						return nil, err
					}

					ok, err := projectService.IsExists(referenceInformation.GetFullName(), referenceInformation.Specifications.Workspace)
					if err != nil {
						return nil, err
					}
					if !ok {
						return nil, errors.New("Invalid Reference in Project '" + project.Header.Name + "'. '" + reference.Header.Name + "' not found.")
					}
				}
			}
		}
	}

	sortedProjectMap := make(map[string]applicationprojectStruct.ProjectBaseStruct)
	sortedProjects, err := sortUnOrderedProjectsByReference(projects, sortedProjectMap)
	if err != nil {
		return nil, err
	}

	// PrintRefInfo(sortedProjects)

	return sortedProjects, nil
}
func sortUnOrderedProjectsByReference(projects []applicationprojectStruct.ProjectBaseStruct, sortedProjectMap map[string]applicationprojectStruct.ProjectBaseStruct) ([]applicationprojectStruct.ProjectBaseStruct, error) {
	projectService := ioc.Get[application_project_contract.ProjectInterface]()
	var sortedProjects []applicationprojectStruct.ProjectBaseStruct = make([]applicationprojectStruct.ProjectBaseStruct, 0)
	var unOrderedProjects []applicationprojectStruct.ProjectBaseStruct = make([]applicationprojectStruct.ProjectBaseStruct, 0)

	logrus.Debugf("'%d' projects ordering", len(projects))
	for _, project := range projects {
		logrus.Debugf("project '%v' processing for order", project.Header.Name)
		if _, ok := sortedProjectMap[project.GetUniqueKey()]; !ok {
			projectReferences := project.Specifications.References
			logrus.Debugf("project '%v' has '%d' references with map key %v", project.Header.Name, len(projectReferences), project.GetUniqueKey())
			if projectReferences == nil || len(projectReferences) == 0 {
				logrus.Debugf("project '%v' has no reference", project.Header.Name)
				sortedProjectMap[project.GetUniqueKey()] = project
				sortedProjects = append(sortedProjects, project)
			} else {
				var allInMap bool = true
				logrus.Debugf("project '%v' has '%d' reference(s)", project.Header.Name, len(projectReferences))
				for _, reference := range projectReferences {
					logrus.Debugf("validating reference (%v) for project (%v)", reference.Header.Name, project.Header.Name)
					if reference.Specifications.ID == 0 {
						if _, ok := sortedProjectMap[reference.GetUniqueKey()]; !ok {
							ok, err := projectService.IsExists(reference.GetFullName(), reference.Specifications.Workspace)
							if err != nil {
								return nil, err
							}
							if !ok { //Bu kontrol buraya gelmeden, "sortProjectsByReference(projects []project.ProjectBaseStruct)" burda da yapılıyor, algoritma iyileştirilebilir
								allInMap = false
								logrus.Debugf("reference (%v) for project (%v), is not in ordered list yet", reference.Header.Name, project.Header.Name)
								break
							}
						}
					}
				}
				if allInMap {
					sortedProjectMap[project.GetUniqueKey()] = project
					sortedProjects = append(sortedProjects, project)
					logrus.Debugf("all references in ordered list for project '%v'. Project is adding to ordered list", project.Header.Name)
				} else {
					unOrderedProjects = append(unOrderedProjects, project)
					logrus.Debugf("all references is not in ordered list for project '%v'. Project is adding to unordered list", project.Header.Name)
				}
			}
		} else {
			logrus.Debugf("project '%v' also ordered", project.Header.Name)
		}
	}
	logrus.Debugf("'%d' project(s) are not ordered", len(unOrderedProjects))

	if len(unOrderedProjects) > 0 {
		sortedChilds, err := sortUnOrderedProjectsByReference(unOrderedProjects, sortedProjectMap)
		if err != nil {
			return nil, err
		}

		sortedProjects = append(sortedProjects, sortedChilds...)
	}

	return sortedProjects, nil
}
