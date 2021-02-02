package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		if !v.IsZero() {
			fmt.Fprintf(buf, "%d", v.Int())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if !v.IsZero() {
			fmt.Fprintf(buf, "%d", v.Uint())
		}

	case reflect.String:
		if !v.IsZero() {
			fmt.Fprintf(buf, "%q", v.String())
		}

	case reflect.Ptr:
		if !v.IsZero() {
			return encode(buf, v.Elem())
		}

	case reflect.Array, reflect.Slice: // (value ...)
		existsNotZeroValue := false
		for i := 0; i < v.Len(); i++ {
			if !v.Index(i).IsZero() {
				existsNotZeroValue = true
				break
			}
		}
		if !existsNotZeroValue {
			return nil
		}

		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 && !v.Index(i).IsZero() {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		existsNotZeroValue := false
		for i := 0; i < v.NumField(); i++ {
			if !v.Field(i).IsZero() {
				existsNotZeroValue = true
				break
			}
		}
		if !existsNotZeroValue {
			return nil
		}

		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).IsZero() {
				continue
			}

			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		existsNotZeroValue := false
		for _, key := range v.MapKeys() {
			if !key.IsZero() {
				existsNotZeroValue = true
				break
			}
		}
		if !existsNotZeroValue {
			return nil
		}

		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if v.MapIndex(key).IsZero() {
				continue
			}

			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
