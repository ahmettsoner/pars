package handlers

import (
	"parsdevkit.net/modules/resource/object_resource_payload/events"
)

type ResourceCreatedEventHandler struct{}

func (s *ResourceCreatedEventHandler) Handle(evt events.ResourceCreated) {
}
