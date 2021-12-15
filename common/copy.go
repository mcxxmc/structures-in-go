package common

// Copy tries to make a deep copy of the input
//
// currently, only supports basic types  //TODO: more cases
func Copy(val interface{}) interface{} {
	switch val.(type) {
	case int:
		return val.(int)
	case string:
		return val.(string)[:]
	case []int:
		tmp := make([]interface{}, len(val.([]int)))
		for i, _ := range val.([]int) {
			tmp[i] = val.([]int)[i]
		}
		return tmp
	case []interface{}:
		tmp := make([]interface{}, len(val.([]interface{})))
		for i, _ := range val.([]interface{}) {
			tmp[i] = Copy(val.([]interface{})[i])
		}
		return tmp
	default:
		if tmp, ok := val.(Template); ok {
			return tmp.Copy()
		}
		return val
	}
}

// CopyInterfaces CopyList tries to make a deep copy of a []interface object.
func CopyInterfaces(val []interface{}) []interface{} {
	tmp := make([]interface{}, len(val))
	for i, _ := range val {
		tmp[i] = Copy(val[i])
	}
	return tmp
}
