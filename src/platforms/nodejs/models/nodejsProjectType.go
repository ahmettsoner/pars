package models

type NodeJSProjectType string

var NodeJSProjectTypes = struct {
	Library NodeJSProjectType
}{
	Library: "library",
}

func (c NodeJSProjectType) String() string {
	switch c {
	case "library":
		return "Library"
	default:
		return "Unknown"
	}
}
