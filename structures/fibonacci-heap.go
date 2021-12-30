package structures

import (
	"errors"
	"math"
)
//TODO: make a node list a real list with loop

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

// NumOfElements returns the number of nodes in the heap.
func (fib *FibonacciHeap) NumOfElements() int {
	return fib.n
}

// inserts a new node to the root list. It does not update fib.n. It does not update Min unless the heap is empty.
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
}

// (simply) removes a node from the root list. It does not modify the attributes of the node and fib.n,
// and does not update Min unless the heap is empty after the deletion.
func (fib *FibonacciHeap) removeFromRoot(node *FibNode) {
	if node.Right == node {  // is the only node
		fib.Min = nil
	} else {
		left := node.Left
		right := node.Right
		left.Right = right
		right.Left = left
	}
}

// Insert inserts a new val into the heap and returns a pointer to the inserted node.
func (fib *FibonacciHeap) Insert(val interface{}) *FibNode {
	newNode := NewFibNode(val)
	fib.insert(newNode)
	fib.n ++
	if fib.compare(fib.Min.Val, newNode.Val) == 1 {  // replace fib.Min if necessary
		fib.Min = newNode
	}
	return newNode
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
		if fib.compare(fib.Min.Val, other.Min.Val) == 1 {
			h.Min = other.Min
		}
	}

	h.n = fib.n + other.n

	return h
}

// links node x and y and makes y a child of x. x and y should both be in the root list!
func (fib *FibonacciHeap) link(y, x *FibNode) {
	fib.removeFromRoot(y)
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

	cur := fib.Min
	count := make(map[*FibNode]bool)
	for cur != nil && !count[cur] {  // loop through all the root nodes
		x := cur
		right := cur.Right
		count[x] = true
		d := x.Degree

		for d < dn && a[d] != nil {  // find root nodes with the same degree and link them together
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
		cur = right
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
			child.Left.Right = nil  // break the child list; it is ok to break here because we will reset siblings later
		}
		for child != nil {
			nextChild := child.Right
			fib.insert(child)  // this will change the siblings of the child
			child.Parent = nil
			child = nextChild
		}
		fib.removeFromRoot(z)
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

// removes x from the child list of y and decrease the degree of y by 1. It does not update x.
func (fib *FibonacciHeap) removeChild(x, y *FibNode) {
	if y.Degree == 1 {
		y.Child = nil
	} else if y.Child == x {
		y.Child = x.Right
	}
	// reconnect the list
	x.Left.Right = x.Right
	x.Right.Left = x.Left
	y.Degree --
}

// cuts the link between node x and its parent node y, making x a root node.
func (fib *FibonacciHeap) cut(x, y *FibNode) {
	fib.removeChild(x, y)
	fib.insert(x)
	x.Parent = nil
	x.Marked = false
}

// stops until reaching a root node or an unmarked node.
func (fib *FibonacciHeap) cascadingCut(y *FibNode) {
	z := y.Parent
	if z != nil {
		if !y.Marked {
			y.Marked = true
		} else {
			fib.cut(y, z)
			fib.cascadingCut(z)
		}
	}
}

// DecreaseKey decrease the val of a node to newVal.
func (fib *FibonacciHeap) DecreaseKey(node *FibNode, newVal interface{}) error {
	if fib.compare(newVal, node.Val) == 1 {
		return errors.New("new key cannot be greater than the old key")
	}

	node.Val = newVal
	y := node.Parent

	if y != nil && fib.compare(y.Val, node.Val) == 1 {
		fib.cut(node, y)
		fib.cascadingCut(y)
	}
	if fib.compare(fib.Min.Val, node.Val) == 1 {
		fib.Min = node
	}

	return nil
}

// Delete deletes a node from the heap.
//
// It requires a min interface{} as input so for all the values in the heap, FibonacciHeap.compare(a, min) == 1.
func (fib *FibonacciHeap) Delete(node *FibNode, min interface{}) error {
	err := fib.DecreaseKey(node, min)
	if err != nil {
		return err
	}
	fib.ExtractMin()
	return nil
}

func NewFibonacciHeap(compare func(a, b interface{}) int) *FibonacciHeap {
	return &FibonacciHeap{compare: compare}
}


