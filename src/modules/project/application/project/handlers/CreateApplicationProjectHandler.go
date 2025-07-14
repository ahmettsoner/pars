package handlers

import (
	application_project_payload_commands "parsdevkit.net/modules/project/application_project_payload/commands"
)

type CreateApplicationProjectHandler struct{}

func (h *CreateApplicationProjectHandler) Handle(cmd application_project_payload_commands.CreateApplicationProject) error {
	return nil
}
