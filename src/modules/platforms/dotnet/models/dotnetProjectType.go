package models

type DotnetProjectType string

var DotnetProjectTypes = struct {
	WebApi  DotnetProjectType
	WebApp  DotnetProjectType
	Library DotnetProjectType
	Console DotnetProjectType
}{
	WebApi:  "webapi",
	WebApp:  "webapp",
	Library: "library",
	Console: "console",
}

func (c DotnetProjectType) String() string {
	switch c {
	case "webapi":
		return "Web Api"
	case "webapp":
		return "Web App"
	case "library":
		return "Library"
	case "console":
		return "Console"
	default:
		return "Unknown"
	}
}
