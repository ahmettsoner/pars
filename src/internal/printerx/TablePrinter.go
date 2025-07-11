package printerx

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

type TablePrinter struct{}

func (p *TablePrinter) Print(out io.Writer, data any) error {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("table output requires a slice")
	}
	if v.Len() == 0 {
		fmt.Fprintln(out, "(no data)")
		return nil
	}

	elemType := v.Index(0).Type()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	// 1. Header
	headers := make([]string, elemType.NumField())
	maxWidths := make([]int, elemType.NumField())

	for i := 0; i < elemType.NumField(); i++ {
		headers[i] = elemType.Field(i).Name
		maxWidths[i] = len(headers[i])
	}

	// 2. Rows
	rows := [][]string{}
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		row := make([]string, elem.NumField())
		for j := 0; j < elem.NumField(); j++ {
			val := fmt.Sprintf("%v", elem.Field(j).Interface())

			// Emoji destekli status gösterimi
			if strings.EqualFold(headers[j], "Status") {
				switch strings.ToLower(val) {
				case "active":
					val = "✅ Active"
				case "inactive":
					val = "❌ Inactive"
				}
			}

			row[j] = val
			if len(val) > maxWidths[j] {
				maxWidths[j] = len(val)
			}
		}
		rows = append(rows, row)
	}

	// 3. Print header
	headerLine := ""
	sepLine := ""
	for i, h := range headers {
		format := fmt.Sprintf("%%-%ds  ", maxWidths[i])
		headerLine += fmt.Sprintf(format, h)
		sepLine += strings.Repeat("─", maxWidths[i]) + "  "
	}
	fmt.Fprintln(out, headerLine)
	fmt.Fprintln(out, sepLine)

	// 4. Print rows
	for _, row := range rows {
		line := ""
		for i, val := range row {
			format := fmt.Sprintf("%%-%ds  ", maxWidths[i])
			line += fmt.Sprintf(format, val)
		}
		fmt.Fprintln(out, line)
	}

	return nil
}
