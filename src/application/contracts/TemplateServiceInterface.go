package contracts

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/schemas"
)

type TemplateServiceInterface[T schemas.SchemaInterface] interface {
	GetByName(name string) (*T, error)
	Save(model T) (*T, error)
	SaveTemplate(model T) (*T, error)
	UndoSaveTemplate(model T) (*T, error)
	DeleteTemplate(model T) (*T, error)
	ClearTemplateHistory(model T) error
	List() (*([]T), error)
	ListBySetAndLayers(set string, layers ...string) (*([]T), error)
	ListByFilter(set, workspace string, layers []string, tags []string, labels []label.Label) (*([]T), error)
	Remove(name, workspace string, permanent bool) (*T, error)
	IsExists(name, workspace string) (bool, error)
	GetHash(name string) (string, error)
	ListByWorkspace(workspace string) (*([]T), error)
	Context(model T) any
}
