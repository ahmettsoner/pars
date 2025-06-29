package contracts

import (
	"parsdevkit.net/core/schemas"
)

type TaskServiceInterface[T schemas.Schema] interface {
	GetByName(name string) (*T, error)
	Save(mommonl T) (*T, error)
	List() (*([]T), error)
	ListBySetAndLayers(set string, layers ...string) (*([]T), error)
	Remove(name, workspace string, permanent bool) (*T, error)
	IsExists(name, workspace string) (bool, error)
	GetHash(name string) (string, error)
}
