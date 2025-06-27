package repositories

import (
	"errors"

	"parsdevkit.net/persistence/contexts"

	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type SettingsRepository struct {
	DbContext *contexts.DbContext
}

func NewSettingsRepository(environment string) *SettingsRepository {
	return &SettingsRepository{DbContext: contexts.New(environment)}
}

func (s *SettingsRepository) Get(key string) (*entities.Settings, error) {
	entity := new(entities.Settings)
	result := s.DbContext.Database.Where("key = ?", key).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *SettingsRepository) GetValue(key string) (string, error) {
	entity := new(entities.Settings)
	result := s.DbContext.Database.Where("key = ?", key).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", result.Error
	}

	return entity.Value, nil
}

func (s *SettingsRepository) SetValue(key string, value string) error {
	existingSettings, err := s.Get(key)
	if err != nil {
		return err
	}

	if existingSettings == nil {
		entity := entities.Settings{Key: key, Value: value}
		result := s.DbContext.Database.Create(entity)
		if result.Error != nil {
			return result.Error
		}
	} else {
		existingSettings.Value = value
		s.Update(existingSettings)
	}

	return nil
}

func (s *SettingsRepository) List() (*([]entities.Settings), error) {
	var entities = make(([]entities.Settings), 0)
	result := s.DbContext.Database.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *SettingsRepository) Create(entity *entities.Settings) error {

	result := s.DbContext.Database.Create(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SettingsRepository) Update(entity *entities.Settings) error {
	result := s.DbContext.Database.Save(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SettingsRepository) Delete(entity *entities.Settings) error {
	result := s.DbContext.Database.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
