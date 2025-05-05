package string

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func StartsWith(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

func EndsWith(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

func ToUpperCase(str string) string {
	return strings.ToUpper(str)
}

func ToLowerCase(str string) string {
	return strings.ToLower(str)
}

func ToPascalCase(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, "")
}
func ToCamelCase(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word[:1]) + word[1:]
		} else {
			words[i] = strings.ToLower(word)
		}
	}
	return strings.Join(words, "")
}

func ToSnakeCase(str string) string {
	var builder strings.Builder
	var prev rune

	for _, char := range str {
		if unicode.IsUpper(char) && unicode.IsLower(prev) {
			builder.WriteRune('_')
		}
		builder.WriteRune(unicode.ToLower(char))
		prev = char
	}

	return builder.String()
}
func ToKebabCase(str string) string {
	words := strings.Fields(str)

	var kebabCaseWords []string
	for _, word := range words {
		kebabCaseWords = append(kebabCaseWords, strings.ToLower(word))
	}

	return strings.Join(kebabCaseWords, "-")
}
func ToUpperFlatCase(str string) string {
	str = strings.ReplaceAll(str, " ", "")

	return strings.ToUpper(str)
}
func ToLowerFlatCase(str string) string {
	str = strings.ReplaceAll(str, " ", "")

	return strings.ToLower(str)
}
func ToCobolCase(str string) string {
	str = strings.ReplaceAll(str, " ", "-")

	return strings.ToUpper(str)
}
func ToTrainCase(str string) string {
	var builder strings.Builder
	prevWasSpace := true

	for _, char := range str {
		if unicode.IsSpace(char) {
			builder.WriteRune('-')
			prevWasSpace = true
		} else if prevWasSpace {
			builder.WriteRune(unicode.ToUpper(char))
			prevWasSpace = false
		} else {
			builder.WriteRune(unicode.ToLower(char))
		}
	}

	return builder.String()
}
func ToNormalCase(str string) string {
	str = strings.ReplaceAll(str, "-", " ")
	str = strings.ReplaceAll(str, "_", " ")

	str = strings.Title(strings.ToLower(str))

	str = strings.TrimSpace(str)

	return str
}

func Normalize(str string) string {
	return strings.TrimSpace(strings.ReplaceAll(str, "\n", " "))
}

func TrimSpace(str string) string {
	return strings.TrimSpace(str)
}
func Trim(str string, cutset string) string {
	return strings.Trim(str, cutset)
}
func TrimSuffix(str string, cutset string) string {
	return strings.TrimRight(str, cutset)
}
func TrimPrefix(str string, cutset string) string {
	return strings.TrimRight(str, cutset)
}

func Replace(str, old, new string, n int) string {
	return strings.Replace(str, old, new, n)
}

func Split(str, sep string) []string {
	return strings.Split(str, sep)
}

func Concat(sep string, items ...string) string {
	return strings.Join(items, sep)
}

func Repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

func JoinFields(slice interface{}, sep, key string) string {
	val := reflect.ValueOf(slice)
	if val.Type().Kind() != reflect.Slice {
		return ""
	}

	var result []string
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i)
		if item.Type().Kind() == reflect.Struct && key != "" {
			fieldValue := item.FieldByName(key)
			if !fieldValue.IsValid() {
				continue
			}
			result = append(result, fmt.Sprint(fieldValue.Interface()))
		} else {
			if item.CanInterface() {
				result = append(result, fmt.Sprint(item.Interface()))
			}
		}
	}

	return strings.Join(result, sep)
}
func SplitFields(input, sep string) []interface{} {
	parts := strings.Split(input, sep)

	var result []interface{}
	for _, part := range parts {
		result = append(result, part)
	}

	return result
}

func Pluralize(word string) string {
	irregulars := map[string]string{
		"person":   "people",
		"man":      "men",
		"woman":    "women",
		"child":    "children",
		"foot":     "feet",
		"tooth":    "teeth",
		"goose":    "geese",
		"mouse":    "mice",
		"ox":       "oxen",
		"deer":     "deer",
		"sheep":    "sheep",
		"fish":     "fish",
		"aircraft": "aircraft",
		"series":   "series",
	}

	switch word[len(word)-1:] {
	case "s", "x", "ch", "sh":
		return word + "es"
	case "y":
		if len(word) > 2 && word[len(word)-2] != 'a' {
			return word[:len(word)-1] + "ies"
		}
	}

	if plural, ok := irregulars[word]; ok {
		return plural
	}

	return word + "s"
}
func UnPluralize(word string) string {
	irregulars := map[string]string{
		"person":   "people",
		"man":      "men",
		"woman":    "women",
		"child":    "children",
		"foot":     "feet",
		"tooth":    "teeth",
		"goose":    "geese",
		"mouse":    "mice",
		"ox":       "oxen",
		"deer":     "deer",
		"sheep":    "sheep",
		"fish":     "fish",
		"aircraft": "aircraft",
		"series":   "series",
	}

	switch word[len(word)-1:] {
	case "s":
		return word[:len(word)-1]
	case "es":
		if len(word) > 2 && word[len(word)-2] != 'a' {
			return word[:len(word)-2] + "y"
		}
	case "ies":
		return word[:len(word)-3] + "y"
	}

	if singular, ok := irregulars[word]; ok {
		return singular
	}

	return word[:len(word)-1]
}
func Indent(text, indent string) string {
	lines := strings.Split(text, "\n")

	result := ""
	for _, line := range lines {
		result += indent + strings.TrimSpace(line) + "\n"
	}

	return result
}
func IndentAuto(text string) string {
	lines := strings.Split(text, "\n")

	if len(lines) > 0 {
		indent := getIndentation(lines[0])

		result := ""
		for _, line := range lines {
			result += indent + strings.TrimSpace(line) + "\n"
		}

		return result
	}

	return text
}
func getIndentation(line string) string {
	indent := ""
	for _, char := range line {
		if char == ' ' || char == '\t' {
			indent += string(char)
		} else {
			break
		}
	}
	return indent
}
func ArrayToStringSlice(arr []interface{}) []string {
	strSlice := make([]string, 0, len(arr))
	for _, val := range arr {
		if strVal, ok := val.(string); ok {
			strSlice = append(strSlice, strVal)
		}
	}
	return strSlice
}
