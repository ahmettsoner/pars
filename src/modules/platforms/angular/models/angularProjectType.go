package models

type AngularProjectType string

var AngularProjectTypes = struct {
	SPA     AngularProjectType
	Library AngularProjectType
}{
	SPA:     "spa",
	Library: "library",
}

func (c AngularProjectType) String() string {
	switch c {
	case "spa":
		return "Single Page Application"
	case "library":
		return "Library"
	default:
		return "Unknown"
	}
}
