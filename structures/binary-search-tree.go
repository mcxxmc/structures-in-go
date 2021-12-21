package structures

// BinarySearchTree
//
// The binary search tree. Please use NewBSTree() as the safe constructor.
//
// Attributes:
//
//	Root *TreeNode
//
//	compare func(a, b interface{}) bool
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
//
// .
//
// Note that this BinarySearchTree does not perform type checking; please include any necessary type checking
// in the customized compare function.
type BinarySearchTree struct {
	Root    *TreeNode
	compare func(a, b interface{}) int
}

// InOrderTreeWalk returns all the values of the tree in an in-order-tree-walk manner.
func (bt *BinarySearchTree) InOrderTreeWalk() []interface{} {
	r := make([]interface{}, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			r = append(r, node.Val)
			inorder(node.Right)
		}
	}
	inorder(bt.Root)
	return r
}

// Search returns the pointer to the FIRST corresponding TreeNode if that TreeNode exists in the tree.
func (bt *BinarySearchTree) Search(val interface{}) (*TreeNode, bool) {
	cur := bt.Root
	for {
		if cur == nil {
			break
		}
		c := bt.compare(cur.Val, val)
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

// MaxSince returns the pointer to the max (rightmost) TreeNode in the subtree since the current node.
func (bt *BinarySearchTree) MaxSince(node *TreeNode) *TreeNode {
	cur := node
	for {
		if cur == nil || cur.Right == nil {
			break
		}
		cur = cur.Right
	}
	return cur
}

// MinSince returns the pointer to the min (leftmost) TreeNode in the subtree since the current node.
func (bt *BinarySearchTree) MinSince(node *TreeNode) *TreeNode {
	cur := node
	for {
		if cur == nil || cur.Left == nil {
			break
		}
		cur = cur.Left
	}
	return cur
}

// Max returns the pointer to the max (rightmost) TreeNode in the tree.
func (bt *BinarySearchTree) Max() *TreeNode {
	return bt.MaxSince(bt.Root)
}

// Min returns the pointer to the min (leftmost) TreeNode in the tree.
func (bt *BinarySearchTree) Min() *TreeNode {
	return bt.MinSince(bt.Root)
}

// Successor find the minimum tree node that is bigger than (to the right of) the current node.
//
// It will return nil if the current node is nil.
func (bt *BinarySearchTree) Successor(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return bt.MinSince(node.Right)
	}
	y := node.Parent
	x := node
	for y != nil && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}

// Predecessor find the maximum tree node that is smaller than (to the left of) the current node.
//
// It will return nil if the current node is nil.
func (bt *BinarySearchTree) Predecessor(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		return bt.MaxSince(node.Left)
	}
	y := node.Parent
	x := node
	for y != nil && x == y.Left {
		x = y
		y = y.Parent
	}
	return y
}

func (bt *BinarySearchTree) insert(val interface{}, safe bool) bool {
	node := NewTreeNode(val)
	if bt.Root == nil {
		bt.Root = node
		return true
	}

	cur := bt.Root
	for {
		c := bt.compare(cur.Val, val)
		if c == 1 { // cur.Val > val
			if cur.Left == nil {
				cur.Left = node
				node.Parent = cur
				break
			} else {
				cur = cur.Left
			}
		} else {
			if c == 0 && safe {
				return false
			}
			if cur.Right == nil {
				cur.Right = node
				node.Parent = cur
				break
			} else {
				cur = cur.Right
			}
		}
	}
	return true
}

// Insert inserts a new val as a new node.
//
// Does not insert if the val already exists in the tree.
func (bt *BinarySearchTree) Insert(val interface{}) bool {
	return bt.insert(val, true)
}

// UnsafeInsert inserts a new val as a new node and allows the same val to be inserted for multiple times.
func (bt *BinarySearchTree) UnsafeInsert(val interface{}) {
	bt.insert(val, false)
}

// uses subtree n2 to replace subtree n1 and updates the parent information (and this is exactly the only thing it does).
func (bt *BinarySearchTree) transplant(n1, n2 *TreeNode) {
	if n1.Parent == nil {
		bt.Root = n2
	} else if n1 == n1.Parent.Left {
		n1.Parent.Left = n2
	} else {
		n1.Parent.Right = n2
	}
	if n2 != nil {
		n2.Parent = n1.Parent
	}
}

// DeleteNode deletes the node from the tree.
//
// it returns a boolean value indicating if the deletion is successful.
func (bt *BinarySearchTree) DeleteNode(node *TreeNode) bool {
	if node == nil {
		return false
	}

	if node.Left == nil {  // if the node does not have a left child (it may also do not have a right child)
		bt.transplant(node, node.Right)
	} else if node.Right == nil {  // does not have a right child (but must have a left child)
		bt.transplant(node, node.Left)
	} else {  // has both right child and left child
		y := bt.MaxSince(node.Left)  // the max node that is smaller than the current node
		if y.Parent != node {
			bt.transplant(y, y.Left)
			y.Left = node.Left
			y.Left.Parent = y
		}
		bt.transplant(node, y)
		y.Right = node.Right
		y.Right.Parent = y
	}

	return true
}

// Delete deletes the First node with the corresponding value if it exists.
//
// it returns a boolean value indicating if the deletion is successful.
func (bt *BinarySearchTree) Delete(v interface{}) bool {
	node, _ := bt.Search(v)
	return bt.DeleteNode(node)
}

// Height returns the height of the tree.
//
// Warning: it uses dfs and is expensive.
func (bt *BinarySearchTree) Height() int {
	max := 0
	count := 0
	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur != nil {
			count ++
			if count > max {
				max = count
			}
			dfs(cur.Left)
			dfs(cur.Right)
			count --
		}
	}
	dfs(bt.Root)
	return max
}

// Rebuild returns a tree with the same set of elements that are in different order and the distribution will be more
// condense.
func (bt *BinarySearchTree) Rebuild() *BinarySearchTree {
	values := bt.InOrderTreeWalk()

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

// NewIntBSTree returns a BinarySearchTree with int val and default compare method
func NewIntBSTree() *BinarySearchTree {
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
