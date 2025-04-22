package applicationProject

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	applicationproject "parsdevkit.net/structs/project/application-project"

	"parsdevkit.net/operation/services"

	"parsdevkit.net/core/utils"

	"github.com/sirupsen/logrus"
)

type ApplicationProjectEngine struct{}

func (s ApplicationProjectEngine) CreateProjectsFromTemplate(workspaceName string, init bool, data any, templateFiles ...string) error {

	var allProjects []applicationproject.ProjectBaseStruct = make([]applicationproject.ProjectBaseStruct, 0)

	logrus.Debugf("found %v template(s) to create project", len(templateFiles))
	for _, templateFilePath := range templateFiles {

		var tmplFile = filepath.Join(utils.GetManagerTemplatesLocation(), templateFilePath)
		tmplContent, err := os.ReadFile(tmplFile)
		if err != nil {
			log.Fatal(err)
		}
		var outputBuffer bytes.Buffer
		err = template.Must(template.New("ProjectFromTemplate").Parse(string(tmplContent))).Execute(&outputBuffer, data)
		if err != nil {
			log.Fatal(err)
		}
		mainStr := outputBuffer.String()

		projectSerializer := ApplicationProjectSerializer{}
		projects, err := projectSerializer.GetProjectStuctsFromString(workspaceName, mainStr)
		if err != nil {
			return err
		}
		allProjects = append(allProjects, projects...)
	}
	logrus.Debugf("found %v project(s) to create", len(allProjects))
	if err := s.CreateProjects(allProjects, init); err != nil {
		return err
	}

	return nil
}
func (s ApplicationProjectEngine) CreateProjectsFromFile(workspaceName string, init bool, files ...string) error {
	logrus.Debugf("found %v file(s) to create project", len(files))
	if len(files) > 0 {

		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {

				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		projectSerializer := ApplicationProjectSerializer{}
		projectsFromFile, err := projectSerializer.GetProjectStuctsFromFile(workspaceName, allFiles...)
		if err != nil {
			return err
		}

		logrus.Debugf("found %v project(s) to create", len(projectsFromFile))
		if err := s.CreateProjects(projectsFromFile, init); err != nil {
			return err
		}
	}
	return nil
}

func (s ApplicationProjectEngine) RemoveProjectsFromFile(workspaceName string, permanent bool, files ...string) error {
	logrus.Debugf("found %v file(s) to create project", len(files))

	if len(files) > 0 {
		allFiles := make([]string, 0)
		for _, path := range files {
			if !utils.IsEmpty(path) {

				files, err := utils.GetFilesInPath(path)
				if err != nil {
					return err
				}

				allFiles = append(allFiles, files...)
			}
		}

		projectSerializer := ApplicationProjectSerializer{}
		projectsFromFile, err := projectSerializer.GetProjectStuctsFromFile(workspaceName, allFiles...)
		if err != nil {
			return err
		}

		logrus.Debugf("found %v project(s) to remove", len(projectsFromFile))
		if err := s.RemoveProjects(projectsFromFile, permanent); err != nil {
			return err
		}
	}
	return nil
}

func (s ApplicationProjectEngine) CreateProjects(projects []applicationproject.ProjectBaseStruct, init bool) error {

	projectsReadyToCreate := make([]applicationproject.ProjectBaseStruct, 0)
	projectsForUpdate := make([]applicationproject.ProjectBaseStruct, 0)
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	// projectReferenceMap := make(map[string]map[string]applicationproject.ProjectSpecification)
	for _, project := range projects {
		if ok := projectService.IsExists(project.GetFullName(), project.Specifications.Workspace); ok {

			newModelHash, err := utils.CalculateHashFromObject(project)
			if err != nil {
				return err
			}
			structHash := projectService.GetHash(project.GetFullName(), project.Specifications.Workspace)

			if newModelHash != structHash {
				projectsForUpdate = append(projectsForUpdate, project)
			}
		} else {
			projectsReadyToCreate = append(projectsReadyToCreate, project)
			// for _, reference := range applicationproject.Specifications.Configuration.References {
			// 	if _, ok := projectReferenceMap[applicationproject.Specifications.GetUniqueKey()][reference.GetUniqueKey()]; ok {
			// 		continue
			// 	} else {
			// 		if ok := projectService.IsExists(reference.GetFullName(), reference.Workspace.Name); ok {
			// 			projectReferenceMap[applicationproject.Specifications.GetUniqueKey()][reference.GetUniqueKey()] = reference
			// 		}
			// 	}
			// }
		}
	}
	logrus.Debugf("'%d' project(s) detected that will create", len(projectsReadyToCreate))
	logrus.Debugf("'%d' project(s) detected that will update", len(projectsForUpdate))

	logrus.Debugf("'%d' project(s) creating", len(projectsReadyToCreate))
	logrus.Debugf("updating %v project(s) ", len(projectsForUpdate))
	orderedByReferenceProjects, err := s.SortProjectsByReference(projectsReadyToCreate)
	if err != nil {
		return err
	}

	logrus.Debugf("'%d' ordered project(s) processing", len(orderedByReferenceProjects))
	for index, project := range orderedByReferenceProjects {

		projectSerializer := ApplicationProjectSerializer{}
		projectReferences, err := projectSerializer.GetProjectReferences(project)
		if err != nil {
			return err
		}

		project.Specifications.Configuration.References = projectReferences

		logrus.Debugf("trying to create %v", project.Name)
		if _, err := projectService.Create(project, init); err != nil {
			return err
		}

		fmt.Printf("%v (%d) Project created\n", project.Name, index)

	}

	for _, project := range projectsForUpdate {
		existingProject, err := projectService.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}

		if !reflect.DeepEqual(project.Specifications.Configuration.Layers, existingProject.Specifications.Configuration.Layers) {
			newItems := make([]applicationproject.Layer, 0)
			updatedItems := make([]struct {
				Old applicationproject.Layer
				New applicationproject.Layer
			}, 0)
			deletedItems := make([]applicationproject.Layer, 0)

			for _, newLayer := range project.Specifications.Configuration.Layers {
				found := false
				for _, existingLayer := range existingProject.Specifications.Configuration.Layers {
					if newLayer.Name == existingLayer.Name {
						found = true
						if !reflect.DeepEqual(newLayer, existingLayer) {
							updatedItems = append(updatedItems, struct {
								Old applicationproject.Layer
								New applicationproject.Layer
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
			newItems := make([]applicationproject.Package, 0)
			updatedItems := make([]struct {
				Old applicationproject.Package
				New applicationproject.Package
			}, 0)
			deletedItems := make([]applicationproject.Package, 0)

			for _, newItem := range project.Specifications.Configuration.Dependencies {
				found := false
				for _, existingItem := range existingProject.Specifications.Configuration.Dependencies {
					if newItem.Name == existingItem.Name {
						found = true
						if !reflect.DeepEqual(newItem, existingItem) {
							updatedItems = append(updatedItems, struct {
								Old applicationproject.Package
								New applicationproject.Package
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
			newItems := make([]applicationproject.ProjectBaseStruct, 0)
			updatedItems := make([]struct {
				Old applicationproject.ProjectBaseStruct
				New applicationproject.ProjectBaseStruct
			}, 0)
			deletedItems := make([]applicationproject.ProjectBaseStruct, 0)

			for _, newRef := range project.Specifications.Configuration.References {
				found := false
				for _, existingRef := range existingProject.Specifications.Configuration.References {
					if newRef.Name == existingRef.Name {
						found = true
						// if !reflect.DeepEqual(newRef, existingRef) {
						if newRef.Specifications.Name != existingRef.Specifications.Name || newRef.Specifications.Group != existingRef.Specifications.Group || newRef.Specifications.Workspace != existingRef.Specifications.Workspace {
							updatedItems = append(updatedItems, struct {
								Old applicationproject.ProjectBaseStruct
								New applicationproject.ProjectBaseStruct
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
					if newRef.Name == existingRef.Name {
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
func (s ApplicationProjectEngine) RemoveProjects(projects []applicationproject.ProjectBaseStruct, permanent bool) error {

	projectsReadyToDelete := make([]applicationproject.ProjectBaseStruct, 0)
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	for _, project := range projects {
		if ok := projectService.IsExists(project.GetFullName(), project.Specifications.Workspace); ok {
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

func (s ApplicationProjectEngine) SortProjectsByReference(projects []applicationproject.ProjectBaseStruct) ([]applicationproject.ProjectBaseStruct, error) {

	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	logrus.Debugf("'%d' projects preparing for ordering", len(projects))
	projectMap := make(map[string]applicationproject.ProjectSpecification)
	for _, project := range projects {
		projectMap[project.GetUniqueKey()] = project.Specifications
	}
	logrus.Debugf("'%d' project(s) mapped", len(projects))

	for _, project := range projects {

		logrus.Debugf("checking references for project (%v)", project.Name)
		for _, reference := range project.Specifications.Configuration.References {
			logrus.Debugf("validating reference (%v) for project (%v)", reference.Name, project.Name)
			//TODO: kontrol edilecek, id checkler iptal ediliyor
			if reference.Specifications.ID == 0 {
				if _, ok := projectMap[reference.GetUniqueKey()]; !ok {
					projectSerializer := ApplicationProjectSerializer{}
					referenceInformation, err := projectSerializer.GetProjectReference(project, reference)
					if err != nil {
						return nil, err
					}
					if ok := projectService.IsExists(referenceInformation.GetFullName(), referenceInformation.Specifications.Workspace); !ok {
						return nil, errors.New("Invalid Reference in Project '" + project.Name + "'. '" + reference.Name + "' not found.")
					}
				}
			}
		}
	}

	sortedProjectMap := make(map[string]applicationproject.ProjectBaseStruct)
	var sortedProjects []applicationproject.ProjectBaseStruct = SortUnOrderedProjectsByReference(projects, sortedProjectMap)

	// PrintRefInfo(sortedProjects)

	return sortedProjects, nil
}
func SortUnOrderedProjectsByReference(projects []applicationproject.ProjectBaseStruct, sortedProjectMap map[string]applicationproject.ProjectBaseStruct) []applicationproject.ProjectBaseStruct {
	projectService := services.NewApplicationProjectService(utils.GetEnvironment())
	var sortedProjects []applicationproject.ProjectBaseStruct = make([]applicationproject.ProjectBaseStruct, 0)
	var unOrderedProjects []applicationproject.ProjectBaseStruct = make([]applicationproject.ProjectBaseStruct, 0)

	logrus.Debugf("'%d' projects ordering", len(projects))
	for _, project := range projects {
		logrus.Debugf("project '%v' processing for order", project.Name)
		if _, ok := sortedProjectMap[project.GetUniqueKey()]; !ok {
			projectReferences := project.Specifications.Configuration.References
			logrus.Debugf("project '%v' has '%d' references with map key %v", project.Name, len(projectReferences), project.GetUniqueKey())
			if projectReferences == nil || len(projectReferences) == 0 {
				logrus.Debugf("project '%v' has no reference", project.Name)
				sortedProjectMap[project.GetUniqueKey()] = project
				sortedProjects = append(sortedProjects, project)
			} else {
				var allInMap bool = true
				logrus.Debugf("project '%v' has '%d' reference(s)", project.Name, len(projectReferences))
				for _, reference := range projectReferences {
					logrus.Debugf("validating reference (%v) for project (%v)", reference.Name, project.Name)
					if reference.Specifications.ID == 0 {
						if _, ok := sortedProjectMap[reference.GetUniqueKey()]; !ok {
							if ok := projectService.IsExists(reference.GetFullName(), reference.Specifications.Workspace); !ok { //Bu kontrol buraya gelmeden, "SortProjectsByReference(projects []project.ProjectBaseStruct)" burda da yapılıyor, algoritma iyileştirilebilir
								allInMap = false
								logrus.Debugf("reference (%v) for project (%v), is not in ordered list yet", reference.Name, project.Name)
								break
							}
						}
					}
				}
				if allInMap {
					sortedProjectMap[project.GetUniqueKey()] = project
					sortedProjects = append(sortedProjects, project)
					logrus.Debugf("all references in ordered list for project '%v'. Project is adding to ordered list", project.Name)
				} else {
					unOrderedProjects = append(unOrderedProjects, project)
					logrus.Debugf("all references is not in ordered list for project '%v'. Project is adding to unordered list", project.Name)
				}
			}
		} else {
			logrus.Debugf("project '%v' also ordered", project.Name)
		}
	}
	logrus.Debugf("'%d' project(s) are not ordered", len(unOrderedProjects))

	if len(unOrderedProjects) > 0 {
		var sortedChilds = SortUnOrderedProjectsByReference(unOrderedProjects, sortedProjectMap)
		sortedProjects = append(sortedProjects, sortedChilds...)
	}

	return sortedProjects
}
