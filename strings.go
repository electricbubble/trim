package trim

import (
	"reflect"
	"strings"
)

func StringSpace(value any) {
	if value == nil {
		return
	}

	rVal := reflect.ValueOf(value)

	if rVal.Kind() == reflect.Invalid || rVal.Kind() != reflect.Ptr {
		return
	}

	rVal = rVal.Elem()
	if !rVal.CanSet() {
		return
	}

	if rVal.Kind() == reflect.Interface {
		if !rVal.IsValid() || rVal.IsNil() || rVal.IsZero() {
			return
		}
		rVal = reflect.ValueOf(rVal.Addr().Interface())
	}

	switch rVal.Kind() {
	case reflect.Struct:
		for i := 0; i < rVal.NumField(); i++ {
			if !rVal.Field(i).Addr().CanInterface() {
				continue
			}
			StringSpace(rVal.Field(i).Addr().Interface())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rVal.Len(); i++ {
			elem := rVal.Index(i)
			if !elem.Addr().CanInterface() {
				break
			}
			StringSpace(elem.Addr().Interface())
		}
	case reflect.String:
		rVal.SetString(strings.TrimSpace(rVal.String()))
	case reflect.Ptr:
		rv := reflect.ValueOf(value).Elem()
		if rv.IsNil() {
			return
		}
		StringSpace(rv.Interface())
	default:
		return
	}
}
