package services

import (
	"parsdevkit.net/structs/group"
)

type GroupServiceInterface interface {
	GetByName(name string) (*group.GroupBaseStruct, error)
	Save(model group.GroupBaseStruct) (*group.GroupBaseStruct, error)
	List() (*([]group.GroupBaseStruct), error)
	Remove(name string, permanent bool) (*group.GroupBaseStruct, error)
	GetHash(name string) string
}
