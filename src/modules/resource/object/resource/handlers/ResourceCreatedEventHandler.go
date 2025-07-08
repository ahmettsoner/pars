package handlers

import (
	"fmt"

	"parsdevkit.net/application"
	"parsdevkit.net/application/ioc"
	"parsdevkit.net/engines"
	"parsdevkit.net/modules/resource/object_resource_contract"
	"parsdevkit.net/modules/resource/object_resource_payload"
	"parsdevkit.net/modules/resource/object_resource_payload/events"
)

type ResourceCreatedEventHandler struct{}

func (s *ResourceCreatedEventHandler) Handle(event events.ResourceCreated) {
	fmt.Println("[EventHandler] Resource '%s' generated:", event.Data.GetHeader().Name)

	model, ok := event.Data.(object_resource_payload.ResourceBaseStruct)
	if !ok {
		panic(fmt.Errorf("invalid item type: expected object_resource_payload.ResourceBaseStruct, got %T", model))
	}
	if _, err := s.generate(model); err != nil {
		panic(err)
	}
}

func (s *ResourceCreatedEventHandler) generate(model object_resource_payload.ResourceBaseStruct) (*object_resource_payload.ResourceBaseStruct, error) {

	resourceService := ioc.Get[object_resource_contract.ResourceInterface]()

	result, err := resourceService.GetByName(model.Header.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	// TODO: Birden fazla template işlenebilmeli
	templateEngine := engines.NewCodeTemplateOperations(application.GetEnvironment())
	err = templateEngine.GenerateByResource(model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
