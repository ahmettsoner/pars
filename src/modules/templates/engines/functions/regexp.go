package functions

import (
	"fmt"

	regexp "parsdevkit.net/core/utils/regexp"
)

type RegexpFuncs struct{}

func (r RegexpFuncs) Find(expression, input string) string {
	result, err := regexp.Find(expression, input)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (r RegexpFuncs) FindAll(expression string, n int, input string) []string {
	result, err := regexp.FindAll(expression, n, input)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (r RegexpFuncs) Match(expression, input string) bool {
	result, err := regexp.Match(expression, input)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (r RegexpFuncs) QuoteMeta(input string) string {
	return regexp.QuoteMeta(input)
}

func (r RegexpFuncs) Replace(expression, replacement, input string) string {
	result, err := regexp.Replace(expression, replacement, input)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (r RegexpFuncs) ReplaceLiteral(expression, replacement, input string) string {
	result, err := regexp.ReplaceLiteral(expression, replacement, input)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}

func (r RegexpFuncs) Split(expression string, n int, input string) []string {
	result, err := regexp.Split(expression, n, input)
	if err != nil {
		fmt.Print("An error occured...")
	}
	return result
}
