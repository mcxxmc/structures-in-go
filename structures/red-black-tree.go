package structures

const red = true
const black = false

// RedBlackTree
//
// The red-black tree structure. Please use NewRedBlackTree() as the safe constructor.
//
// Attributes:
//
// Root *RBTreeNode
//
// sentinel *RBTreeNode
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
//
// .
//
// Note that this RedBlackTree does not perform type checking; please include any necessary type checking
// in the customized compare function.
type RedBlackTree struct {
	Root *RBTreeNode
	sentinel *RBTreeNode
	compare  func(a, b interface{}) int
}

// InOrderTreeWalk returns all the values of the tree in an in-order-tree-walk manner.
func (rbt *RedBlackTree) InOrderTreeWalk() []interface{} {
	r := make([]interface{}, 0)
	var inorder func(node *RBTreeNode)
	inorder = func(node *RBTreeNode) {
		if node != nil && node != rbt.sentinel {
			inorder(node.Left)
			r = append(r, node.Val)
			inorder(node.Right)
		}
	}
	inorder(rbt.Root)
	return r
}

// Search returns the pointer to the FIRST corresponding RBTreeNode if that RBTreeNode exists in the tree.
func (rbt *RedBlackTree) Search(val interface{}) (*RBTreeNode, bool) {
	cur := rbt.Root
	for {
		if cur == nil || cur == rbt.sentinel {
			break
		}
		c := rbt.compare(cur.Val, val)
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

// MaxSince returns the pointer to the max (rightmost) RBTreeNode in the subtree since the current node.
func (rbt *RedBlackTree) MaxSince(node *RBTreeNode) *RBTreeNode {
	cur := node
	for {
		if cur == nil || cur == rbt.sentinel || cur.Right == rbt.sentinel {
			break
		}
		cur = cur.Right
	}
	return cur
}

// MinSince returns the pointer to the min (leftmost) RBTreeNode in the subtree since the current node.
func (rbt *RedBlackTree) MinSince(node *RBTreeNode) *RBTreeNode {
	cur := node
	for {
		if cur == nil || cur == rbt.sentinel || cur.Left == rbt.sentinel {
			break
		}
		cur = cur.Left
	}
	return cur
}

// Max returns the pointer to the max (rightmost) RBTreeNode in the tree.
func (rbt *RedBlackTree) Max() *RBTreeNode {
	return rbt.MaxSince(rbt.Root)
}

// Min returns the pointer to the min (leftmost) RBTreeNode in the tree.
func (rbt *RedBlackTree) Min() *RBTreeNode {
	return rbt.MinSince(rbt.Root)
}

// Successor find the minimum tree node that is bigger than (to the right of) the current node.
//
// It will return nil if the current node is nil.
func (rbt *RedBlackTree) Successor(node *RBTreeNode) *RBTreeNode {
	if node == nil || node == rbt.sentinel {
		return nil
	}
	if node.Right != rbt.sentinel {
		return rbt.MinSince(node.Right)
	}
	y := node.Parent
	x := node
	for y != rbt.sentinel && x == y.Right {
		x = y
		y = y.Parent
	}
	if y == rbt.sentinel {
		return nil
	}
	return y
}

// Predecessor find the maximum tree node that is smaller than (to the left of) the current node.
//
// It will return nil if the current node is nil.
func (rbt *RedBlackTree) Predecessor(node *RBTreeNode) *RBTreeNode {
	if node == nil || node == rbt.sentinel {
		return nil
	}
	if node.Left != rbt.sentinel {
		return rbt.MaxSince(node.Left)
	}
	y := node.Parent
	x := node
	for y != rbt.sentinel && x == y.Left {
		x = y
		y = y.Parent
	}
	if y == rbt.sentinel {
		return nil
	}
	return y
}

// left-rotates the subtree for balance
func (rbt *RedBlackTree) leftRotate(node *RBTreeNode) {
	if node.Right == rbt.sentinel {
		return
	}
	y := node.Right
	node.Right = y.Left
	if y.Left != rbt.sentinel {
		y.Left.Parent = node
	}
	y.Parent = node.Parent
	if node.Parent == rbt.sentinel {
		rbt.Root = y
	} else if node == node.Parent.Left {
		node.Parent.Left = y
	} else {
		node.Parent.Right = y
	}
	y.Left = node
	node.Parent = y
}

// right-rotates the subtree for balance
func (rbt *RedBlackTree) rightRotate(node *RBTreeNode) {
	if node.Left == rbt.sentinel {
		return
	}
	x := node.Left
	node.Left = x.Right
	if x.Right != rbt.sentinel {
		x.Right.Parent = node
	}
	x.Parent = node.Parent
	if node.Parent == rbt.sentinel {
		rbt.Root = x
	} else if node == node.Parent.Left {
		node.Parent.Left = x
	} else {
		node.Parent.Right = x
	}
	x.Right = node
	node.Parent = x
}

func (rbt *RedBlackTree) insert(val interface{}, safe bool) bool {
	node := NewRBTreeNode(val, red)
	prev := rbt.sentinel
	cur := rbt.Root
	for cur != nil && cur != rbt.sentinel {
		prev = cur
		c := rbt.compare(cur.Val, val)
		if c == 1 {  // cur.Val > val
			cur = cur.Left
		} else {
			if c == 0 && safe {
				return false
			}
			cur = cur.Right
		}
	}

	node.Parent = prev
	if prev == rbt.sentinel {
		rbt.Root = node
	} else if rbt.compare(prev.Val, val) == 1 {
		prev.Left = node
	} else {
		prev.Right = node
	}

	node.Left = rbt.sentinel
	node.Right = rbt.sentinel
	rbt.insertFixup(node)
	return true
}

// restores the red-black tree property
func (rbt *RedBlackTree) insertFixup(node *RBTreeNode) {
	for node.Parent.Color == red {
		if node.Parent == node.Parent.Parent.Left {
			y := node.Parent.Parent.Right

			if y.Color == red {  // case 1
				node.Parent.Color = black
				y.Color = black
				node.Parent.Parent.Color = red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {  // case 2
					node = node.Parent
					rbt.leftRotate(node)
				}
				node.Parent.Color = black  // case 3
				node.Parent.Parent.Color = red
				rbt.rightRotate(node.Parent.Parent)
			}
		} else if node.Parent == node.Parent.Parent.Right {
			y := node.Parent.Parent.Left

			if y.Color == red {
				node.Parent.Color = black
				y.Color = black
				node.Parent.Parent.Color = red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					rbt.rightRotate(node)
				}
				node.Parent.Color = black
				node.Parent.Parent.Color = red
				rbt.leftRotate(node.Parent.Parent)
			}
		}
	}
	rbt.Root.Color = black
}

// Insert inserts a new val as a new node.
//
// Does not insert if the val already exists in the tree.
func (rbt *RedBlackTree) Insert(val interface{}) bool {
	return rbt.insert(val, true)
}

// UnsafeInsert inserts a new val as a new node and allows the same val to be inserted for multiple times.
func (rbt *RedBlackTree) UnsafeInsert(val interface{}) {
	rbt.insert(val, false)
}

// uses subtree n2 to replace subtree n1 by connecting n2 and the parent of n1.
//
// It does not update the child of n1 or n2.
func (rbt *RedBlackTree) transplant(n1, n2 *RBTreeNode) {
	if n1.Parent == rbt.sentinel {
		rbt.Root = n2
	} else if n1 == n1.Parent.Left {
		n1.Parent.Left = n2
	} else {
		n1.Parent.Right = n2
	}
	n2.Parent = n1.Parent
}

// restores the red-black tree property.
func (rbt *RedBlackTree) deleteFixup(node *RBTreeNode) {
	for node != rbt.Root && node != rbt.sentinel && node.Color == black {
		if node == node.Parent.Left {
			w := node.Parent.Right
			if w.Color == red {  // case 1
				w.Color = black
				node.Parent.Color = red
				rbt.leftRotate(node.Parent)
				w = node.Parent.Right
			}
			if w.Left.Color == black && w.Right.Color == black {  // case 2
				w.Color = red
				node = node.Parent
			} else {
				if w.Right.Color == black {  // case 3
					w.Left.Color = black
					w.Color = red
					rbt.rightRotate(w)
					w = node.Parent.Right
				}
				// case 4
				w.Color = node.Parent.Color
				node.Parent.Color = black
				w.Right.Color = black
				rbt.leftRotate(node.Parent)
				node = rbt.Root
			}
		} else if node == node.Parent.Right {
			w := node.Parent.Left
			if w.Color == red {
				w.Color = black
				node.Parent.Color = red
				rbt.rightRotate(node.Parent)
				w = node.Parent.Left
			}
			if w.Right.Color == black && w.Left.Color == black {
				w.Color = red
				node = node.Parent
			} else {
				if w.Left.Color == black {
					w.Right.Color = black
					w.Color = red
					rbt.leftRotate(w)
					w = node.Parent.Left
				}
				w.Color = node.Parent.Color
				node.Parent.Color = black
				w.Left.Color = black
				rbt.rightRotate(node.Parent)
				node = rbt.Root
			}
		}
	}
	node.Color = black
}

// DeleteNode deletes the node from the tree.
//
// it returns a boolean value indicating if the deletion is successful.
func (rbt *RedBlackTree) DeleteNode(node *RBTreeNode) bool {
	if node == nil {
		return false
	}

	y := node
	yOriginalColor := y.Color
	var x *RBTreeNode

	if node.Left == rbt.sentinel {
		x = node.Right
		rbt.transplant(node, node.Right)
	} else if node.Right == rbt.sentinel {
		x = node.Left
		rbt.transplant(node, node.Left)
	} else {
		y = rbt.MinSince(node.Right)  // the successor of node
		yOriginalColor = y.Color
		x = y.Right
		if y.Parent != node {
			rbt.transplant(y, y.Right)
			y.Right = node.Right
			y.Right.Parent = y
		} else {
			x.Parent = y
		}
		rbt.transplant(node, y)
		y.Left = node.Left
		y.Left.Parent = y
		y.Color = node.Color
	}

	if yOriginalColor == black {
		rbt.deleteFixup(x)
	}
	return true
}

// Delete deletes the First node with the corresponding value if it exists.
//
// it returns a boolean value indicating if the deletion is successful.
func (rbt *RedBlackTree) Delete(val interface{}) bool {
	node, _ := rbt.Search(val)
	return rbt.DeleteNode(node)
}

// Height returns the height of the tree.
//
// Warning: it uses dfs and is expensive.
func (rbt *RedBlackTree) Height() int {
	max := 0
	count := 0
	var dfs func(cur *RBTreeNode)
	dfs = func(cur *RBTreeNode) {
		if cur != rbt.sentinel {
			count ++
			if count > max {
				max = count
			}
			dfs(cur.Left)
			dfs(cur.Right)
			count --
		}
	}
	dfs(rbt.Root)
	return max
}

// NewRedBlackTree returns a new RedBlackTree object.
func NewRedBlackTree(compare func(a, b interface{}) int) *RedBlackTree {
	sentinel := NewRBTreeNode(nil, black)
	return &RedBlackTree{compare: compare, sentinel: sentinel}
}
