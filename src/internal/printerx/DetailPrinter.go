package printerx

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"parsdevkit.net/pkg/utilities/object"
)

// type Resource struct {
// 	Name      string
// 	Type      string
// 	Status    string
// 	CreatedAt string
// 	Secure    bool
// 	Path      string
// }
// res := Resource{
// 	Name: "app-config", Type: "Config", Status: "Active",
// 	CreatedAt: "2025-07-09 10:30", Secure: true, Path: "/dev/app/app-config",
// }

// printer := &printerx.DetailPrinter{}
// printer.Print(os.Stdout, res)

type DetailPrinter struct {
	Emoji     bool
	ShowEmpty bool
	Indent    int
}

func (p *DetailPrinter) Print(out io.Writer, obj any) error {
	return p.printValue(out, obj, p.Indent)
}

func (p *DetailPrinter) printValue(out io.Writer, obj any, indent int) error {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	maxKeyLen := 0
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		if len(name) > maxKeyLen {
			maxKeyLen = len(name)
		}
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		isEmpty := object.IsEmpty(value)
		if isEmpty && !p.ShowEmpty {
			continue // boş ve gösterme
		}
		displayValue := value
		if isEmpty && p.ShowEmpty {
			displayValue = object.EmptyValue(value)
		}

		key := field.Name
		icon := ""
		if p.Emoji {
			icon = emojiForField(key, strings.ToLower(key))
		}

		prefix := strings.Repeat("  ", indent)
		format := fmt.Sprintf("%%s%%s%%-%ds : ", maxKeyLen)

		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Slice && rv.Type().Elem().Kind() == reflect.Struct {
			if rv.Len() == 0 {
				fmt.Fprintf(out, format+"(empty)\n", prefix, icon, key)
				continue
			}

			fmt.Fprintf(out, format+"\n", prefix, icon, key)
			for i := 0; i < rv.Len(); i++ {
				item := rv.Index(i).Interface()
				itemVal := reflect.ValueOf(item)
				if itemVal.Kind() == reflect.Ptr {
					itemVal = itemVal.Elem()
				}
				itemType := itemVal.Type()

				// İlk alanı "- Name : Value" gibi yaz
				if itemVal.NumField() > 0 {
					firstField := itemType.Field(0)
					firstValue := itemVal.Field(0).Interface()

					displayFirst := firstValue
					if object.IsEmpty(firstValue) && p.ShowEmpty {
						displayFirst = object.EmptyValue(firstValue)
					}

					firstKey := firstField.Name
					firstIcon := ""
					if p.Emoji {
						firstIcon = emojiForField(firstKey, strings.ToLower(firstKey))
					}

					// "-" işareti + ilk alan
					fmt.Fprintf(out, "%s  - %s%-10s : %v\n", prefix, firstIcon, firstKey, displayFirst)

					// Diğer alanları alt satırlara yaz
					for j := 1; j < itemVal.NumField(); j++ {
						subField := itemType.Field(j)
						subValue := itemVal.Field(j).Interface()

						if object.IsEmpty(subValue) && !p.ShowEmpty {
							continue
						}
						if object.IsEmpty(subValue) && p.ShowEmpty {
							subValue = object.EmptyValue(subValue)
						}

						subKey := subField.Name
						subIcon := ""
						if p.Emoji {
							subIcon = emojiForField(subKey, strings.ToLower(subKey))
						}
						fmt.Fprintf(out, "%s    %s%-10s : %v\n", prefix, subIcon, subKey, subValue)
					}
				}
			}
			continue
		}

		// 🔥 Bu eksikti: normal alanları yazdır
		fmt.Fprintf(out, format+"%v\n", prefix, icon, key, displayValue)
	}
	return nil
}

func emojiForField(key string, value string) string {
	key = strings.ToLower(key)
	switch key {
	case "name":
		return "📛 "
	case "type":
		return "📦 "
	case "status":
		if strings.Contains(strings.ToLower(value), "active") {
			return "✅ "
		}
		return "❌ "
	case "createdat", "created":
		return "📅 "
	case "secure":
		return "🔐 "
	case "path":
		return "📁 "
	case "tags":
		return "🏷️ "
	case "projects":
		return "📂 "
	default:
		return ""
	}
}
