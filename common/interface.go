package common

// Value the interface for values
//
// can be used as:
//
// 1. Val in trees
type Value interface {
	String() string     // stringify
	Copy() interface{}  // makes a deep copy
}
