package contracts

import (
	"parsdevkit.net/application/models/label"
	"parsdevkit.net/application/models/layer"
	"parsdevkit.net/application/models/section"
	"parsdevkit.net/application/schemas"
)

type ResourceServiceInterface[T schemas.SchemaInterface] interface {
	GetByName(name string) (*T, error)
	Save(model T) (*T, error)
	SaveResource(model T) (*T, error)
	UndoSaveResource(model T) (*T, error)
	DeleteResource(model T) (*T, error)
	ClearResourceHistory(model T) error
	List() (*([]T), error)
	ListByWorkspace(workspace string) (*([]T), error)
	ListBySet(set string) (*([]T), error)
	ListByWorkspaceAndSet(workspace, set string) (*([]T), error)
	ListBySetAndLayers(set string, layers ...string) (*([]T), error)
	ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]T), error)
	ListByFilter(set, workspace string, layers []string, tags []string, labels []label.Label) (*([]T), error)
	Remove(name, workspace string, force, permanent bool) (*T, error)
	IsExists(name, workspace string) (bool, error)
	GetHash(name string) (string, error)
	Context(workspace, project, resource, template schemas.SchemaInterface, layer layer.LayerIdentifier, section section.SectionIdentifier) any
}
