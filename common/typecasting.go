package common

import (
	"strconv"
	"strings"
)

const DefaultStringLength = 4

// CastString converts an interface to a string
//
// For now, only support basic types.
func CastString(val interface{}) string {
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

// FixedMinLength converts the string to a minimum length of n by appending " " before and after
func FixedMinLength(s string, n int) string {
	ls := len(s)
	if ls >= n {
		return s
	} else {
		diff1 := (n - ls) / 2
		diff2 := n - ls - diff1
		return strings.Repeat(" ", diff1) + s + strings.Repeat(" ", diff2)
	}
}

// FixedMinLengthDefault converts the string to a minimum default length by appending " " before and after
func FixedMinLengthDefault(s string) string {
	return FixedMinLength(s, DefaultStringLength)
}