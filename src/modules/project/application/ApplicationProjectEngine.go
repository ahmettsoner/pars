package application

import (
	"errors"
	"reflect"

	"fmt"

	groupStruct "parsdevkit.net/structs/group"
	applicationprojectStruct "parsdevkit.net/structs/project/application-project"
	workspaceStruct "parsdevkit.net/structs/workspace"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
	"parsdevkit.net/core"
	"parsdevkit.net/core/schemas"
)

type ApplicationProjectEngine struct{}

func (s ApplicationProjectEngine) Validate(data []schemas.Schema) bool {
	for _, item := range data {
		_, ok := item.(*applicationprojectStruct.ProjectBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}
func (s ApplicationProjectEngine) Process(ctx *core.ApplicationContext, data []schemas.Schema) error {
	applicationprojects := make([]applicationprojectStruct.ProjectBaseStruct, 0, len(data))

	for _, item := range data {
		applicationproject, ok := item.(*applicationprojectStruct.ProjectBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Process: expected applicationprojectStruct.ProjectBaseStruct, got %T", item)
		}

		if err := s.completeProjectInformation(ctx, applicationproject); err != nil {
			return err
		}
		applicationprojects = append(applicationprojects, *applicationproject)
	}

	return s.createProjects(applicationprojects, true)
}
func (s ApplicationProjectEngine) Destroy(ctx *core.ApplicationContext, data []schemas.Schema) error {
	applicationprojects := make([]applicationprojectStruct.ProjectBaseStruct, 0, len(data))

	for _, item := range data {
		applicationproject, ok := item.(*applicationprojectStruct.ProjectBaseStruct)
		if !ok {
			return fmt.Errorf("invalid item type in Destroy: expected applicationprojectStruct.ProjectBaseStruct, got %T", item)
		}

		if err := s.completeProjectInformation(ctx, applicationproject); err != nil {
			return err
		}
		applicationprojects = append(applicationprojects, *applicationproject)
	}

	return s.removeProjects(applicationprojects, true)
}

func (s ApplicationProjectEngine) createProjects(projects []applicationprojectStruct.ProjectBaseStruct, init bool) error {

	projectsReadyToCreate := make([]applicationprojectStruct.ProjectBaseStruct, 0)
	projectsForUpdate := make([]applicationprojectStruct.ProjectBaseStruct, 0)
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	// projectReferenceMap := make(map[string]map[string]applicationprojectStruct.ProjectSpecification)

	for _, project := range projects {
		ok, err := projectService.IsExists(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}
		if ok {
			newModelHash, err := utils.CalculateHashFromObject(project)
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
			// for _, reference := range applicationprojectStruct.Specifications.Configuration.References {
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

		project.Specifications.Configuration.References = projectReferences

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

		if !reflect.DeepEqual(project.Specifications.Configuration.Layers, existingProject.Specifications.Configuration.Layers) {
			newItems := make([]applicationprojectStruct.Layer, 0)
			updatedItems := make([]struct {
				Old applicationprojectStruct.Layer
				New applicationprojectStruct.Layer
			}, 0)
			deletedItems := make([]applicationprojectStruct.Layer, 0)

			for _, newLayer := range project.Specifications.Configuration.Layers {
				found := false
				for _, existingLayer := range existingProject.Specifications.Configuration.Layers {
					if newLayer.Name == existingLayer.Name {
						found = true
						if !reflect.DeepEqual(newLayer, existingLayer) {
							updatedItems = append(updatedItems, struct {
								Old applicationprojectStruct.Layer
								New applicationprojectStruct.Layer
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

			for _, existingLayer := range existingProject.Specifications.Configuration.Layers {
				found := false
				for _, newLayer := range project.Specifications.Configuration.Layers {
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

		if !reflect.DeepEqual(project.Specifications.Configuration.Dependencies, existingProject.Specifications.Configuration.Dependencies) {
			newItems := make([]applicationprojectStruct.Package, 0)
			updatedItems := make([]struct {
				Old applicationprojectStruct.Package
				New applicationprojectStruct.Package
			}, 0)
			deletedItems := make([]applicationprojectStruct.Package, 0)

			for _, newItem := range project.Specifications.Configuration.Dependencies {
				found := false
				for _, existingItem := range existingProject.Specifications.Configuration.Dependencies {
					if newItem.Name == existingItem.Name {
						found = true
						if !reflect.DeepEqual(newItem, existingItem) {
							updatedItems = append(updatedItems, struct {
								Old applicationprojectStruct.Package
								New applicationprojectStruct.Package
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

			for _, existingItem := range existingProject.Specifications.Configuration.Dependencies {
				found := false
				for _, newItem := range project.Specifications.Configuration.Dependencies {
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
				err := projectService.AddPackageToProject(project, newItems...)
				if err != nil {
					return err
				}

			}
			if len(updatedItems) > 0 {
				for _, item := range updatedItems {
					err := projectService.RemovePackageToProject(project, item.Old)
					if err != nil {
						return err
					}
					err = projectService.AddPackageToProject(project, item.New)
					if err != nil {
						return err
					}
				}
			}
			if len(deletedItems) > 0 {
				err := projectService.RemovePackageToProject(project, deletedItems...)
				if err != nil {
					return err
				}
			}
		}

		if !reflect.DeepEqual(project.Specifications.Configuration.References, existingProject.Specifications.Configuration.References) {
			newItems := make([]applicationprojectStruct.ProjectBaseStruct, 0)
			updatedItems := make([]struct {
				Old applicationprojectStruct.ProjectBaseStruct
				New applicationprojectStruct.ProjectBaseStruct
			}, 0)
			deletedItems := make([]applicationprojectStruct.ProjectBaseStruct, 0)

			for _, newRef := range project.Specifications.Configuration.References {
				found := false
				for _, existingRef := range existingProject.Specifications.Configuration.References {
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

			for _, existingRef := range existingProject.Specifications.Configuration.References {
				found := false
				for _, newRef := range project.Specifications.Configuration.References {
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
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
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

func (s ApplicationProjectEngine) completeProjectInformation(ctx *core.ApplicationContext, project *applicationprojectStruct.ProjectBaseStruct) error {

	logrus.Debugf("filling project (%v) information", project.Header.Name)

	activeWorkspace, err := s.getWorkspace(ctx, *project)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	project.Specifications.Workspace = activeWorkspace.Header.Name
	project.Specifications.WorkspaceObject = activeWorkspace.Specifications
	logrus.Debugf("workspace (%v) detected for (%v)", activeWorkspace.Header.Name, project.Header.Name)

	group, err := s.getGroup(*project)
	if err != nil {
		return err
	}

	//WARN: Doğru mu oldu?
	project.Specifications.Group = group.Header.Name
	project.Specifications.GroupObject = *&group.Specifications
	logrus.Debugf("group (%v) detected for (%v)", group.Header.Name, project.Header.Name)

	projectReferences, err := s.getProjectReferences(*project)
	if err != nil {
		return err
	}
	project.Specifications.Configuration.References = projectReferences
	logrus.Debugf("project references (%d) restored for (%v)", len(project.Specifications.Configuration.References), project.Specifications.Path)

	return nil
}

func (s ApplicationProjectEngine) getWorkspace(ctx *core.ApplicationContext, project applicationprojectStruct.ProjectBaseStruct) (*workspaceStruct.WorkspaceBaseStruct, error) {
	//TODO: Bu şekilde interface'ten tip dönüşümü tamamlanamadı, yapı buna dönüştürülmeli
	// if projectStruct, ok := project.(project.ProjectBaseStruct); !ok {
	// 	return nil, fmt.Errorf("incompatible model type: expected %T, got %T", project, projectStruct)
	// } else {
	// }

	workspaceName := project.Specifications.Workspace
	if utils.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *workspaceStruct.WorkspaceBaseStruct = nil

	if !utils.IsEmpty(workspaceName) {
		workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
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
func (s ApplicationProjectEngine) getGroup(project applicationprojectStruct.ProjectBaseStruct) (*groupStruct.GroupBaseStruct, error) {
	result := groupStruct.GroupBaseStruct{}

	groupName := project.Specifications.Group

	if !utils.IsEmpty(groupName) {
		groupService := services.NewGroupService(utils.GetEnvironment())
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

	for _, reference := range prj.Specifications.Configuration.References {
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

	workspaceService := services.NewWorkspaceService(utils.GetEnvironment())
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	var projectReference *applicationprojectStruct.ProjectBaseStruct = nil

	logrus.Debugf("reference (%v) processing for (%v)", reference.Header.Name, prj.Header.Name)

	if reference.Specifications.ID == 0 {

		logrus.Debugf("found new project defination for (%v) referenced by (%v)", reference.Header.Name, prj.Header.Name)

		if reference.Specifications.WorkspaceObject.ID == 0 {
			if utils.IsEmpty(reference.Specifications.Workspace) {
				reference.Specifications.Workspace = prj.Specifications.Workspace
				logrus.Debugf("decided to using same workspace (%v) for reference (%v)", reference.Specifications.Workspace, reference.Header.Name)
			}
			selectedWorkspace, err := workspaceService.GetByName(reference.Specifications.Workspace)
			if err != nil {
				return nil, err
			}
			reference.Specifications.Workspace = selectedWorkspace.Header.Name
			reference.Specifications.WorkspaceObject = selectedWorkspace.Specifications
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

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	logrus.Debugf("'%d' projects preparing for ordering", len(projects))
	projectMap := make(map[string]applicationprojectStruct.ProjectSpecification)
	for _, project := range projects {
		projectMap[project.GetUniqueKey()] = project.Specifications
	}
	logrus.Debugf("'%d' project(s) mapped", len(projects))

	for _, project := range projects {

		logrus.Debugf("checking references for project (%v)", project.Header.Name)
		for _, reference := range project.Specifications.Configuration.References {
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
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	var sortedProjects []applicationprojectStruct.ProjectBaseStruct = make([]applicationprojectStruct.ProjectBaseStruct, 0)
	var unOrderedProjects []applicationprojectStruct.ProjectBaseStruct = make([]applicationprojectStruct.ProjectBaseStruct, 0)

	logrus.Debugf("'%d' projects ordering", len(projects))
	for _, project := range projects {
		logrus.Debugf("project '%v' processing for order", project.Header.Name)
		if _, ok := sortedProjectMap[project.GetUniqueKey()]; !ok {
			projectReferences := project.Specifications.Configuration.References
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
