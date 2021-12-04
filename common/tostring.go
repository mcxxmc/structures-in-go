package common

import (
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
