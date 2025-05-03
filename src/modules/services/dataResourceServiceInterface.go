package services

import (
	dataresource "parsdevkit.net/structs/resource/data-resource"
)

type DataResourceServiceInterface interface {
	GetByName(name string) (*dataresource.ResourceBaseStruct, error)
	Save(model dataresource.ResourceBaseStruct) (*dataresource.ResourceBaseStruct, error)
	List() (*([]dataresource.ResourceBaseStruct), error)
	ListByWorkspace(workspace string) (*([]dataresource.ResourceBaseStruct), error)
	ListBySet(set string) (*([]dataresource.ResourceBaseStruct), error)
	ListByWorkspaceAndSet(workspace, set string) (*([]dataresource.ResourceBaseStruct), error)
	ListBySetAndLayers(set string, layers ...string) (*([]dataresource.ResourceBaseStruct), error)
	ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]dataresource.ResourceBaseStruct), error)
	Remove(name, workspace string, force, permanent bool) (*dataresource.ResourceBaseStruct, error)
	IsExists(name, workspace string) bool
	GetHash(name string) string
}
