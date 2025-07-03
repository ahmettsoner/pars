package project

import (
	"parsdevkit.net/pkg/utilities/file"

	"parsdevkit.net/pkg/errors"
)

func ParseProjectFullName(fullname string) (string, string, error) {
	parts := file.PathToArray(fullname)

	switch len(parts) {
	case 1:
		return "", parts[0], nil
	case 2:
		return parts[0], parts[1], nil
	default:
		return "", "", &errors.InvalidProjectFullnameFormatError{Value: fullname}
	}
}
