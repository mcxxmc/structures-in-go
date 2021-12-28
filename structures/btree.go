package structures

import (
	"fmt"
)

// BTree
//
// The B-tree structure.
//
// Attributes:
//
//		Root *TreeNode
//
//		t int
//
//		compare func(a, b interface{}) bool
//
// .
//
// t int
//
//		The minimum degree.
//		Every node except root must contain at least t-1 keys. The root may contain minimum 1 key.
//		All nodes (including root) may contain at most 2*t – 1 keys.
//
// .
//
// compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b
//
// .
//
// The first input of the compare function should be the same type as the value of the tree node; the second input may have
// variant types. A tricky compare method can relax the conditions for Search and Delete; see examples for details.
//
// .
//
// Note that this BTree does not perform type checking; please include any necessary type checking
// in the customized compare function
type BTree struct {
	Root *BTreeNode
	t int
	compare func(a, b interface{}) int
	num int  // track number of elements in the tree
}

// T returns the minimum degree
func (bt *BTree) T() int {
	return bt.t
}

// NumOfElements returns the number of elements in the BTree
func (bt *BTree) NumOfElements() int {
	return bt.num
}

// Search searches for an element in the BTree.
//
// Returns the first corresponding node, the index of the corresponding value on the node,
// and a boolean indicating whether the searching is successful.
func (bt *BTree) Search(val interface{})  (*BTreeNode, int, bool) {
	cur := bt.Root

	for cur != nil {
		i := 0
		for i < cur.N && bt.compare(cur.Keys[i], val) == -1 {
			i ++
		}

		if i < cur.N && bt.compare(cur.Keys[i], val) == 0 {
			return cur, i, true
		} else if cur.IsLeaf {
			return nil, -1, false
		} else {
			cur = cur.Children[i]  // here, "i" can be equal to cur.n, which means the last child
		}
	}

	return nil, -1, false
}

func (bt *BTree) Predecessor(node *BTreeNode) interface{} {
	cur := node
	for !cur.IsLeaf {
		cur = cur.Children[cur.N]
	}
	return cur.Keys[cur.N - 1]
}

func (bt *BTree) Successor(node *BTreeNode) interface{} {
	cur := node
	for !cur.IsLeaf {
		cur = cur.Children[0]
	}
	return cur.Keys[0]
}

// finds the proper index for the given key in the current node.Keys.
func (bt *BTree) findKey(node *BTreeNode, val interface{}) int {
	r := 0
	for r < node.N && bt.compare(node.Keys[r], val) == -1 {
		r ++
	}
	return r
}

// splits the index th child of the node; index starts from 0.
func (bt *BTree) splitChild(node *BTreeNode, index int) {
	t := bt.t
	y := node.Children[index]
	z := NewBTreeNode(t, y.IsLeaf)
	z.N = t - 1

	for i := 0; i < t - 1; i ++ {
		z.Keys[i] = y.Keys[i + t]
	}
	if !z.IsLeaf {
		for i := 0; i < t; i ++ {
			z.Children[i] = y.Children[i + t]
		}
	}

	y.N = t - 1  // the mid-key (t th key) is popped up so the total number of keys left is 2 * t -2

	for i := node.N; i >= index + 1; i -- {  // note that we assume node itself is not full
		node.Children[i + 1] = node.Children[i]
	}
	node.Children[index + 1] = z
	z.Parent = node

	for i := node.N - 1; i >= index; i -- {
		node.Keys[i + 1] = node.Keys[i]
	}
	node.Keys[index] = y.Keys[t - 1]
	node.N ++
}

// inserts val into the node; the node must not be full
func (bt *BTree) insertNotFull(node *BTreeNode, val interface{}) {
	i := node.N - 1
	if node.IsLeaf {
		for i >= 0 && bt.compare(node.Keys[i], val) == 1 {
			node.Keys[i + 1] = node.Keys[i]
			i --
		}
		node.Keys[i + 1] = val
		node.N ++
	} else {
		for i >= 0 && bt.compare(node.Keys[i], val) == 1 {
			i --
		}
		i ++  // to get the child after i th key
		if node.Children[i].N == 2 * bt.t - 1 {
			bt.splitChild(node, i)
			if bt.compare(node.Keys[i], val) == -1 {  // the popped up mid-key may be smaller than val
				i ++
			}
		}
		bt.insertNotFull(node.Children[i], val)
	}
}

// Insert inserts a new value into the b-tree.
func (bt *BTree) Insert(val interface{}) {
	cur := bt.Root
	if cur.N == 2 * bt.t - 1 {
		s := NewBTreeNode(bt.t, false)
		bt.Root = s
		s.Children[0] = cur
		cur.Parent = s
		bt.splitChild(s, 0)  // this will also set keys for s
		bt.insertNotFull(s, val)
	} else {
		bt.insertNotFull(cur, val)
	}

	bt.num ++
}

// removes the key k from the sub-tree rooted with this node.
func (bt *BTree) removeFromNode(node *BTreeNode, val interface{}) bool {
	idx := bt.findKey(node, val)
	if idx < node.N && bt.compare(node.Keys[idx], val) == 0 {
		if node.IsLeaf {
			return bt.removeFromLeaf(node, idx)
		} else {
			return bt.removeFromNonLeaf(node, idx)
		}
	} else {
		if node.IsLeaf {  // not found
			return false
		}

		flag := idx == node.N

		if node.Children[idx].N < bt.t {
			bt.fill(node, idx)
		}

		if flag && idx > node.N {
			return bt.removeFromNode(node.Children[idx - 1], val)
		} else {
			return bt.removeFromNode(node.Children[idx], val)
		}
	}
}

// removes a key from a leaf node.
func (bt *BTree) removeFromLeaf(node *BTreeNode, index int) bool {
	for i := index + 1; i < node.N; i ++ {
		node.Keys[i - 1] = node.Keys[i]
	}
	node.N --
	return true
}

// removes a key from a non-leaf node
func (bt *BTree) removeFromNonLeaf(node *BTreeNode, index int) bool {
	k := node.Keys[index]

	if node.Children[index].N >= bt.t {
		predecessor := bt.Predecessor(node.Children[index])
		node.Keys[index] = predecessor
		return bt.removeFromNode(node.Children[index], predecessor)
	} else if node.Children[index + 1].N >= bt.t {
		successor := bt.Successor(node.Children[index + 1])
		node.Keys[index] = successor
		return bt.removeFromNode(node.Children[index + 1], successor)
	} else {
		bt.combineChildren(node, index, index + 1)
		return bt.removeFromNode(node.Children[index], k)
	}
}

// combines 2 neighboring children of the node.
func (bt *BTree) combineChildren(node *BTreeNode, index1 int, index2 int) {
	t := bt.t
	y1 := node.Children[index1]
	y2 := node.Children[index2]

	y1.Keys[t - 1] = node.Keys[index1]  // add "k" to y1.Keys

	for i := 0; i < y2.N; i ++ {  // copying keys
		y1.Keys[i + t] = y2.Keys[i]
	}

	if !y1.IsLeaf {
		for i := 0; i < y2.N; i ++ {  // copying children
			y1.Children[i + t] = y2.Children[i]
		}
	}

	for i := index2; i < node.N; i ++ {  // move keys
		node.Keys[i - 1] = node.Keys[i]  // node.Keys[index] is deleted here
	}

	for i := index2 + 1; i < node.N + 1; i ++ {  // move children
		node.Children[i - 1] = node.Children[i]  // node.Children[index2] is deleted here
	}

	y1.N += y2.N + 1
	node.N --
	return
}

// borrows a key from the previous / left sibling.
func (bt *BTree) borrowPrev(node *BTreeNode, index int) {
	cur := node.Children[index]
	toBorrow := node.Children[index - 1]

	// the borrowed key will be smaller; so move all keys in cur on step to the right
	for i := cur.N - 1; i >= 0; i -- {
		cur.Keys[i + 1] = cur.Keys[i]
	}

	// also, moves all children if necessary
	if !cur.IsLeaf {
		for i := cur.N; i >= 0; i -- {
			cur.Children[i + 1] = cur.Children[i]
		}
	}

	cur.Keys[0] = node.Keys[index - 1]  // use the previous key in the node as the first key of the child

	if !cur.IsLeaf {  // do not forget the child of the sibling
		cur.Children[0] = toBorrow.Children[toBorrow.N]
		cur.Children[0].Parent = cur
	}

	node.Keys[index - 1] = toBorrow.Keys[toBorrow.N - 1]  // use the last key in toBorrow as the new key
	cur.N ++
	toBorrow.N --
}

// borrows a key from the next / right sibling.
func (bt *BTree) borrowNext(node *BTreeNode, index int) {
	cur := node.Children[index]
	toBorrow := node.Children[index + 1]

	// the borrowed key will be bigger
	cur.Keys[cur.N] = node.Keys[index]
	if ! cur.IsLeaf {
		cur.Children[cur.N + 1] = toBorrow.Children[0]
		cur.Children[cur.N + 1].Parent = cur
	}

	node.Keys[index] = toBorrow.Keys[0]

	// updates toBorrow
	for i := 0; i < toBorrow.N - 1; i ++ {
		toBorrow.Keys[i] = toBorrow.Keys[i + 1]
	}

	if !toBorrow.IsLeaf {
		for i := 0; i < toBorrow.N; i ++ {
			toBorrow.Children[i] = toBorrow.Children[i + 1]
		}
	}

	cur.N ++
	toBorrow.N --
}

// fills the child node which has less than t - 1 keys.
func (bt *BTree) fill(node *BTreeNode, index int) {
	// the main idea is to decide where to borrow
	t := bt.t
	if index != 0 && node.Children[index - 1].N >= t {
		bt.borrowPrev(node, index)
	} else if index != node.N && node.Children[index + 1].N >= t {
		bt.borrowNext(node, index)
	} else {  // none of the siblings have extra keys
		if index != node.N {
			bt.combineChildren(node, index, index + 1)
		} else {
			bt.combineChildren(node, index - 1, index)
		}
	}
}

func (bt *BTree) Delete(val interface{}) bool {
	b := bt.removeFromNode(bt.Root, val)
	if b {
		bt.num --
	}

	// check if the root has no keys && no children
	if bt.Root.N == 0 {
		if !bt.Root.IsLeaf {
			bt.Root = bt.Root.Children[0]
			bt.Root.Parent = nil
		}
	}

	return b
}

// Values returns all the values in the tree in an ordered manner.
//
// Note: it uses (modified) dfs and is expensive.
func (bt *BTree) Values() []interface{} {
	r := make([]interface{}, bt.num)
	index := 0
	var dfs func(node *BTreeNode)
	dfs = func(node *BTreeNode) {
		if node == nil {
			return
		}
		for i := 0; i < node.N; i ++ {
			if !node.IsLeaf {
				dfs(node.Children[i])
			}
			r[index] = node.Keys[i]
			index ++
		}
		dfs(node.Children[node.N])
	}
	dfs(bt.Root)
	return r
}

// NewBTree returns a NewBtree object
//
// t must > 1; otherwise it will return nil.
func NewBTree(t int, compare func(a, b interface{}) int) *BTree {
	if t < 2 {
		fmt.Println("the minimum degree t must be > 1")
		return nil
	}
	return &BTree{Root: NewBTreeNode(t, true), t: t, compare: compare}
}
