package models

import (
	"fmt"
	"strings"
)

type ProjectType string

var ProjectTypes = struct {
	Project ProjectType
	WebApi  ProjectType
	WebApp  ProjectType
	SPA     ProjectType
	Library ProjectType
	Console ProjectType
	Desktop ProjectType
	Mobile  ProjectType
}{
	Project: "project",
	WebApi:  "webapi",
	WebApp:  "webapp",
	SPA:     "spa",
	Console: "console",
	Library: "library",
	Desktop: "desktop",
	Mobile:  "mobile",
}

func (c ProjectType) String() string {
	switch c {
	case "project":
		return "Project"
	case "webapi":
		return "Web Api"
	case "webapp":
		return "Web Application"
	case "spa":
		return "Single Page Application"
	case "console":
		return "Console Application"
	case "library":
		return "Library"
	case "desktop":
		return "Desktop Application"
	case "mobile":
		return "Mobile Application"
	default:
		return "Unknown"
	}
}

func ProjectTypeEnumFromString(enum string) (ProjectType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower("Project"):
		return ProjectTypes.Project, nil
	case strings.ToLower("WebApi"):
		return ProjectTypes.WebApi, nil
	case strings.ToLower("WebApp"):
		return ProjectTypes.WebApp, nil
	case strings.ToLower("Console"):
		return ProjectTypes.Console, nil
	case strings.ToLower("Library"):
		return ProjectTypes.Library, nil
	case strings.ToLower("Desktop"):
		return ProjectTypes.Desktop, nil
	case strings.ToLower("Mobile"):
		return ProjectTypes.Mobile, nil
	case strings.ToLower("SPA"):
		return ProjectTypes.SPA, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *ProjectType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := ProjectTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
