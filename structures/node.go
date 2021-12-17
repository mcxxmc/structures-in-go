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

// returns false if the Val is nil
func (node *TreeNode) hasVal() bool {
	return node.Val != nil
}

// DummyTreeNode returns a dummy tree node with val as nil
func DummyTreeNode() *TreeNode {
	return &TreeNode{Val: nil}
}

func NewTreeNode(val interface{}) *TreeNode {
	return &TreeNode{Val: val}
}

// BTreeNode The node used as the internal node & the leaf node for B tree
//
// Number of children of a node is equal to the number of keys in it plus 1.
type BTreeNode struct {
	Values []interface{}
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
//
// Number of children of a node is equal to the number of keys in it plus 1.
func NewBTreeNode(t int) *BTreeNode {
	return &BTreeNode{Values: make([]interface{}, 2 * t - 1), Children: make([]*BTreeNode, 2 * t)}
}

// Node
//
// The basic node.
//
// Attributes:
//
// Val interface{}
//
// Next *BiNode  the child node
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