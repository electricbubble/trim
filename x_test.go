package trim

import (
	"reflect"
	"slices"
	"testing"
)

func requireNil(t *testing.T, v any) {
	t.Helper()

	if v == nil {
		return
	}

	rVal := reflect.ValueOf(v)
	rKind := rVal.Kind()

	isNilableKind := slices.Contains([]reflect.Kind{
		reflect.Ptr, reflect.UnsafePointer,
		reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface,
	}, rKind)

	if isNilableKind && rVal.IsNil() {
		t.Fatalf("Expected nil, but got: %#v", v)
	}
}

func requireEqual(t *testing.T, expected, actual any) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual)
	}
}
