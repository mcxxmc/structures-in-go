package structures

import (
	"fmt"
	"math"
	"some-data-structures/common"
	"strings"
)

// BinarySearchTree the binary search tree
//
// Root *Node
//
// Compare func(a, b interface{}) bool
//
// Compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b
//
// note that this BinarySearchTree does not perform type checking; please include any necessary type checking
// in the customized Compare function
//
// also, the Compare function should never be changed after a tree is created
type BinarySearchTree struct {
	Root *Node
	Compare func(a, b interface{}) int
}

// Insert inserts a new val as a new node
func (bt *BinarySearchTree) Insert(val interface{}) {
	if bt.Root == nil {
		bt.Root = NewNode(val)
		return
	}

	cur := bt.Root
	for {
		if bt.Compare(cur.Val, val) == 1 {  // cur.Val > val
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

// Search returns the corresponding node if the element exists in the tree
func (bt *BinarySearchTree) Search(val interface{}) (*Node, bool) {
	cur := bt.Root
	for {
		if cur == nil {
			break
		}
		c := bt.Compare(cur.Val, val)
		if c == 0 {
			return cur, true
		} else if c == 1 {  // cur.Val > val
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return nil, false
}

// Delete deletes the element if it exists; note that it only deletes 1 element at once
//
// val interface{} should be exactly same with the element that you want to delete
//
// TODO: make a version that takes in a function for comparing e.g., f(a, b interface) bool {return a.(int) == b.(int)}
func (bt *BinarySearchTree) Delete(val interface{}) bool {
	parent := DummyNode()
	isLeftChild := false
	cur := bt.Root

	// first find the node to delete
	for {
		if cur == nil {
			break
		}

		c := bt.Compare(cur.Val, val)

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
		return false
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

	return true
}

// PrintTree prints the tree (using bfs)
//
// TODO: better visualization (include "/" and "\")
func (bt *BinarySearchTree) PrintTree() {
	values := bt.Values(true)
	depth := len(values)
	indentations := getIndentation(depth)
	
	for i, layer := range values {
		s := ""
		s += strings.Repeat(" ", indentations[i][0])
		switcher := false  // TODO: remove this; [i][1] and [1][2] is always same for i != depth - 1
		for _, element := range layer {
			s += common.FixedMinLengthDefault(common.CastString(element))
			if switcher {
				s += strings.Repeat(" ", indentations[i][2])
				switcher = false
			} else {
				s += strings.Repeat(" ", indentations[i][1])
				switcher = true
			}
		}
		fmt.Println(s)
		s = ""
	}
}

// used for calculating indentations for visualizing a bst
func getIndentation(depth int) [][]int {
	if depth < 1 {
		return make([][]int, 0)
	}
	totalDisplayLength := (3 + common.DefaultStringLength) * int(math.Pow(float64(2), float64(depth - 1))) - 4

	ans := make([][]int, depth)
	for i := depth - 1; i >= 0; i -- {
		tmp := make([]int, 3)
		if i == depth - 1 {
			tmp[0] = 0
			tmp[1] = 2
			tmp[2] = 4
		} else if i == 0 {
			tmp[0] = (totalDisplayLength - common.DefaultStringLength) / 2
			tmp[1] = 0
			tmp[2] = 0
		} else {
			tmp[2] = ans[i + 1][1] + ans[i + 1][2] + common.DefaultStringLength
			tmp[1] = tmp[2]
			numOfNodes := int(math.Pow(float64(2), float64(i)))
			tmp[0] = (totalDisplayLength - (numOfNodes * tmp[1] / 2) - ((numOfNodes / 2 - 1) * tmp[2]) -
				(common.DefaultStringLength * numOfNodes)) / 2
		}
		ans[i] = tmp
	}
	return ans
}

// Depth returns the depth of the tree (using dfs?)
func (bt *BinarySearchTree) Depth() int {
	return len(bt.Values(false))
}

// Values returns all the values in the tree (in a bfs manner)
//
// replacement: if set to true, nil will be used to fill the empty space of non-exist child nodes
func (bt *BinarySearchTree) Values(replacement bool) [][]interface{} {
	if bt.Root == nil {  // when the tree is empty
		return make([][]interface{}, 0)
	}

	queue1 := NewQueue() // nodes of the current level
	queue1.Add(bt.Root)
	queue2 := NewQueue() // nodes of the next level
	ans := make([][]interface{}, 0)
	cache := make([]interface{}, 0)

	hasRealValInThisLoop := false

	for {
		if queue1.IsEmpty() {
			if !hasRealValInThisLoop {
				break
			}

			tmp := make([]interface{}, len(cache))
			copy(tmp, cache)
			ans = append(ans, tmp)
			cache = cache[:0]  // clear the cache

			if queue2.IsEmpty() {
				break
			} else {
				queue1 = queue2.Copy()
				queue2.Reset()
				hasRealValInThisLoop = false
			}
		} else {
			cur := queue1.Pop().(*Node)
			if !hasRealValInThisLoop && cur.hasVal() {
				hasRealValInThisLoop = true
			}

			cache = append(cache, cur.Val)

			if cur != nil {
				if cur.Left != nil {
					queue2.Add(cur.Left)
				} else if replacement {
					queue2.Add(DummyNode())
				}

				if cur.Right != nil {
					queue2.Add(cur.Right)
				} else if replacement {
					queue2.Add(DummyNode())
				}
			} else {  // cur will be nil only if replacement == true
				queue2.Add(DummyNode())
				queue2.Add(DummyNode())
			}
		}
	}
	return ans
}

// FlattenedValues returns the values in a flattened 1d slice.
//
// replacement: if set to true, nil will be used to fill the empty space of non-exist child nodes
func (bt *BinarySearchTree) FlattenedValues(replacement bool) []interface{} {
	return common.Flatten2D(bt.Values(replacement))
}

// Copy makes a deep copy of the tree
func (bt *BinarySearchTree) Copy() *BinarySearchTree {
	values := bt.FlattenedValues(false)
	newTree := NewBinaryTree(bt.Compare)
	for _, val := range values {
		newTree.Insert(common.Copy(val))
	}
	return newTree
}

// Rebuild returns a tree with the same set of elements that are in different order and the distribution will be more
// condense
func (bt *BinarySearchTree) Rebuild() *BinarySearchTree {
	values := bt.FlattenedValues(false)
	sorter := common.NewSorter(bt.Compare)
	sorter.Sort(values)

	newTree := NewBinaryTree(bt.Compare)

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

func NewBinaryTree(compare func(a, b interface{}) int) *BinarySearchTree {
	return &BinarySearchTree{Compare: compare}
}

// NewIntBinaryTree returns a BinarySearchTree with int val and default compare method
func NewIntBinaryTree() *BinarySearchTree {
	compare := func(a, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) == b.(int) {
			return 0
		}
		return -1
	}
	return NewBinaryTree(compare)
}
