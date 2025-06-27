package repositories

import (
	"errors"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type GroupRepository struct {
	DbContext *contexts.DbContext
}

func NewGroupRepository(environment string) *GroupRepository {
	return &GroupRepository{DbContext: contexts.New(environment)}
}

func (s *GroupRepository) Get(id int) (*entities.Group, error) {
	entity := new(entities.Group)
	result := s.DbContext.Database.First(entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return entity, nil
}
func (s *GroupRepository) GetByName(name string) (*entities.Group, error) {
	entity := new(entities.Group)
	result := s.DbContext.Database.Where("json_extract(document, '$.Name')= ?", name).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}
func (s *GroupRepository) List() (*([]entities.Group), error) {
	var entities = make(([]entities.Group), 0)
	result := s.DbContext.Database.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *GroupRepository) Save(entity *entities.Group) error {

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

func (s *GroupRepository) ListByStartWithPath(paths []string) (*([]entities.Group), error) {
	var entities = make(([]entities.Group), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Path') IN (?)", paths).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}
func (s *GroupRepository) ListByPath(paths ...string) (*([]entities.Group), error) {
	var entities = make(([]entities.Group), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Path') IN (?)", paths).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *GroupRepository) Delete(entity *entities.Group) error {
	result := s.DbContext.Database.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
