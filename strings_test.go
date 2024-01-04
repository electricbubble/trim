package trim

import (
	"io"
	"testing"
)

func TestStringSpace(t *testing.T) {
	t.Run("nil_any", func(t *testing.T) {
		StringSpace(nil)

		var v any
		StringSpace(v)
		requireNil(t, v)
		StringSpace(&v)
		requireNil(t, v)
	})
	t.Run("var_error", func(t *testing.T) {
		var v error
		StringSpace(v)
		requireNil(t, v)
		StringSpace(&v)
		requireNil(t, v)

		err := io.EOF
		StringSpace(&v)
		requireEqual(t, err.Error(), io.EOF.Error())
	})
	t.Run("int", func(t *testing.T) {
		v := 6
		StringSpace(v)
		requireEqual(t, 6, v)
		StringSpace(&v)
		requireEqual(t, 6, v)
	})
	t.Run("chan", func(t *testing.T) {
		v := make(chan struct{}, 1)
		StringSpace(v)
		StringSpace(&v)
	})

	t.Run("string", func(t *testing.T) {
		v := "a  "
		expected := "a  "
		StringSpace(v)
		requireEqual(t, expected, v)
	})
	t.Run("&string", func(t *testing.T) {
		v := "a  "
		expected := "a"
		StringSpace(&v)
		requireEqual(t, expected, v)
	})
	t.Run("*string", func(t *testing.T) {
		v := toPtr("a  ")
		expected := toPtr("a")
		StringSpace(v)
		requireEqual(t, *expected, *v)
	})
	t.Run("&(*string)", func(t *testing.T) {
		v := toPtr("a  ")
		expected := toPtr("a")
		StringSpace(&v)
		requireEqual(t, *expected, *v)
	})

	t.Run("any(string)", func(t *testing.T) {
		v := any("  b  ")
		expected := any("  b  ")
		StringSpace(v)
		requireEqual(t, expected, v)
	})
	t.Run("&(any(string))", func(t *testing.T) {
		v := any("  b  ")
		expected := any("  b  ")
		StringSpace(&v)
		requireEqual(t, expected, v)
	})
	t.Run("*any(string)", func(t *testing.T) {
		v := toPtr(any("  b  "))
		expected := any("  b  ")
		StringSpace(v)
		requireEqual(t, expected, *v)
	})
	t.Run("&(*any(string))", func(t *testing.T) {
		v := toPtr(any("  b  "))
		expected := any("  b  ")
		StringSpace(&v)
		requireEqual(t, expected, *v)
	})

	t.Run("any(*string)", func(t *testing.T) {
		v := any(toPtr("  c  "))
		expected := any(toPtr("c"))
		StringSpace(v)
		requireEqual(t, *(expected.(*string)), *(v.(*string)))
	})
	t.Run("&(any(*string))", func(t *testing.T) {
		v := any(toPtr("  c  "))
		expected := any(toPtr("c"))
		StringSpace(&v)
		requireEqual(t, *(expected.(*string)), *(v.(*string)))
	})
	t.Run("*any(*string)", func(t *testing.T) {
		v := toPtr(any(toPtr("  c  ")))
		expected := toPtr(any(toPtr("c")))
		StringSpace(v)
		requireEqual(t, *((*expected).(*string)), *((*v).(*string)))
	})
	t.Run("&(*any(*string))", func(t *testing.T) {
		v := toPtr(any(toPtr("  c  ")))
		expected := toPtr(any(toPtr("c")))
		StringSpace(&v)
		requireEqual(t, *((*expected).(*string)), *((*v).(*string)))
	})

	t.Run("[]string", func(t *testing.T) {
		vv := []string{"         a ", " b", " c "}
		expected := []string{"         a ", " b", " c "}
		StringSpace(vv)
		requireEqual(t, expected, vv)
	})
	t.Run("&[]string", func(t *testing.T) {
		vv := []string{"         a ", " b            ", " c "}
		expected := []string{"a", "b", "c"}
		StringSpace(&vv)
		requireEqual(t, expected, vv)
	})
	t.Run("*[]string", func(t *testing.T) {
		vv := toPtr([]string{"         a ", " b", " c           "})
		expected := []string{"a", "b", "c"}
		StringSpace(vv)
		requireEqual(t, expected, *vv)
	})
	t.Run("&(*[]string)", func(t *testing.T) {
		vv := toPtr([]string{"         a ", " b", " c ", "d"})
		expected := []string{"a", "b", "c", "d"}
		StringSpace(&vv)
		requireEqual(t, expected, *vv)
	})

	t.Run("[]any(string)", func(t *testing.T) {
		vv := []any{"  q", " w ", " e ", "      "}
		expected := []any{"  q", " w ", " e ", "      "}
		StringSpace(vv)
		requireEqual(t, expected, vv)
	})
	t.Run("&([]any(string))", func(t *testing.T) {
		vv := []any{"  q", " w ", " e ", "      "}
		expected := []any{"  q", " w ", " e ", "      "}
		StringSpace(&vv)
		requireEqual(t, expected, vv)
	})
	t.Run("[]any(*string)", func(t *testing.T) {
		vv := []any{toPtr("  q"), toPtr(" w "), toPtr(" e "), toPtr("      ")}
		expected := []any{"  q", " w ", " e ", "      "}
		StringSpace(vv)
		for i, a := range vv {
			requireEqual(t, expected[i].(string), *(a.(*string)))
		}
	})
	t.Run("&([]any(*string))", func(t *testing.T) {
		vv := []any{toPtr("  q"), toPtr(" w "), toPtr(" e "), toPtr("      ")}
		expected := []any{"q", "w", "e", ""}
		StringSpace(&vv)
		for i, a := range vv {
			requireEqual(t, expected[i].(string), *(a.(*string)))
		}
	})

	t.Run("type_string", func(t *testing.T) {
		var v Model = " a1 "
		expected := Model("a1")
		StringSpace(v)
		requireEqual(t, Model(" a1 "), v)
		StringSpace(&v)
		requireEqual(t, expected, v)
	})
	t.Run("[]type_string", func(t *testing.T) {
		vv := []Model{"   a1 ", " b2", "   "}
		expected := []Model{"   a1 ", " b2", "   "}
		StringSpace(vv)
		requireEqual(t, expected, vv)
	})
	t.Run("&[]type_string", func(t *testing.T) {
		vv := []Model{"   a1 ", " b2", "   "}
		expected := []Model{"a1", "b2", ""}
		StringSpace(&vv)
		requireEqual(t, expected, vv)
	})
	t.Run("[](*type_string)", func(t *testing.T) {
		vv := []*Model{toPtr[Model]("   a1 "), toPtr[Model](" b2"), toPtr[Model]("   "), nil}
		expected := []*Model{toPtr[Model]("   a1 "), toPtr[Model](" b2"), toPtr[Model]("   "), nil}
		StringSpace(vv)
		requireEqual(t, expected, vv)
	})
	t.Run("&([](*type_string))", func(t *testing.T) {
		vv := []*Model{toPtr[Model]("   a1 "), toPtr[Model](" b2"), toPtr[Model]("   "), nil}
		expected := []*Model{toPtr[Model]("a1"), toPtr[Model]("b2"), toPtr[Model](""), nil}
		StringSpace(&vv)
		requireEqual(t, expected, vv)
	})
	t.Run("&([3](*type_string))", func(t *testing.T) {
		vv := [3]Model{"a1 ", " b2", " c3 "}
		expected := [3]Model{"a1", "b2", "c3"}
		StringSpace(&vv)
		requireEqual(t, expected, vv)
	})

	t.Run("struct", func(t *testing.T) {
		v := Conf{
			URL:      "   https://hi.bye",
			Port:     8080,
			token:    "    token       ",
			internal: internal{secret: "  d"},
			Public:   Public{Name: "    hi"},
		}
		expected := Conf{
			URL:      "   https://hi.bye",
			Port:     8080,
			token:    "    token       ",
			internal: internal{secret: "  d"},
			Public:   Public{Name: "    hi"},
		}
		StringSpace(v)
		requireEqual(t, expected, v)
	})
	t.Run("&struct", func(t *testing.T) {
		v := Conf{
			URL:      "   https://hi.bye",
			Port:     8080,
			token:    "    token       ",
			internal: internal{secret: "  d"},
			Public:   Public{Name: "    hi"},
		}
		expected := Conf{
			URL:      "https://hi.bye",
			Port:     8080,
			token:    "    token       ",
			internal: internal{secret: "  d"},
			Public:   Public{Name: "hi"},
		}
		StringSpace(&v)
		requireEqual(t, expected, v)

	})

	t.Run("&S.S", func(t *testing.T) {
		v := Conf{
			URL:    "   https://hi.bye",
			Public: Public{Name: "     hi "},
		}
		expected := Conf{
			URL:    "   https://hi.bye",
			Public: Public{Name: "hi"},
		}
		StringSpace(&v.Public)
		requireEqual(t, expected, v)
	})
}

type Model string

func toPtr[T any](v T) *T {
	return &v
}

type Conf struct {
	URL      string
	Port     int
	token    string
	internal internal

	Public Public
}

type internal struct {
	secret string
}

type Public struct {
	Name string
}
