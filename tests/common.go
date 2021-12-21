package tests

func compareInt (a, b interface{}) int {
	if a.(int) > b.(int) {
	return 1
	} else if a.(int) == b.(int) {
	return 0
	}
	return -1
}
