package common

import (
	"errors"
	"reflect"
	"strconv"
)

// ToString converts an interface to a string
//
// For now, only support basic types.
func ToString(val interface{}) string {
	switch val.(type) {
	case string:
		return val.(string)
	case int:
		return strconv.Itoa(val.(int))
	case nil:
		return " "
	default:
		if tmp, ok := val.(Template); ok {
			return tmp.String()
		}
		return " "
	}
}

// ToInterfaces converts an interface{} to []interface{} if applicable
// TODO: more cases
func ToInterfaces(values interface{}) ([]interface{}, error) {
	if rt := reflect.TypeOf(values); rt.Kind() != reflect.Slice && rt.Kind() != reflect.Array {
		err := errors.New("values should be a slice or array")
		return nil, err
	}
	switch values.(type) {
	case []int:
		tmp := values.([]int)
		r := make([]interface{}, len(tmp))
		for i := 0; i < len(tmp); i ++ {
			r[i] = tmp[i]
		}
		return r, nil
	case []interface{}:
		return values.([]interface{}), nil
	case []Template:
		tmp := values.([]Template)
		r := make([]interface{}, len(tmp))
		for i := 0; i < len(tmp); i ++ {
			r[i] = tmp[i]
		}
		return r, nil
	default:
		err := errors.New("cannot handle the input type")
		return nil, err
	}
}
