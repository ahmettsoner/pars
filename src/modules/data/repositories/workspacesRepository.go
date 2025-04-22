package repositories

import (
	"errors"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/contexts"

	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type WorkspaceRepository struct {
	DbContext *contexts.DbContext
}

func NewWorkspaceRepository(environment string) *WorkspaceRepository {
	return &WorkspaceRepository{DbContext: contexts.New(environment)}
}

func (s *WorkspaceRepository) Get(id int) (*entities.Workspace, error) {
	entity := new(entities.Workspace)
	result := s.DbContext.Database.First(entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return entity, nil
}
func (s *WorkspaceRepository) GetByName(name string) (*entities.Workspace, error) {
	entity := new(entities.Workspace)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Name') = ?", name).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}
func (s *WorkspaceRepository) ListByNameStartWith(name string) (*([]entities.Workspace), error) {
	var entities = make(([]entities.Workspace), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Name') LIKE ?", name+"%").Find(&entities)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &entities, nil
}
func (s *WorkspaceRepository) List() (*([]entities.Workspace), error) {
	var entities = make(([]entities.Workspace), 0)
	result := s.DbContext.Database.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *WorkspaceRepository) Save(entity *entities.Workspace) error {

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
		entity.Version = existingValue.Version + 1
		if entity.ID == 0 {
			entity.ID = existingValue.ID
		}
		result := s.DbContext.Database.Save(entity)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s *WorkspaceRepository) ListByStartWithPath(paths []string) (*([]entities.Workspace), error) {
	var entities = make(([]entities.Workspace), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Path') IN (?)", paths).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *WorkspaceRepository) ListByPath(paths ...string) (*([]entities.Workspace), error) {
	var entities = make(([]entities.Workspace), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Path') IN (?)", paths).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *WorkspaceRepository) Delete(entity *entities.Workspace) error {
	result := s.DbContext.Database.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
