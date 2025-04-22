package repositories

import (
	"errors"
	"time"

	"parsdevkit.net/persistence/contexts"
	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type GenerationHistoryRepository struct {
	DbContext *contexts.DbContext
}

func NewGenerationHistoryRepository(environment string) *GenerationHistoryRepository {
	return &GenerationHistoryRepository{DbContext: contexts.New(environment)}
}

func (s *GenerationHistoryRepository) Get(id int) (*entities.GenerationHistory, error) {
	entity := new(entities.GenerationHistory)
	result := s.DbContext.Database.First(&entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *GenerationHistoryRepository) GetLast(set, resource, template, section, layer string) (*entities.GenerationHistory, error) {
	entity := new(entities.GenerationHistory)
	result := s.DbContext.Database.Where("`set` = ? and resource = ? and template = ? and section = ? and layer = ?", set, resource, template, section, layer).Order("Timestamp desc").First(&entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *GenerationHistoryRepository) Create(entity *entities.GenerationHistory) error {

	entity.Timestamp = time.Now()
	result := s.DbContext.Database.Create(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *GenerationHistoryRepository) Update(entity *entities.GenerationHistory) error {
	result := s.DbContext.Database.Save(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *GenerationHistoryRepository) Delete(entity *entities.GenerationHistory) error {
	result := s.DbContext.Database.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *GenerationHistoryRepository) DeleteBySetAndResource(set, resource string) error {
	result := s.DbContext.Database.Where("`set` = ? and resource = ?", set, resource).Delete(&entities.GenerationHistory{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *GenerationHistoryRepository) DeleteBySetAndTemplate(set, template string) error {
	result := s.DbContext.Database.Where("`set` = ? and template = ?", set, template).Delete(&entities.GenerationHistory{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *GenerationHistoryRepository) DeleteBySetAndLayer(set, layer string) error {
	result := s.DbContext.Database.Where("`set` = ? and layer = ?", set, layer).Delete(&entities.GenerationHistory{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
