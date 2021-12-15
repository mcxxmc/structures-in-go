package structures

// LinkedList
//
// The basic linked list structure. Please use NewLinkedList() as the safe constructor.
//
// Attributes:
//
// Head *Node: the head of the linked list. Usually it is a dummy head.
//
// compare func(a, b interface{}) int: the compare method.
type LinkedList struct {
	Head *Node
	compare func(a, b interface{}) int
}

// Search returns the first element satisfying the search condition.
//
// Will return nil if there is no such element.
func (ll *LinkedList) Search(v interface{}) *Node {
	pt := ll.Head
	for pt.Next != nil {
		pt = pt.Next
		if ll.compare(pt.Val, v) == 0 {
			return pt
		}
	}
	return nil
}

// Insert inserts v as a new node to the head of this linked list.
func (ll *LinkedList) Insert(v interface{}) {
	node := NewNode(v)
	node.Next = ll.Head.Next
	ll.Head.Next = node
}

// Delete deletes and returns the first element satisfying the search condition.
func (ll *LinkedList) Delete(v interface{}) *Node {
	prev := ll.Head
	cur := prev.Next
	if cur != nil {
		if ll.compare(cur.Val, v) == 0 {
			prev.Next = cur.Next
			cur.Next = nil
			return cur
		}
		prev = cur
		cur = cur.Next
	}
	return nil
}

// NewLinkedList returns a LinkedList object.
//
// compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b;
// a should always be an element from the struct other than user input.
func NewLinkedList(compare func(a, b interface{}) int) *LinkedList {
	return &LinkedList{Head: DummyNode(), compare: compare}
}

// DoubleLinkedList
//
// The double linked list with a sentinel. Please use NewDoubleLinkedList() as a safe constructor.
//
// Attributes:
//
// Head *BiNode: the head of the linked list. Usually it is a dummy head.
//
// compare func(a, b interface{}) int: the compare method.
type DoubleLinkedList struct {
	Head *BiNode
	compare func(a, b interface{}) int
}

// Search returns the first element satisfying the search condition.
//
// Will return nil if there is no such element.
func (dll *DoubleLinkedList) Search(v interface{}) *BiNode {
	pt := dll.Head
	for pt.Next.Val != nil {
		pt = pt.Next
		if dll.compare(pt.Val, v) == 0 {
			return pt
		}
	}
	return nil
}

// Insert inserts v as a new node to the head of this linked list.
func (dll *DoubleLinkedList) Insert(v interface{}) {
	node := NewBiNode(v)
	node.Next = dll.Head.Next
	dll.Head.Next.Prev = node
	dll.Head.Next = node
	node.Prev = dll.Head
}

// Delete deletes and returns the first element satisfying the search condition.
func (dll *DoubleLinkedList) Delete(v interface{}) *BiNode {
	node := dll.Search(v)
	dll.DeleteNode(node)
	return node
}

// DeleteNode deletes the node from the double linked list.
func (dll *DoubleLinkedList) DeleteNode(n *BiNode) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

func NewDoubleLinkedList(compare func(a, b interface{}) int) *DoubleLinkedList {
	node := DummyBiNode()
	node.Prev = node
	node.Next = node
	return &DoubleLinkedList{Head: node, compare: compare}
}
