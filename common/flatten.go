package common

// Flatten2D turns a 2d slice into a 1d one. Note that it make deep copies.
func Flatten2D(values [][]interface{}) []interface{} {
	ans := make([]interface{}, 0)
	for _, layer := range values {
		tmp := make([]interface{}, len(layer))
		copy(tmp, layer)
		ans = append(ans, tmp...)
	}
	return ans
}
