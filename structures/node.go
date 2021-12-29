package structures

// TreeNode
//
// The basic tree node.
//
// Attributes:
//
// Val interface{}: the value.
//
// Left *TreeNode: the smaller (or equal) left child.
//
// Right *TreeNode: the bigger right child.
//
// Parent *TreeNode: the parent node.
type TreeNode struct {
	Val interface{}
	Left *TreeNode
	Right *TreeNode
	Parent *TreeNode
}

func NewTreeNode(val interface{}) *TreeNode {
	return &TreeNode{Val: val}
}

// RBTreeNode
//
// The tree node for red-black-tree.
//
// Attributes:
//
// Val interface{}: the value.
//
// Color bool: if the node is red.
//
// Left *RBTreeNode: the smaller (or equal) left child.
//
// Right *RBTreeNode: the bigger right child.
//
// Parent *RBTreeNode: the parent node.
type RBTreeNode struct {
	Val   interface{}
	Color bool
	Left  *RBTreeNode
	Right *RBTreeNode
	Parent *RBTreeNode
}

func NewRBTreeNode(val interface{}, isRed bool) *RBTreeNode {
	return &RBTreeNode{Val: val, Color: isRed}
}

// BTreeNode The node used as the internal node & the leaf node for B tree
//
// Number of children of a node is equal to the number of keys in it plus 1.
//
// N int: the number of keys.
type BTreeNode struct {
	IsLeaf bool
	N int
	Keys []interface{}
	Children []*BTreeNode
	Parent *BTreeNode
}

// NewBTreeNode returns a new BTreeNode.
//
// t int
//
//		The minimum degree.
//		Every node except root must contain at least t-1 keys. The root may contain minimum 1 key.
//		All nodes (including root) may contain at most 2*t – 1 keys.
//		t must > 1.
//
// Number of children of a node is equal to the number of keys in it plus 1.
func NewBTreeNode(t int, isLeaf bool) *BTreeNode {
	return &BTreeNode{IsLeaf: isLeaf, N : 0, Keys: make([]interface{}, 2 * t - 1), Children: make([]*BTreeNode, 2 * t)}
}

// Node
//
// The basic node.
//
// Attributes:
//
// Val interface{}
//
// Next *Node  the child node
type Node struct {
	Val interface{}
	Next *Node
}

// DummyNode returns a dummy node with val as nil
func DummyNode() *Node {
	return &Node{Val: nil}
}

func NewNode(val interface{}) *Node {
	return &Node{Val: val}
}

// BiNode
//
// The bi-direction node.
//
// Attributes:
//
// Val interface{}
//
// Prev *BiNode the parent node
//
// Next *BiNode  the child node
type BiNode struct {
	Val interface{}
	Prev *BiNode
	Next *BiNode
}

// DummyBiNode returns a dummy node with val as nil
func DummyBiNode() *BiNode {
	return &BiNode{Val: nil}
}

func NewBiNode(val interface{}) *BiNode {
	return &BiNode{Val: val}
}

// FibNode
//
// The node structure for fibonacci heap.
//
// Attributes:
//
// Marked bool: if the node has lost a child node.
//
// Degree int: the number of the children of this node.
//
// Val interface{}
//
// Parent *FibNode: the parent.
//
// Child *FibNode: the child.
//
// Left *FibNode: the left sibling.
//
// Right *FibNode: the right sibling.
type FibNode struct {
	Marked bool  // false by default
	Degree int  // 0 by default
	Val interface{}
	Parent *FibNode
	Child *FibNode
	Left *FibNode
	Right *FibNode
}

// NewFibNode creates a new Fibonacci node object.
func NewFibNode(val interface{}) *FibNode {
	return &FibNode{Val: val}
}