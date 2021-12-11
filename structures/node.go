package structures

// Node the node
//
// Val interface{}
//
// Left *Node  the smaller (or equal) left child
//
// Right *Node  the bigger right child
type Node struct {
	Val interface{}
	Left *Node
	Right *Node
}

// returns false if the Val is nil
func (node *Node) hasVal() bool {
	return node.Val != nil
}

// DummyNode returns a dummy node with val as nil
func DummyNode() *Node {
	return &Node{Val: nil}
}

func NewNode(val interface{}) *Node {
	return &Node{Val: val}
}

// BiNode the node with 2 direction
//
// Val interface{}
//
// Parent *BiNode the parent node
//
// Left *BiNode  the smaller (or equal) left child
//
// Right *BiNode  the bigger right child
type BiNode struct {
	Val interface{}
	Parent *BiNode
	Left *BiNode
	Right *BiNode
}

// returns false if the Val is nil
func (node *BiNode) hasVal() bool {
	return node.Val != nil
}

// DummyBiNode returns a dummy node with val as nil
func DummyBiNode() *Node {
	return &Node{Val: nil}
}

func NewBiNode(val interface{}) *BiNode {
	return &BiNode{Val: val}
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
