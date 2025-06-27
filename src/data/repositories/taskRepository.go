package repositories

import (
	"errors"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/contexts"

	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DbContext *contexts.DbContext
}

func NewTaskRepository(environment string) *TaskRepository {
	return &TaskRepository{DbContext: contexts.New(environment)}
}

func (s *TaskRepository) Get(id int) (*entities.Task, error) {
	entity := new(entities.Task)
	result := s.DbContext.Database.First(entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return entity, nil
}

func (s *TaskRepository) GetByName(name string) (*entities.Task, error) {
	entity := new(entities.Task)
	result := s.DbContext.Database.Where("json_extract(document, '$.Name')= ?", name).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}
func (s *TaskRepository) GetByNameAndWorkspace(name, workspace string) (*entities.Task, error) {
	entity := new(entities.Task)
	result := s.DbContext.Database.Where("json_extract(document, '$.Name') = ? and json_extract(document, '$.Specifications.Workspace') = ?", name, workspace).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *TaskRepository) ListBySetAndLayers(set string, layers ...string) (*([]entities.Task), error) {
	var entities = make(([]entities.Task), 0)
	rawSQL := `
	SELECT tasks.*
	FROM tasks
	JOIN json_each(tasks.document, '$.Specifications.Layers') AS json_each
	WHERE json_extract(tasks.document, '$.Specifications.Set') = ? and json_extract(json_each.value, '$.Name') IN (?)
`
	result := s.DbContext.Database.Raw(rawSQL, set, layers).Scan(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TaskRepository) ListByWorkspaceSetAndLayers(workspace, set string, layers ...string) (*([]entities.Task), error) {
	var entities = make(([]entities.Task), 0)
	rawSQL := `
	SELECT tasks.*
	FROM tasks
	JOIN json_each(tasks.document, '$.Specifications.Layers') AS json_each
	WHERE json_extract(document, '$.Specifications.Workspace') = ? and json_extract(tasks.document, '$.Specifications.Set') = ? and json_extract(json_each.value, '$.Name') IN (?)
`
	result := s.DbContext.Database.Raw(rawSQL, set, layers).Scan(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TaskRepository) List() (*([]entities.Task), error) {
	var entities = make(([]entities.Task), 0)
	result := s.DbContext.Database.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TaskRepository) ListByWorkspace(workspace string) (*([]entities.Task), error) {
	var entities = make(([]entities.Task), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ?", workspace).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TaskRepository) ListByKind(kind string) (*([]entities.Task), error) {
	var entities = make(([]entities.Task), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Kind') = ?", kind).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TaskRepository) ListByWorkspaceAndKind(workspace, kind string) (*([]entities.Task), error) {
	var entities = make(([]entities.Task), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ? and json_extract(document, '$.Kind') = ?", workspace, kind).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TaskRepository) Save(entity *entities.Task) error {

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

func (s *TaskRepository) Delete(entity *entities.Task) error {
	result := s.DbContext.Database.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
