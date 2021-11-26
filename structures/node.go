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