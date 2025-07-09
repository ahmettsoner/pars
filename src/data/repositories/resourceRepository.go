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

type ResourceRepository struct {
	DbContext *contexts.DbContext
}

func NewResourceRepository(dbCtx *contexts.DbContext) *ResourceRepository {
	return &ResourceRepository{DbContext: dbCtx}
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
	result := s.DbContext.Database.Where("json_extract(document, '$.Header.Name')= ?", name).First(entity)
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
	result := s.DbContext.Database.Where("json_extract(document, '$.Header.Name') = ? and json_extract(document, '$.Specifications.Workspace') = ?", name, workspace).First(entity)
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

func (s *ResourceRepository) ListByFilter(set, workspace string, layers []string, tags []string, labels []map[string]string) (*([]entities.Resource), error) {
	var resources []entities.Resource
	var args []interface{}
	var joins []string
	var where []string

	sql := strings.Builder{}
	sql.WriteString("SELECT DISTINCT resources.* FROM resources")

	// Join'leri ihtiyaca göre ekle
	if len(layers) > 0 {
		joins = append(joins, "LEFT JOIN json_each(resources.document, '$.Specifications.Layers') AS l")
	}
	if len(tags) > 0 {
		joins = append(joins, "LEFT JOIN json_each(resources.document, '$.Header.Metadata.Tags') AS t")
	}
	if len(labels) > 0 {
		joins = append(joins, "LEFT JOIN json_each(resources.document, '$.Specifications.Labels') AS lbl")
	}
	if len(joins) > 0 {
		sql.WriteString("\n" + strings.Join(joins, "\n"))
	}

	// workspace
	where = append(where, "json_extract(resources.document, '$.Specifications.Workspace') = ?")
	args = append(args, workspace)

	// set
	where = append(where, "json_extract(resources.document, '$.Specifications.Set') = ?")
	args = append(args, set)

	if len(layers) > 0 {
		// JOIN gerekiyor
		joins = append(joins, `LEFT JOIN json_each(resources.document, '$.Specifications.Layers') AS l`)

		// IN (...) ile filtreleme
		layerPlaceholders := make([]string, len(layers))
		for i, layer := range layers {
			layerPlaceholders[i] = "?"
			args = append(args, layer)
		}
		where = append(where, fmt.Sprintf("json_extract(l.value, '$.Name') IN (%s)", strings.Join(layerPlaceholders, ",")))
	} else {
		// Layer verilmemişse: Layer alanı tanımlı değil veya boş array
		where = append(where, `
			(
				json_type(json_extract(document, '$.Specifications.Layers')) IS NULL
				OR json_array_length(json_extract(document, '$.Specifications.Layers')) = 0
			)
		`)
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
	result := s.DbContext.Database.Raw(sql.String(), args...).Scan(&resources)
	if result.Error != nil {
		return nil, result.Error
	}
	return &resources, nil
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
	result := s.DbContext.Database.Where("json_extract(document, '$.Header.Kind') = ?", kind).Find(&entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entities, nil
}

func (s *ResourceRepository) ListByWorkspaceAndKind(workspace, kind string) (*([]entities.Resource), error) {
	var entities = make(([]entities.Resource), 0)
	query := s.DbContext.Database.Model(&entities).Where("json_extract(document, '$.Specifications.Workspace') = ? and json_extract(document, '$.Header.Kind') = ?", workspace, kind)
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

func (s *ResourceRepository) Delete(entity *entities.Resource) error {
	result := s.DbContext.Database.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ResourceRepository) DeleteByName(name string) error {
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
