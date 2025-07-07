package handlers

import (
	"fmt"

	"parsdevkit.net/modules/project/application_project_payload/commands"
)

type CreateApplicationProjectHandler struct{}

func (h *CreateApplicationProjectHandler) Handle(cmd commands.CreateApplicationProject) error {
	fmt.Println("[CommandHandler] User created:")
	return nil
}
