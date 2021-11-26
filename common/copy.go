package common

// Copy tries to make a deep copy of the input
//
// currently, only supports basic types
func Copy(val interface{}) interface{} {
	switch val.(type) {
	case int:
		return val.(int)
	case string:
		return val.(string)[:]
	default:
		if tmp, ok := val.(Template); ok {
			return tmp.Copy()
		}
		return val
	}
}
