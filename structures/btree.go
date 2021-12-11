package structures

import "some-data-structures/common"

// BTree
//
// The B-tree structure.
//
// Attributes:
//
//		Root *Node
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

// Values returns a deep copy of all the values in the BTree. (in a bfs manner)
func (bt *BTree) Values() []interface{} {
	r := make([]interface{}, bt.num)
	index := 0
	numVals := 2 * bt.t -1
	toVisit := NewQueue()
	toVisit.Push(bt.Root)

	for toVisit.HasNext() {
		cur := toVisit.Pop().(*BTreeNode)
		values := cur.Values
		children := cur.Children

		for i := 0; i < numVals; i ++ {
			tmp := values[i]
			if tmp == nil {
				child := children[i]
				if child != nil {
					toVisit.Push(child)
				}
				break
			} else {
				r[index] = common.Copy(tmp)
				index ++
				child := children[i]
				if child != nil {
					toVisit.Push(child)
				}
				if i == numVals - 1 {
					lastChild := children[i + 1]
					if lastChild != nil {
						toVisit.Push(lastChild)
					}
				}
			}
		}
	}

	return r
}

// Search searches for an element in the BTree.
//
// Returns the corresponding value (it may be different from the input,
// depending on the design of the compare method) and a boolean indicating whether the searching is successful.
func (bt *BTree) Search(val interface{}) (interface{}, bool) {
	cur := bt.Root
	for {
		if cur == nil {
			break
		}
		for i := 0; i < len(cur.Values); i ++ {
			v := cur.Values[i]
			if v == nil {
				cur = cur.Children[i]
			}
			c := bt.compare(v, val)
			if c == 0 {
				return v, true
			} else if c > 0 {
				cur = cur.Children[i]
			} else if i == len(cur.Values) - 1 {
				cur = cur.Children[i + 1]
			}
		}
	}
	return nil, false
}

func (bt *BTree) Insert(val interface{}) bool {
	return false
}

func (bt *BTree) UnsafeInsert(val interface{}) {

}

func (bt *BTree) Delete(interface{}) (interface{}, bool) {
	return nil, false
}

// NewBTree returns a NewBtree object
func NewBTree(t int, compare func(a, b interface{}) int) *BTree {
	return &BTree{Root: NewBTreeNode(t), t: t, compare: compare}
}
