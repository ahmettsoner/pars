package services

import (
	objectresource "parsdevkit.net/structs/resource/object-resource"
)

type ObjectResourceServiceInterface interface {
	GetByName(name string) (*objectresource.ResourceBaseStruct, error)
	Save(model objectresource.ResourceBaseStruct) (*objectresource.ResourceBaseStruct, error)
	List() (*([]objectresource.ResourceBaseStruct), error)
	ListByWorkspace(workspace string) (*([]objectresource.ResourceBaseStruct), error)
	ListBySet(set string) (*([]objectresource.ResourceBaseStruct), error)
	ListByWorkspaceAndSet(workspace, set string) (*([]objectresource.ResourceBaseStruct), error)
	ListBySetAndLayers(set string, layers ...string) (*([]objectresource.ResourceBaseStruct), error)
	ListByWorkspaceAndSetAndLayers(workspace, set string, layers ...string) (*([]objectresource.ResourceBaseStruct), error)
	Remove(name, workspace string, force, permanent bool) (*objectresource.ResourceBaseStruct, error)
	IsExists(name, workspace string) bool
	GetHash(name string) string
}
