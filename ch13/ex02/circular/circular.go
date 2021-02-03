package circular

import (
	"reflect"
	"unsafe"
)

func IsCircular(v interface{}) bool {
	seen := make(map[reference]bool)
	return isCircular(reflect.ValueOf(v), seen)
}

type reference struct {
	v unsafe.Pointer
	t reflect.Type
}

func isCircular(v reflect.Value, seen map[reference]bool) bool {
	if !v.IsValid() {
		return false
	}
	if v.CanAddr() {
		vptr := unsafe.Pointer(v.UnsafeAddr())
		r := reference{vptr, v.Type()}
		if seen[r] {
			return true // already seen
		}
		seen[r] = true
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCircular(v.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if isCircular(v.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if isCircular(v.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range v.MapKeys() {
			if isCircular(v.MapIndex(k), seen) {
				return true
			}
		}
		return false

	default:
		return false
	}
}
