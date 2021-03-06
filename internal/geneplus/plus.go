package geneplus

import (
	"errors"
	"reflect"
	"strings"
)

var ErrInvalid = errors.New("go away")

func addWithInterface(a, b interface{}) (interface{}, error) {
	switch a.(type) {
	case int:
		if bint, ok := b.(int); ok {
			return a.(int) + bint, nil
		}
	case float64:
		if bfloat, ok := b.(float64); ok {
			return a.(float64) + bfloat, nil
		}
	case string:
		if bstring, ok := b.(string); ok {
			return a.(string) + bstring, nil
		}
	}

	return nil, ErrInvalid
}

func addReflected(a, b interface{}) (interface{}, error) {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		return a.(int) + b.(int), nil
	case reflect.Float64:
		return a.(float64) + b.(float64), nil
	case reflect.String:
		return a.(string) + b.(string), nil
	}

	return nil, ErrInvalid
}

func addInt(a, b int) int {
	return a + b
}

func addFloat64(a, b float64) float64 {
	return a + b
}

func addString(a, b string) string {
	return a + b
}

func addStringBuf(a, b string) string {
	r := strings.Builder{}
	r.WriteString(a)
	r.WriteString(b)
	return r.String()
}

func addWithInterfaceVari(list ...interface{}) (interface{}, error) {
	var a interface{}
	for _, some := range list {
		switch some.(type) {
		case int:
			if a == nil {
				a = 0
			}
			cur, ok := a.(int)
			if !ok {
				return nil, ErrInvalid
			}
			if item, ok := some.(int); ok {
				cur += item
			}
			a = cur
		case float64:
			if a == nil {
				a = float64(0)
			}
			cur, ok := a.(float64)
			if !ok {
				return nil, ErrInvalid
			}
			if item, ok := some.(float64); ok {
				cur += item
			}
			a = cur
		case string:
			if a == nil {
				a = ""
			}
			cur, ok := a.(string)
			if !ok {
				return nil, ErrInvalid
			}
			if item, ok := some.(string); ok {
				cur += item
			}
			a = cur
		}
	}

	return a, nil
}

func addReflectedVari(list ...interface{}) (interface{}, error) {
	var res interface{}

	for _, itm := range list {
		switch reflect.TypeOf(itm).Kind() {
		case reflect.Int:
			if res == nil {
				res = 0
			}
			cur, ok := res.(int)
			if !ok {
				return nil, ErrInvalid
			}

			if item, ok := itm.(int); ok {
				cur += item
			}

			res = cur
		case reflect.Float64:
			if res == nil {
				res = float64(0)
			}
			cur, ok := res.(float64)
			if !ok {
				return nil, ErrInvalid
			}

			if item, ok := itm.(float64); ok {
				cur += item
			}

			res = cur
		case reflect.String:
			if res == nil {
				res = ""
			}
			cur, ok := res.(string)
			if !ok {
				return nil, ErrInvalid
			}

			if item, ok := itm.(string); ok {
				cur += item
			}

			res = cur
		}
	}

	return res, nil
}

func addIntVari(list ...int) int {
	var res int
	for _, itm := range list {
		res += itm
	}

	return res
}

func addFloat64Vari(list ...float64) float64 {
	var res float64
	for _, itm := range list {
		res += itm
	}

	return res
}

func addStringVari(list ...string) string {
	var res string
	for _, itm := range list {
		res += itm
	}

	return res
}

func addStringBufVari(list ...string) string {
	r := strings.Builder{}
	for _, itm := range list {
		r.WriteString(itm)
	}

	return r.String()
}
