package repositories

import (
	"errors"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/contexts"

	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type ResourceRepository struct {
	DbContext *contexts.DbContext
}

func NewResourceRepository(environment string) *ResourceRepository {
	return &ResourceRepository{DbContext: contexts.New(environment)}
}

func (s *ResourceRepository) Get(id int) (*entities.Resource, error) {
	entity := new(entities.Resource)
	result := s.DbContext.Database.First(entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return entity, nil
}

func (s *ResourceRepository) GetByName(name string) (*entities.Resource, error) {
	entity := new(entities.Resource)
	result := s.DbContext.Database.Where("json_extract(document, '$.Name')= ?", name).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}
func (s *ResourceRepository) GetByNameAndWorkspace(name, workspace string) (*entities.Resource, error) {
	entity := new(entities.Resource)
	result := s.DbContext.Database.Where("json_extract(document, '$.Name') = ? and json_extract(document, '$.Specifications.Workspace') = ?", name, workspace).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *ResourceRepository) ListBySet(set string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Set') = ?", set).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *ResourceRepository) ListByWorkspaceAndSet(workspace, set string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ? and json_extract(document, '$.Specifications.Set') = ?", workspace, set).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *ResourceRepository) ListBySetAndLayers(set string, layers ...string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	rawSQL := `
	SELECT resources.*
	FROM resources
	JOIN json_each(resources.document, '$.Specifications.Layers') AS json_each
	WHERE json_extract(resources.document, '$.Specifications.Set') = ? and json_extract(json_each.value, '$.Name') IN (?)
`
	result := s.DbContext.Database.Raw(rawSQL, set, layers).Scan(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ResourceRepository) ListByWorkspaceSetAndLayers(workspace, set string, layers ...string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	rawSQL := `
	SELECT resources.*
	FROM resources
	JOIN json_each(resources.document, '$.Specifications.Layers') AS json_each
	WHERE json_extract(document, '$.Specifications.Workspace') = ? and json_extract(resources.document, '$.Specifications.Set') = ? and json_extract(json_each.value, '$.Name') IN (?)
`
	result := s.DbContext.Database.Raw(rawSQL, workspace, set, layers).Scan(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ResourceRepository) List() (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	result := s.DbContext.Database.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ResourceRepository) ListWorkspace(workspace string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ?", workspace).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ResourceRepository) ListByKind(kind string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Kind') = ?", kind).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ResourceRepository) ListByWorkspaceAndKind(workspace, kind string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	query := s.DbContext.Database.Model(&entities).Where("json_extract(document, '$.Specifications.Workspace') = ? and json_extract(document, '$.Kind') = ?", workspace, kind)
	result := query.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ResourceRepository) Save(entity *entities.Resource) error {

	existingValue, err := s.GetByName(entity.Name)
	if err != nil {
		return err
	}

	documentHash, err := utils.CalculateHash(entity.Document)
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

func (s *ResourceRepository) Delete(entity *entities.Resource) error {
	result := s.DbContext.Database.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
