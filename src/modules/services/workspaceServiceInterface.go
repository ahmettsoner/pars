package services

import (
	"parsdevkit.net/structs/workspace"
)

type WorkspaceServiceInterface interface {
	GetByName(name string) (*workspace.WorkspaceBaseStruct, error)
	Save(model workspace.WorkspaceBaseStruct) (*workspace.WorkspaceBaseStruct, error)
	List() (*([]workspace.WorkspaceBaseStruct), error)
	Remove(name string, force, permanent bool) (*workspace.WorkspaceBaseStruct, error)
	IsExists(name string) (bool, error)
	GetHash(name string) (string, error)
}
