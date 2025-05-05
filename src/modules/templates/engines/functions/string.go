package functions

import (
	_string "parsdevkit.net/core/utils/string"
)

type StringFuncs struct{}

func (s StringFuncs) StartsWith(str, prefix string) bool {
	return _string.StartsWith(str, prefix)
}

func (s StringFuncs) EndsWith(str, suffix string) bool {
	return _string.EndsWith(str, suffix)
}

func (s StringFuncs) Contains(str, substr string) bool {
	return _string.Contains(str, substr)
}

func (s StringFuncs) ToUpperCase(str string) string {
	return _string.ToUpperCase(str)
}

func (s StringFuncs) ToLowerCase(str string) string {
	return _string.ToLowerCase(str)
}

func (s StringFuncs) ToPascalCase(str string) string {
	return _string.ToPascalCase(str)
}
func (s StringFuncs) ToCamelCase(str string) string {
	return _string.ToCamelCase(str)
}

func (s StringFuncs) ToSnakeCase(str string) string {
	return _string.ToSnakeCase(str)
}
func (s StringFuncs) ToKebabCase(str string) string {
	return _string.ToKebabCase(str)
}
func (s StringFuncs) ToUpperFlatCase(str string) string {
	return _string.ToUpperFlatCase(str)
}
func (s StringFuncs) ToLowerFlatCase(str string) string {
	return _string.ToLowerFlatCase(str)
}
func (s StringFuncs) ToCobolCase(str string) string {
	return _string.ToCobolCase(str)
}
func (s StringFuncs) ToTrainCase(str string) string {
	return _string.ToTrainCase(str)
}
func (s StringFuncs) ToNormalCase(str string) string {
	return _string.ToNormalCase(str)
}

func (s StringFuncs) Normalize(str string) string {
	return _string.ToNormalCase(str)
}

func (s StringFuncs) TrimSpace(str string) string {
	return _string.TrimSpace(str)
}
func (s StringFuncs) Trim(str string, cutset string) string {
	return _string.Trim(str, cutset)
}
func (s StringFuncs) TrimSuffix(str string, cutset string) string {
	return _string.TrimSuffix(str, cutset)
}
func (s StringFuncs) TrimPrefix(str string, cutset string) string {
	return _string.TrimPrefix(str, cutset)
}

func (s StringFuncs) Replace(str, old, new string, n int) string {
	return _string.Replace(str, old, new, n)
}

func (s StringFuncs) Split(str, sep string) []string {
	return _string.Split(str, sep)
}

func (s StringFuncs) Concat(sep string, items ...string) string {
	return _string.Concat(sep, items...)
}

func (s StringFuncs) Repeat(str string, count int) string {
	return _string.Repeat(str, count)
}

func (s StringFuncs) JoinFields(slice interface{}, sep, key string) string {
	return _string.JoinFields(slice, sep, key)
}
func (s StringFuncs) SplitFields(input, sep string) []interface{} {
	return _string.SplitFields(input, sep)
}

func (s StringFuncs) Pluralize(word string) string {
	return _string.Pluralize(word)
}
func (s StringFuncs) UnPluralize(word string) string {
	return _string.UnPluralize(word)
}
func (s StringFuncs) Indent(text, indent string) string {
	return _string.Indent(text, indent)
}
func (s StringFuncs) IndentAuto(text string) string {
	return _string.IndentAuto(text)
}

func (s StringFuncs) ArrayToStringSlice(arr []interface{}) []string {
	return _string.ArrayToStringSlice(arr)
}
