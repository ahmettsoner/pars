package repositories

import (
	"errors"

	"parsdevkit.net/core/utils"

	"parsdevkit.net/persistence/contexts"

	"parsdevkit.net/persistence/entities"

	"gorm.io/gorm"
)

type TemplateRepository struct {
	DbContext *contexts.DbContext
}

func NewTemplateRepository(environment string) *TemplateRepository {
	return &TemplateRepository{DbContext: contexts.New(environment)}
}

func (s *TemplateRepository) Get(id int) (*entities.Template, error) {
	entity := new(entities.Template)
	result := s.DbContext.Database.First(entity, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return entity, nil
}

func (s *TemplateRepository) GetByName(name string) (*entities.Template, error) {
	entity := new(entities.Template)
	result := s.DbContext.Database.Where("json_extract(document, '$.Name')= ?", name).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}
func (s *TemplateRepository) GetByNameAndWorkspace(name, workspace string) (*entities.Template, error) {
	entity := new(entities.Template)
	result := s.DbContext.Database.Where("json_extract(document, '$.Name') = ? and json_extract(document, '$.Specifications.Workspace') = ?", name, workspace).First(entity)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return entity, nil
}

func (s *TemplateRepository) ListBySetAndLayers(set string, layers ...string) (*([]entities.Template), error) {
	var entities = make(([]entities.Template), 0)
	rawSQL := `
	SELECT templates.*
	FROM templates
	JOIN json_each(templates.document, '$.Specifications.Layers') AS json_each
	WHERE json_extract(templates.document, '$.Specifications.Set') = ? and json_extract(json_each.value, '$.Name') IN (?)
`
	result := s.DbContext.Database.Raw(rawSQL, set, layers).Scan(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TemplateRepository) ListByWorkspaceSetAndLayers(workspace, set string, layers ...string) (*([]entities.Template), error) {
	var entities = make(([]entities.Template), 0)
	rawSQL := `
	SELECT templates.*
	FROM templates
	JOIN json_each(templates.document, '$.Specifications.Layers') AS json_each
	WHERE json_extract(document, '$.Specifications.Workspace') = ? and json_extract(resources.document, '$.Specifications.Set') = ? and json_extract(json_each.value, '$.Name') IN (?)
`
	result := s.DbContext.Database.Raw(rawSQL, set, layers).Scan(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TemplateRepository) List() (*([]entities.Template), error) {
	var entities = make(([]entities.Template), 0)
	result := s.DbContext.Database.Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TemplateRepository) ListByWorkspace(workspace string) (*([]entities.Template), error) {
	var entities = make(([]entities.Template), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ?", workspace).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TemplateRepository) ListByKind(kind string) (*([]entities.Template), error) {
	var entities = make(([]entities.Template), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Kind') = ?", kind).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TemplateRepository) ListByWorkspaceAndKind(workspace, kind string) (*([]entities.Template), error) {
	var entities = make(([]entities.Template), 0)
	result := s.DbContext.Database.Where("json_extract(document, '$.Specifications.Workspace') = ? and json_extract(document, '$.Kind') = ?", workspace, kind).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *TemplateRepository) Save(entity *entities.Template) error {

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

func (s *TemplateRepository) Delete(entity *entities.Template) error {
	result := s.DbContext.Database.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
