package repositories

import (
	"errors"
	"fmt"
	"strings"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/pkg/utilities/encrypt"

	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	DbContext *contexts.DbContext
}

func NewProjectRepository(dbCtx *contexts.DbContext) *ProjectRepository {
	return &ProjectRepository{DbContext: dbCtx}
}

func (s *ProjectRepository) Get(id int) (*entities.Project, error) {
	entity := new(entities.Project)
	result := s.DbContext.Database.First(entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return entity, nil
}
func (s *ProjectRepository) GetByName(name string) (*entities.Project, error) {
	entity := new(entities.Project)
	result := s.DbContext.Database.Where("json_extract(document, '$.Header.Name')= ?", name).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}
func (s *ProjectRepository) List() (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *ProjectRepository) Save(entity *entities.Project) error {

	existingValue, err := s.GetByName(entity.Name)
	if err != nil {
		return err
	}

	documentHash, err := encrypt.CalculateHash(entity.Document)
	if err != nil {
		return err
	}
	entity.Hash = documentHash
	if existingValue == nil {
		result := s.DbContext.Database.Create(&entity)
		if result.Error != nil {
			return result.Error
		}
	} else {
		if entity.ID == 0 {
			entity.ID = existingValue.ID
		}
		entity.Version = existingValue.Version + 1
		result := s.DbContext.Database.Save(entity)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
func (s *ProjectRepository) ListByStartWithPath(paths []string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Path') IN (?)", paths).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *ProjectRepository) ListByPath(paths ...string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Path') IN (?)", paths).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ProjectRepository) Delete(entity *entities.Project) error {
	result := s.DbContext.Database.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ProjectRepository) DeleteByName(name string) error {
	entity, err := s.GetByName(name)
	if err != nil {
		return err
	}

	result := s.DbContext.Database.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *ProjectRepository) GetIndividualByWorkspaceName(workspaceName string) (*entities.Project, error) {
	entity := new(entities.Project)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Group') = '' and json_extract(document, '$.Specifications.Workspace') = ?", workspaceName).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *ProjectRepository) GetByNameGroupAndWorkspaceName(name string, groupName string, workspaceName string) (*entities.Project, error) {
	entity := new(entities.Project)
	result := s.DbContext.Database.Where("json_extract(document, '$.Header.Name') = ? and json_extract(document, '$.Specifications.Group') = ? and json_extract(document, '$.Specifications.Workspace') = ?", name, groupName, workspaceName).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *ProjectRepository) ListByFilter(set, workspace string, layers []string, tags []string, labels []map[string]string) (*([]entities.Project), error) {
	var projects []entities.Project
	var args []interface{}
	var joins []string
	var where []string

	sql := strings.Builder{}
	sql.WriteString("SELECT DISTINCT projects.* FROM projects")

	// Join'leri ihtiyaca göre ekle
	if len(layers) > 0 {
		joins = append(joins, "LEFT JOIN json_each(projects.document, '$.Specifications.Layers') AS l")
	}
	if len(tags) > 0 {
		joins = append(joins, "LEFT JOIN json_each(projects.document, '$.Header.Metadata.Tags') AS t")
	}
	if len(labels) > 0 {
		joins = append(joins, "LEFT JOIN json_each(projects.document, '$.Specifications.Labels') AS lbl")
	}
	if len(joins) > 0 {
		sql.WriteString("\n" + strings.Join(joins, "\n"))
	}

	// workspace
	where = append(where, "json_extract(projects.document, '$.Specifications.Workspace') = ?")
	args = append(args, workspace)

	// set
	where = append(where, "json_extract(projects.document, '$.Specifications.Set') = ?")
	args = append(args, set)

	// layers
	if len(layers) > 0 {
		layerPlaceholders := make([]string, len(layers))
		for i, layer := range layers {
			layerPlaceholders[i] = "?"
			args = append(args, layer)
		}
		where = append(where, fmt.Sprintf("json_extract(l.value, '$.Name') IN (%s)", strings.Join(layerPlaceholders, ",")))
	}

	// tags
	if len(tags) > 0 {
		tagPlaceholders := make([]string, len(tags))
		for i, tag := range tags {
			tagPlaceholders[i] = "?"
			args = append(args, tag)
		}
		where = append(where, fmt.Sprintf("t.value IN (%s)", strings.Join(tagPlaceholders, ",")))
	}

	// labels
	if len(labels) > 0 {
		var labelConds []string
		for _, lbl := range labels {
			var condParts []string
			if key, ok := lbl["Key"]; ok && key != "" {
				condParts = append(condParts, "json_extract(lbl.value, '$.Key') = ?")
				args = append(args, key)
			}
			if val, ok := lbl["Value"]; ok && val != "" {
				condParts = append(condParts, "json_extract(lbl.value, '$.Value') = ?")
				args = append(args, val)
			}
			if len(condParts) > 0 {
				labelConds = append(labelConds, "("+strings.Join(condParts, " AND ")+")")
			}
		}
		if len(labelConds) > 0 {
			where = append(where, "("+strings.Join(labelConds, " OR ")+")")
		}
	}

	// WHERE varsa ekle
	if len(where) > 0 {
		sql.WriteString("\nWHERE " + strings.Join(where, " AND "))
	}

	// Sorguyu çalıştır
	result := s.DbContext.Database.Raw(sql.String(), args...).Scan(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return &projects, nil
}

// func (s *ProjectRepository) ListBySet(set string) (*([]entities.Project), error) {
// 	var entities = make(([]entities.Project), 0)
// 	result := s.DbContext.Database.Preload("Group").Preload("Layers").Where("set = ?", set).Find(&entities)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &entities, nil
// }

func (s *ProjectRepository) ListBySet(set string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Set') = ?", set).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *ProjectRepository) ListByKind(kind string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Header.Kind') = ?", kind).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *ProjectRepository) ListBySetAndLayers(set string, layers ...string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	rawSQL := `
	SELECT DISTINCT projects.*
	FROM projects
	JOIN json_each(projects.document, '$.Specifications.Configuration.Layers') AS json_each
	WHERE json_extract(projects.document, '$.Specifications.Set') = ? and json_extract(json_each.value, '$.Name') IN (?)
`
	result := s.DbContext.Database.Raw(rawSQL, set, layers).Scan(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ProjectRepository) ListByWorkspace(workspaceName string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ?", workspaceName).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ProjectRepository) ListByWorkspaceNameAndGroup(workspaceName string, groupName string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ? and json_extract(document, '$.Specifications.Group') = ?", workspaceName, groupName).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ProjectRepository) ListByGroup(groupId int) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Group.ID') = ?", groupId).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ProjectRepository) ListByGroupName(group string) (*([]entities.Project), error) {
	var entities = make(([]entities.Project), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Group.Name') = ?", group).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

// func (s *ProjectRepository) ListByStartWithPath(paths []string) (*([]entities.Project), error) {
// 	var entities = make(([]entities.Project), 0)
// 	result := s.DbContext.Database.Preload("Group").Preload("Layers").Where("path IN ?", paths).Find(&entities)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &entities, nil
// }

func (s *ProjectRepository) DeleteByWorkspace(workspaceName string) error {
	entities, err := s.ListByWorkspace(workspaceName)
	if err != nil {
		return err
	}

	for _, entity := range *entities {
		s.Delete(&entity)
	}
	return nil
}

func (s *ProjectRepository) CountByWorkspaceID(workspaceID int) (int64, error) {
	var count int64
	result := s.DbContext.Database.Table("projects").Where("json_extract(document, '$.Specifications.Workspace.ID') = ?", workspaceID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (s *ProjectRepository) CountByWorkspaceIDAndGroup(workspaceName string, groupName string) (int64, error) {
	var count int64
	result := s.DbContext.Database.Table("projects").Where("json_extract(document, '$.Specifications.Workspace') = ? and json_extract(document, '$.Specifications.Group') = ?", workspaceName, groupName).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
