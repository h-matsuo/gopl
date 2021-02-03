package params

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Pack(v interface{}) (string, error) {
	return pack(reflect.ValueOf(v))
}

func pack(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.Struct:
		// ok; continue process
	case reflect.Ptr:
		return pack(v.Elem())
	default:
		return "", fmt.Errorf("root type must be struct, got: %s", v.Type())
	}

	values := &url.Values{}
	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		name := structField.Tag.Get("http")
		if name == "" {
			name = strings.ToLower(structField.Name)
		}
		if err := addQuery(values, name, v.Field(i)); err != nil {
			return "", err
		}
	}

	return values.Encode(), nil
}

func addQuery(values *url.Values, name string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		values.Add(name, "")

	case reflect.String:
		values.Add(name, v.String())

	case reflect.Int:
		values.Add(name, fmt.Sprintf("%d", v.Int()))

	case reflect.Bool:
		values.Add(name, fmt.Sprintf("%t", v.Bool()))

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := addQuery(values, name, v.Index(i)); err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
