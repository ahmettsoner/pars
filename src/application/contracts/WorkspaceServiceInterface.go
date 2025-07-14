package contracts

import "parsdevkit.net/application/schemas"

type WorkspaceServiceInterface[T schemas.SchemaInterface] interface {
	GetByName(name string) (*T, error)
	Save(model T) (*T, error)
	SaveWorkspace(model T) (*T, error)
	UndoSaveWorkspace(model T) (*T, error)
	DeleteWorkspace(model T) (*T, error)
	List() (*([]T), error)
	Remove(name string, force, permanent bool) (*T, error)
	IsExists(name string) (bool, error)
	GetHash(name string) (string, error)
	GetSelectedWorkspace() (*T, error)
	GetActiveWorkspace() (*T, error)
	ChangeCurrentWorkspace(name string) (*T, error)
	ListByNameStartWith(name string) (*([]T), error)
	CreateWorkspaceFolder(model T) (string, error)
}
