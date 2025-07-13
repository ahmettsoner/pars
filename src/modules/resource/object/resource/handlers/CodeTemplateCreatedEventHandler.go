package handlers

import (
	"fmt"

	code_template_payload_events "parsdevkit.net/modules/template/code_template_payload/events"
)

type CodeTemplateCreatedEventHandler struct{}

func (s *CodeTemplateCreatedEventHandler) Handle(evt code_template_payload_events.TemplateCreated) {
	fmt.Println("code template created event handled...")
}
