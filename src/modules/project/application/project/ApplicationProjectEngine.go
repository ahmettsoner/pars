package application_project

import (
	"errors"
	"reflect"

	"parsdevkit.net/internal/flowx"

	"fmt"

	"parsdevkit.net/internal/diffx"

	"parsdevkit.net/application/ioc"
	applicationProject "parsdevkit.net/application/structs/project"

	clean_steps "parsdevkit.net/modules/project/application_project/flows/clean"
	create_steps "parsdevkit.net/modules/project/application_project/flows/create"
	remove_steps "parsdevkit.net/modules/project/application_project/flows/remove"
	update_steps "parsdevkit.net/modules/project/application_project/flows/update"
	"parsdevkit.net/modules/project/application_project_contract"

	"parsdevkit.net/modules/group/basic_group_contract"
	basic_group_payload_structs "parsdevkit.net/modules/group/basic_group_payload/structs"
	application_project_payload_structs "parsdevkit.net/modules/project/application_project_payload/structs"
	"parsdevkit.net/modules/workspace/basic_workspace_contract"
	basic_workspace_payload_structs "parsdevkit.net/modules/workspace/basic_workspace_payload/structs"
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
		_, ok := item.(*application_project_payload_structs.ProjectBaseStruct)
		if !ok {
			return false
		}
	}

	return true
}

func (s ApplicationProjectEngine) Process(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.create(ctx, dataStruct, true)
	if err != nil {
		return err
	}
	err = s.update(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}
func (s ApplicationProjectEngine) Destroy(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.remove(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}
func (s ApplicationProjectEngine) prepareToCreate(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	service := ioc.Get[application_project_contract.ProjectInterface]()
	readyToCreateStructs := make([]application_project_payload_structs.ProjectBaseStruct, 0)

	for _, project := range projects {
		if err := s.completeInformation(ctx, &project); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if !ok {
			readyToCreateStructs = append(readyToCreateStructs, project)
		}
	}
	logrus.Debugf("'%d' project(s) detected that will create", len(readyToCreateStructs))

	readyToCreateStructs, err := s.sortProjectsByReference(readyToCreateStructs)
	if err != nil {
		return nil, err
	}

	return readyToCreateStructs, nil
}
func (s ApplicationProjectEngine) create(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct, init bool) error {

	readyToCreateStructs, err := s.prepareToCreate(ctx, projects)
	if err != nil {
		return err
	}

	for _, project := range readyToCreateStructs {

		fmt.Printf("\n\n🛠️  Creating: %s.%s\n\n", project.Header.Name, project.GetKey())

		projectFlow := flowx.NewFlow("CreateNewProject").
			Step(&create_steps.SaveProject{}).
			Step(&create_steps.PrepareProjectFolder{}).
			Step(&create_steps.GenerateProject{}).
			Step(&create_steps.CleanUpFolders{})
		for _, item := range project.Specifications.Layers {
			if _string.IsEmpty(item.Name) {
				continue
			}
			projectFlow.Step(create_steps.NewCreateProjectLayer(item))
		}
		for _, item := range project.Specifications.Dependencies {
			projectFlow.Step(create_steps.NewAddProjectDependency(item))
		}
		for _, item := range project.Specifications.References {
			projectFlow.Step(create_steps.NewAddProjectReference(item))
		}

		fc := flowx.NewContextWithData(map[string]any{
			"init":    init,
			"project": project,
		})

		if err := projectFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Project Create işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s ApplicationProjectEngine) prepareToUpdate(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	service := ioc.Get[application_project_contract.ProjectInterface]()
	readyToUpdateStructs := make([]application_project_payload_structs.ProjectBaseStruct, 0)

	for _, project := range projects {
		if err := s.completeInformation(ctx, &project); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			newModelHash, err := encrypt.CalculateHashFromObject(project)
			if err != nil {
				return nil, err
			}
			structHash, err := service.GetHash(project.GetFullName(), project.Specifications.Workspace)
			if err != nil {
				return nil, err
			}

			if newModelHash != structHash {
				readyToUpdateStructs = append(readyToUpdateStructs, project)
			}
			//burda else ile kayıt zaten güncel bilgisi yazdırılabilir
		}
		//burda else ile kayıt bulunamadı bilgisi yazdırılabilir
	}
	logrus.Debugf("'%d' project(s) detected that will update", len(readyToUpdateStructs))

	return readyToUpdateStructs, nil
}
func (s ApplicationProjectEngine) update(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct, init bool) error {

	service := ioc.Get[application_project_contract.ProjectInterface]()
	readyToUpdateStructs, err := s.prepareToUpdate(ctx, projects)
	if err != nil {
		return err
	}
	for _, project := range readyToUpdateStructs {

		fmt.Printf("\n\n🛠️  Updating: %s.%s\n\n", project.Header.Name, project.GetKey())

		existingProject, err := service.GetByFullNameWorkspace(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return err
		}

		projectFlow := flowx.NewFlow("UpdateExistingProject").
			Step(&update_steps.UpdateProject{})

		if !reflect.DeepEqual(project.Specifications.Layers, existingProject.Specifications.Layers) {

			result := diffx.DiffSlice(existingProject.Specifications.Layers, project.Specifications.Layers)

			if len(result.Created) > 0 {

				for _, item := range result.Created {
					if _string.IsEmpty(item.Name) {
						continue
					}
					projectFlow.Step(update_steps.NewCreateProjectLayer(item))
				}
			}
			if len(result.Updated) > 0 {
				//TODO Burda değişiklik tespit edilerek eğer move ve rename yapılabilir dosya ve klasörlere, silmekten daha güvenli
				for _, item := range result.Updated {
					if _string.IsEmpty(item.Old.Name) {
						continue
					}
					projectFlow.Step(update_steps.NewUpdateProjectLayer(item.Old, item.New))
				}
			}
			if len(result.Deleted) > 0 {
				for _, item := range result.Deleted {
					if _string.IsEmpty(item.Name) {
						continue
					}
					projectFlow.Step(update_steps.NewDeleteProjectLayer(item))
				}
			}
		}

		if !reflect.DeepEqual(project.Specifications.Dependencies, existingProject.Specifications.Dependencies) {
			result := diffx.DiffSlice(existingProject.Specifications.Dependencies, project.Specifications.Dependencies)

			if len(result.Created) > 0 {
				for _, item := range result.Created {
					projectFlow.Step(update_steps.NewCreateProjectDependency(item))
				}

			}
			if len(result.Updated) > 0 {
				for _, item := range result.Updated {
					projectFlow.Step(update_steps.NewUpdateProjectDependency(item.Old, item.New))
				}
			}
			if len(result.Deleted) > 0 {
				for _, item := range result.Deleted {
					projectFlow.Step(update_steps.NewDeleteProjectDependency(item))
				}
			}
		}

		if !reflect.DeepEqual(project.Specifications.References, existingProject.Specifications.References) {
			result := diffx.DiffSlice(existingProject.Specifications.References, project.Specifications.References)

			if len(result.Created) > 0 {
				for _, item := range result.Created {
					projectFlow.Step(update_steps.NewCreateProjectReference(item))
				}
			}
			if len(result.Updated) > 0 {
				for _, item := range result.Updated {
					projectFlow.Step(update_steps.NewUpdateProjectReference(item.Old, item.New))
				}
			}
			if len(result.Deleted) > 0 {
				for _, item := range result.Deleted {
					projectFlow.Step(update_steps.NewDeleteProjectReference(item))
				}
			}
		}
		fc := flowx.NewContextWithData(map[string]any{
			"init":    init,
			"project": project,
		})

		if err := projectFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Project Update işleminde hata oluştu: %w", &err)
		}
	}
	return nil
}

func (s ApplicationProjectEngine) prepareToRemove(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	service := ioc.Get[application_project_contract.ProjectInterface]()
	readyToRemoveStructs := make([]application_project_payload_structs.ProjectBaseStruct, 0)

	for _, project := range projects {
		if err := s.completeInformation(ctx, &project); err != nil {
			return nil, err
		}
		ok, err := service.IsExists(project.GetFullName(), project.Specifications.Workspace)
		if err != nil {
			return nil, err
		}
		if ok {
			readyToRemoveStructs = append(readyToRemoveStructs, project)
		}
	}
	logrus.Debugf("'%d' project(s) detected that will remove", len(readyToRemoveStructs))

	return readyToRemoveStructs, nil
}
func (s ApplicationProjectEngine) remove(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct, permanent bool) error {

	readyToRemoveStructs, err := s.prepareToRemove(ctx, projects)
	if err != nil {
		return err
	}

	for _, project := range readyToRemoveStructs {

		fmt.Printf("\n\n🛠️  Removing: %s.%s\n\n", project.Header.Name, project.GetKey())

		projectFlow := flowx.NewFlow("RemoveExistingProject").
			Step(&remove_steps.DestroyProject{}).
			Step(&remove_steps.DeleteProject{}).
			Step(&remove_steps.RemoveProjectFiles{})

		fc := flowx.NewContextWithData(map[string]any{
			"permanent": permanent,
			"project":   project,
		})

		if err := projectFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Project Remove işleminde hata oluştu: %w", &err)
		}
	}

	return nil
}

func (s ApplicationProjectEngine) completeInformation(ctx *application.ApplicationContext, model *application_project_payload_structs.ProjectBaseStruct) error {

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

	service := ioc.Get[application_project_contract.ProjectInterface]()
	projectType, err := service.GetDefaultPlatformProjectType(*model)
	if err != nil {
		return err
	}

	model.Specifications.Layers = append(model.Specifications.Layers, applicationProject.Layer{})
	model.Specifications.ProjectType = projectType
	return nil
}

func (s ApplicationProjectEngine) getWorkspace(ctx *application.ApplicationContext, project application_project_payload_structs.ProjectBaseStruct) (*basic_workspace_payload_structs.WorkspaceBaseStruct, error) {
	//TODO: Bu şekilde interface'ten tip dönüşümü tamamlanamadı, yapı buna dönüştürülmeli
	// if projectStruct, ok := project.(project.ProjectBaseStruct); !ok {
	// 	return nil, fmt.Errorf("incompatible model type: expected %T, got %T", project, projectStruct)
	// } else {
	// }

	workspaceName := project.Specifications.Workspace
	if _string.IsEmpty(workspaceName) {
		workspaceName = ctx.CurrentWorkspace.Name
	}

	var result *basic_workspace_payload_structs.WorkspaceBaseStruct = nil

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
func (s ApplicationProjectEngine) getGroup(project application_project_payload_structs.ProjectBaseStruct) (*basic_group_payload_structs.GroupBaseStruct, error) {
	result := basic_group_payload_structs.GroupBaseStruct{}

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

func (s ApplicationProjectEngine) getProjectReferences(prj application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	projectReferences := make([]application_project_payload_structs.ProjectBaseStruct, 0)

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

func (s ApplicationProjectEngine) getProjectReference(prj application_project_payload_structs.ProjectBaseStruct, reference application_project_payload_structs.ProjectBaseStruct) (*application_project_payload_structs.ProjectBaseStruct, error) {

	workspaceService := ioc.Get[basic_workspace_contract.WorkspaceInterface]()
	service := ioc.Get[application_project_contract.ProjectInterface]()
	var projectReference *application_project_payload_structs.ProjectBaseStruct = nil

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

		selectedProject, err := service.GetByFullNameWorkspace(reference.GetFullName(), reference.Specifications.Workspace)
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

func (s ApplicationProjectEngine) sortProjectsByReference(projects []application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	service := ioc.Get[application_project_contract.ProjectInterface]()

	logrus.Debugf("'%d' projects preparing for ordering", len(projects))
	projectMap := make(map[string]application_project_payload_structs.ProjectSpecification)
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

					ok, err := service.IsExists(referenceInformation.GetFullName(), referenceInformation.Specifications.Workspace)
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

	sortedProjectMap := make(map[string]application_project_payload_structs.ProjectBaseStruct)
	sortedProjects, err := s.sortUnOrderedProjectsByReference(projects, sortedProjectMap)
	if err != nil {
		return nil, err
	}

	// PrintRefInfo(sortedProjects)

	return sortedProjects, nil
}
func (s ApplicationProjectEngine) sortUnOrderedProjectsByReference(projects []application_project_payload_structs.ProjectBaseStruct, sortedProjectMap map[string]application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {
	service := ioc.Get[application_project_contract.ProjectInterface]()
	var sortedProjects []application_project_payload_structs.ProjectBaseStruct = make([]application_project_payload_structs.ProjectBaseStruct, 0)
	var unOrderedProjects []application_project_payload_structs.ProjectBaseStruct = make([]application_project_payload_structs.ProjectBaseStruct, 0)

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
							ok, err := service.IsExists(reference.GetFullName(), reference.Specifications.Workspace)
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
		sortedChilds, err := s.sortUnOrderedProjectsByReference(unOrderedProjects, sortedProjectMap)
		if err != nil {
			return nil, err
		}

		sortedProjects = append(sortedProjects, sortedChilds...)
	}

	return sortedProjects, nil
}

func (s ApplicationProjectEngine) GetConfig() engines.EngineConfig {
	return engines.EngineConfig{
		Name:  application_project_payload_structs.MODULE_KEY,
		Order: 2000,
	}
}

func CastArrayToConcrate(data []schemas.SchemaInterface) ([]application_project_payload_structs.ProjectBaseStruct, error) {
	r := make([]application_project_payload_structs.ProjectBaseStruct, 0, len(data))

	for _, item := range data {
		model, ok := item.(*application_project_payload_structs.ProjectBaseStruct)
		if !ok {
			return nil, fmt.Errorf("invalid item type: expected application_project_payload_structs.ProjectBaseStruct, got %T", item)
		}

		r = append(r, *model)
	}

	return r, nil
}

func (s ApplicationProjectEngine) Clean(ctx *application.ApplicationContext, data []schemas.SchemaInterface) error {
	dataStruct, err := CastArrayToConcrate(data)
	if err != nil {
		return err
	}

	err = s.clean(ctx, dataStruct, true)
	if err != nil {
		return err
	}

	return nil
}

func (s ApplicationProjectEngine) clean(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct, init bool) error {

	readyToUpdateStructs, err := s.prepareToClean(ctx, projects)
	if err != nil {
		return err
	}
	for _, project := range readyToUpdateStructs {

		fmt.Printf("\n\n🛠️  Cleaning: %s.%s\n\n", project.Header.Name, project.GetKey())

		projectFlow := flowx.NewFlow("CleanProject").
			Step(&clean_steps.CleanProject{})

		fc := flowx.NewContextWithData(map[string]any{
			"project": project,
		})

		if err := projectFlow.Run(fc); err != nil {
			fc.Log("Flow failed: %v", err)
			return fmt.Errorf("xxx: Project Update işleminde hata oluştu: %w", &err)
		}
	}
	return nil
}

func (s ApplicationProjectEngine) prepareToClean(ctx *application.ApplicationContext, projects []application_project_payload_structs.ProjectBaseStruct) ([]application_project_payload_structs.ProjectBaseStruct, error) {

	service := ioc.Get[application_project_contract.ProjectInterface]()
	readyToUpdateStructs := make([]application_project_payload_structs.ProjectBaseStruct, 0)

	for _, project := range projects {

		activeWorkspace, err := s.getWorkspace(ctx, project)
		if err != nil {
			return nil, err
		}

		existingProject, err := service.GetByFullNameWorkspace(project.GetFullName(), activeWorkspace.Header.Name)
		if err != nil {
			return nil, err
		}
		readyToUpdateStructs = append(readyToUpdateStructs, *existingProject)
		//burda else ile kayıt bulunamadı bilgisi yazdırılabilir
	}
	logrus.Debugf("'%d' project(s) detected that will update", len(readyToUpdateStructs))

	return readyToUpdateStructs, nil
}
