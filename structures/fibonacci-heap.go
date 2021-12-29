package structures

// FibonacciHeap
//
// The fibonacci heap structure, which is an ordered min-heap. Please use NewFibonacciHeap() as the safe constructor.
//
// Attributes:
//
// Min *FibNode: a pointer to the minimum node.
//
// compare func(a, b interface{}) int
//
// .
//
// compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b.
//
// .
//
// The first input of the compare function should be the same type as the value of the tree node; the second input may have
// variant types. A tricky compare method can relax the conditions for Search and Delete; see examples for details.
type FibonacciHeap struct {
	n int
	Min *FibNode
	compare func(a, b interface{}) int
}

func NewFibonacciHeap(compare func(a, b interface{}) int) *FibonacciHeap {
	return &FibonacciHeap{compare: compare}
}


