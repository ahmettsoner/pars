package services

import commontask "parsdevkit.net/structs/task/common-task"

type CommonTaskServiceInterface interface {
	GetByName(name string) (*commontask.TaskBaseStruct, error)
	Save(mommonl commontask.TaskBaseStruct) (*commontask.TaskBaseStruct, error)
	List() (*([]commontask.TaskBaseStruct), error)
	ListByWorkspace(workspace string) (*([]commontask.TaskBaseStruct), error)
	ListBySetAndLayers(set string, layers ...string) (*([]commontask.TaskBaseStruct), error)
	ListByWorkspaceSetAndLayers(workspace, set string, layers ...string) (*([]commontask.TaskBaseStruct), error)
	Remove(name, workspace string, permanent bool) (*commontask.TaskBaseStruct, error)
	IsExists(name, workspace string) bool
	GetHash(name string) string
}
