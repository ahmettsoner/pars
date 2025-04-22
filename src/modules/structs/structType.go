package structs

import (
	"fmt"
	"strings"
)

type StructType string

var StructTypes = struct {
	Workspace   StructType
	Project     StructType
	Group       StructType
	Resource    StructType
	Stack       StructType
	Environment StructType
	Layer       StructType
	Template    StructType
	Task        StructType
}{
	Workspace:   "Workspace",
	Project:     "Project",
	Group:       "Group",
	Resource:    "Resource",
	Stack:       "Stack",
	Environment: "Environment",
	Layer:       "Layer",
	Template:    "Template",
	Task:        "Task",
}

func (c StructType) String() string {
	switch c {
	case StructTypes.Workspace:
		return "Workspace"
	case StructTypes.Project:
		return "Project"
	case StructTypes.Group:
		return "Group"
	case StructTypes.Resource:
		return "Resource"
	case StructTypes.Stack:
		return "Stack"
	case StructTypes.Environment:
		return "Environment"
	case StructTypes.Layer:
		return "Layer"
	case StructTypes.Template:
		return "Template"
	case StructTypes.Task:
		return "Task"
	default:
		return "Unknown"
	}
}
func StructTypeEnumFromString(enum string) (StructType, error) {
	switch strings.ToLower(enum) {
	case strings.ToLower(StructTypes.Workspace.String()):
		return StructTypes.Workspace, nil
	case strings.ToLower(StructTypes.Project.String()):
		return StructTypes.Project, nil
	case strings.ToLower(StructTypes.Group.String()):
		return StructTypes.Group, nil
	case strings.ToLower(StructTypes.Resource.String()):
		return StructTypes.Resource, nil
	case strings.ToLower(StructTypes.Stack.String()):
		return StructTypes.Stack, nil
	case strings.ToLower(StructTypes.Environment.String()):
		return StructTypes.Environment, nil
	case strings.ToLower(StructTypes.Layer.String()):
		return StructTypes.Layer, nil
	case strings.ToLower(StructTypes.Template.String()):
		return StructTypes.Template, nil
	case strings.ToLower(StructTypes.Task.String()):
		return StructTypes.Task, nil
	default:
		return "Unknown", fmt.Errorf("unknown state: %s", enum)
	}
}

func (s *StructType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}

	enum, err := StructTypeEnumFromString(value)
	if err != nil {
		return err
	}

	*s = enum
	return nil
}
