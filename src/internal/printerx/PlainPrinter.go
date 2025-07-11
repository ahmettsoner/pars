package printerx

import (
	"fmt"
	"io"
	"reflect"
)

type PlainPrinter struct{}

func (p *PlainPrinter) Print(out io.Writer, data any) error {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("plain output requires a slice")
	}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		fmt.Fprintln(out, "-")
		for j := 0; j < item.NumField(); j++ {
			field := item.Type().Field(j)
			value := item.Field(j)
			fmt.Fprintf(out, "  %s: %v\n", field.Name, value.Interface())
		}
	}
	return nil
}
