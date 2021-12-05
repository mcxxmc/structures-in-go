package structures

import (
	"some-data-structures/common"
)

// BinarySearchTree the binary search tree
//
// Attributes:
//
//     Root *Node
//
//     compare func(a, b interface{}) bool
//
// .
//
// compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b
//
// .
//
// the first input of the compare function should be the same type as the value of the tree node; the second input may have
// variant types. A tricky compare method can relax the conditions for Search and Delete; see examples for details.
//
// .
//
// note that this BinarySearchTree does not perform type checking; please include any necessary type checking
// in the customized compare function
type BinarySearchTree struct {
	Root    *Node
	compare func(a, b interface{}) int
}

// Insert inserts a new val as a new node
func (bt *BinarySearchTree) Insert(val interface{}) {
	if bt.Root == nil {
		bt.Root = NewNode(val)
		return
	}

	cur := bt.Root
	for {
		if bt.compare(cur.Val, val) == 1 { // cur.Val > val
			if cur.Left == nil {
				cur.Left = NewNode(val)
				break
			} else {
				cur = cur.Left
			}
		} else {
			if cur.Right == nil {
				cur.Right = NewNode(val)
				break
			} else {
				cur = cur.Right
			}
		}
	}
}

// Search returns the value of the FIRST corresponding Node if that Node exists in the tree
func (bt *BinarySearchTree) Search(val interface{}) (interface{}, bool) {
	cur := bt.Root
	for {
		if cur == nil {
			break
		}
		c := bt.compare(cur.Val, val)
		if c == 0 {
			return cur.Val, true
		} else if c == 1 {  // cur.Val > val
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return nil, false
}

// Delete deletes the node if it exists; note that it only deletes 1 node at once
//
// returns the val of the deleted node and a boolean value indicating if the deletion is successful
func (bt *BinarySearchTree) Delete(val interface{}) (interface{}, bool) {
	parent := DummyNode()
	isLeftChild := false
	cur := bt.Root

	// first find the node to delete
	for {
		if cur == nil {
			break
		}

		c := bt.compare(cur.Val, val)

		if c == 0 {
			break
		} else if c == 1 {
			parent = cur
			cur = cur.Left
			isLeftChild = true
		} else {
			parent = cur
			cur = cur.Right
			isLeftChild = false
		}
	}

	// if the element is not found
	if cur == nil {
		return nil, false
	}

	// if the node to delete does not have a left child (it may also do not have a right child)
	if cur.Left == nil {
		// if the node is the root node
		if !parent.hasVal() {
			bt.Root = cur.Right
		} else if cur.Right == nil {  // leave node
			if isLeftChild {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		} else {  // intermediate node
			if isLeftChild {
				parent.Left = cur.Right
			} else {
				parent.Right = cur.Right
			}
		}
	} else if cur.Right == nil {  // does not have a right child (but must have a left child)
		// if the node is the root node
		if !parent.hasVal() {
			bt.Root = cur.Left
		} else {  // intermediate node
			if isLeftChild {
				parent.Left = cur.Left
			} else {
				parent.Right = cur.Left
			}
		}
	} else {  // have both left child and right child
		parentOfChosen := cur
		chosenLeaveNode := cur.Left
		isChosenLeft := true

		// find the max node in the left BST; use that node to replace the cur node
		for {
			if chosenLeaveNode.Right == nil {
				break
			}
			parentOfChosen = chosenLeaveNode
			chosenLeaveNode = chosenLeaveNode.Right
			isChosenLeft = false
		}

		if isChosenLeft {
			parentOfChosen.Left = chosenLeaveNode.Left
		} else {
			parentOfChosen.Right = chosenLeaveNode.Left
		}

		// if the node to delete is the root node
		if !parent.hasVal() {
			bt.Root = chosenLeaveNode
		} else {
			if isLeftChild {
				parent.Left = chosenLeaveNode
			} else {
				parent.Right = chosenLeaveNode
			}
		}

		// don't forget to change the children of the chosen node
		chosenLeaveNode.Left = cur.Left
		chosenLeaveNode.Right = cur.Right
	}

	return cur.Val, true
}

// Depth returns the depth of the tree (using dfs?)
func (bt *BinarySearchTree) Depth() int {
	return len(bt.Values(false))
}

// Values returns all the values in the tree (in a bfs manner)
//
// fill: if set to true, nil will be used to fill the empty space of non-exist child nodes
func (bt *BinarySearchTree) Values(fill bool) [][]interface{} {
	if bt.Root == nil {  // when the tree is empty
		return make([][]interface{}, 0)
	}

	queue1 := NewQueue() // nodes of the current level
	queue1.Push(bt.Root)
	queue2 := NewQueue() // nodes of the next level
	ans := make([][]interface{}, 0)
	cache := make([]interface{}, 0)

	hasRealValInNextLoop := false

	for {
		if !queue1.HasNext() {
			tmp := make([]interface{}, len(cache))
			copy(tmp, cache)
			ans = append(ans, tmp)
			cache = cache[:0]  // clear the cache

			if !hasRealValInNextLoop {
				break
			}

			if !queue2.HasNext() {
				break
			} else {
				queue1 = queue2.Copy()
				queue2.Reset()
				hasRealValInNextLoop = false
			}
		} else {
			cur := queue1.Pop().(*Node)

			if cur != nil {
				cache = append(cache, cur.Val)
				if cur.Left != nil {
					queue2.Push(cur.Left)
					hasRealValInNextLoop = true
				} else if fill {
					queue2.Push(nil)
				}

				if cur.Right != nil {
					queue2.Push(cur.Right)
					hasRealValInNextLoop = true
				} else if fill {
					queue2.Push(nil)
				}
			} else if fill {
				cache = append(cache, nil)
				queue2.Push(nil)
				queue2.Push(nil)
			}
		}
	}
	return ans
}

// FlattenedValues returns the values in a flattened 1d slice.
//
// fill: if set to true, nil will be used to fill the empty space of non-exist child nodes
func (bt *BinarySearchTree) FlattenedValues(fill bool) []interface{} {
	return common.Flatten2D(bt.Values(fill))
}

// Copy makes a deep copy of the tree
func (bt *BinarySearchTree) Copy() *BinarySearchTree {
	values := bt.FlattenedValues(false)
	newTree := NewBSTree(bt.compare)
	for _, val := range values {
		newTree.Insert(common.Copy(val))
	}
	return newTree
}

// Rebuild returns a tree with the same set of elements that are in different order and the distribution will be more
// condense
func (bt *BinarySearchTree) Rebuild() *BinarySearchTree {
	values := bt.FlattenedValues(false)
	sorter := common.NewSorter(bt.compare)
	sorter.Sort(values)

	newTree := NewBSTree(bt.compare)

	newValues := make([]interface{}, len(values))
	curIndex := 0

	var rearrange func(left, right int)
	rearrange = func(left, right int) {
		if right >= left {
			mid := (left + right) / 2
			newValues[curIndex] = values[mid]
			curIndex += 1
			rearrange(left, mid - 1)
			rearrange(mid + 1, right)
		}
	}

	rearrange(0, len(values) - 1)

	for _, val := range newValues {
		newTree.Insert(val)
	}

	return newTree
}

// NewBSTree returns a new BinarySearchTree
func NewBSTree(compare func(a, b interface{}) int) *BinarySearchTree {
	return &BinarySearchTree{compare: compare}
}

// NewBSTreeInt returns a BinarySearchTree with int val and default compare method
func NewBSTreeInt() *BinarySearchTree {
	compare := func(a, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) == b.(int) {
			return 0
		}
		return -1
	}
	return NewBSTree(compare)
}
