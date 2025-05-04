package regexp

import (
	"fmt"
	"regexp"
)

func Find(expression, input string) (string, error) {
	re, err := regexp.Compile(expression)
	if err != nil {
		return "", err
	}
	return re.FindString(input), nil
}

func FindAll(expression string, n int, input string) ([]string, error) {
	re, err := regexp.Compile(expression)
	if err != nil {
		return nil, err
	}
	return re.FindAllString(input, n), nil
}

func Match(expression, input string) (bool, error) {
	re, err := regexp.Compile(expression)
	if err != nil {
		return false, fmt.Errorf("error compiling expression: %w", err)
	}

	return re.MatchString(input), nil
}

func QuoteMeta(input string) string {
	return regexp.QuoteMeta(input)
}

func Replace(expression, replacement, input string) (string, error) {
	re, err := regexp.Compile(expression)
	if err != nil {
		return "", fmt.Errorf("error compiling expression: %w", err)
	}

	return re.ReplaceAllString(input, replacement), nil
}

func ReplaceLiteral(expression, replacement, input string) (string, error) {
	re, err := regexp.Compile(expression)
	if err != nil {
		return "", fmt.Errorf("error compiling expression: %w", err)
	}
	return re.ReplaceAllLiteralString(input, replacement), nil
}

func Split(expression string, n int, input string) ([]string, error) {
	re, err := regexp.Compile(expression)
	if err != nil {
		return nil, fmt.Errorf("error compiling expression: %w", err)
	}

	return re.Split(input, n), nil
}
