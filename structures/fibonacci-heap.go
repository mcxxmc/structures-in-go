package structures

import "math"

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

// inserts a new node to the root list. It does not update Min unless the heap is empty.
func (fib *FibonacciHeap) insert(newNode *FibNode) {
	if fib.Min == nil {  // creating a new root list
		newNode.Left = newNode
		newNode.Right = newNode
		fib.Min = newNode
	} else {  // insert into the existing root list
		min := fib.Min
		rightSibling := min.Right
		min.Right = newNode
		newNode.Left = min
		rightSibling.Left = newNode
		newNode.Right = rightSibling
	}
	fib.n ++
}

// (simply) removes a node from the root list. It does not modify the attributes of the node, and does not update Min
// unless the heap is empty after the deletion.
func (fib *FibonacciHeap) remove(node *FibNode) {
	if node.Left == node {  // is the only node
		fib.Min = nil
		fib.n = 0
	} else {
		left := node.Left
		right := node.Right
		left.Right = right
		right.Left = left
	}
}

// Insert inserts a new val into the heap.
func (fib *FibonacciHeap) Insert(val interface{}) {
	newNode := NewFibNode(val)
	fib.insert(newNode)
	if fib.compare(fib.Min.Val, newNode.Val) == 1 {  // replace fib.Min if necessary
		fib.Min = newNode
	}
}

// Minimum returns the min node of the heap; it won't change the heap.
func (fib *FibonacciHeap) Minimum() *FibNode {
	return fib.Min
}

// Union returns the union of the 2 fibonacci heaps.
//
// IMPORTANT: it will destroy the old ones.
func (fib *FibonacciHeap) Union(other *FibonacciHeap) *FibonacciHeap {
	h := NewFibonacciHeap(fib.compare)
	h.Min = fib.Min

	// combine the 2 root lists
	if fib.Min == nil {
		h.Min = other.Min
	} else if other.Min != nil {
		fibMin := fib.Min
		rightSibling := fibMin.Right
		otherMin := other.Min
		otherRightSibling := otherMin.Right

		fibMin.Right = otherRightSibling
		otherRightSibling.Left = fibMin
		otherMin.Right = rightSibling
		rightSibling.Left = otherMin

		// find new Min
		if fib.compare(fib.Min, other.Min) == 1 {
			h.Min = other.Min
		}
	}

	h.n = fib.n + other.n

	return h
}

// links node x and y and makes y a child of x. x and y should both be in the root list!
func (fib *FibonacciHeap) link(y, x *FibNode) {
	fib.remove(y)
	fib.appendChild(x, y)
	y.Marked = false
}

// append the child node to node x as a new child.
func (fib *FibonacciHeap) appendChild(x, child *FibNode) {
	if x.Child == nil {  // x does not have a child
		x.Child = child
		child.Left = child
		child.Right = child
		child.Parent = x
	} else {  // x has at least one child
		oldChild := x.Child
		rightSibling := x.Child.Right

		oldChild.Right = child
		child.Left = oldChild
		rightSibling.Left = child
		child.Right = rightSibling

		child.Parent = x
	}
	x.Degree ++
}

// consolidating the root nodes by reducing the number of nodes in the root list repeatedly.
func (fib *FibonacciHeap) consolidate() {
	dn := int(math.Log2(float64(fib.n)))  // the upper boundary
	a := make([]*FibNode, dn + 1)

	fib.Min.Left.Right = nil  // break the root list
	cur := fib.Min
	for cur != nil {  // loop through all the root nodes
		x := cur
		d := x.Degree

		for a[d] != nil {  // find root nodes with the same degree and link them together
			y := a[d]
			if fib.compare(x.Val, y.Val) == 1 {
				x, y = y, x
				cur = x  // update cur here as well
			}
			fib.link(y, x)
			a[d] = nil
			d ++
		}

		a[d] = x
		cur = cur.Right
	}

	fib.Min = nil

	for i := 0; i < len(a); i ++ {
		if a[i] != nil {
			fib.insert(a[i])
			if fib.compare(fib.Min.Val, a[i].Val) == 1 {
				fib.Min = a[i]
			}
		}
	}
}

// ExtractMin pops out the minimum node of the heap; it will change the heap.
func (fib *FibonacciHeap) ExtractMin() *FibNode {
	z := fib.Min
	if z != nil {
		child := z.Child
		if child != nil {
			child.Left.Right = nil  // break the child list
		}
		for child != nil {
			nextChild := child.Right
			fib.insert(child)  // this will change the siblings of the child
			child.Parent = nil
			child = nextChild
		}
		fib.remove(z)
		if z == z.Right {  // is the only node
			fib.Min = nil
		} else {
			fib.Min = z.Right
			fib.consolidate()
		}
		fib.n --
	}
	return z
}

func NewFibonacciHeap(compare func(a, b interface{}) int) *FibonacciHeap {
	return &FibonacciHeap{compare: compare}
}


