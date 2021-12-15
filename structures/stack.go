package structures

import (
	"some-data-structures/common"
)

// Stack
//
// The stack structure. Please use NewStack() or NewStackWithValues() as the safe constructor.
type Stack struct {
	stack []interface{}
	top int  // points to the top element of the stack; will be -1 if the stack is empty
}

func (sk *Stack) Len() int {
	return sk.top + 1
}

func (sk *Stack) HasNext() bool {
	return sk.top > -1
}

// Pop pops out the element on the top of the stack.
//
// Please check if the stack is empty before using this method.
//
// e.g.,
//
// if stack.HasNext() { stack.Pop() }
func (sk *Stack) Pop() interface{} {
	if sk.top < 0 {
		return nil
	}
	tmp := sk.stack[sk.top]
	sk.top --
	return tmp
}

// Push pushes v into the stack.
func (sk *Stack) Push(v interface{}) {
	if len(sk.stack) - 1 == sk.top {
		sk.stack = append(sk.stack, make([]interface{}, defaultSize)...)
	}
	sk.top ++
	sk.stack[sk.top] = v
}

// Values returns the values in the stack.
//
// This is NOT a safe method, and you should avoid using it.
func (sk *Stack) Values() []interface{} {
	return sk.stack
}

// Empty completely empties the queue.
//
// Note: it does not empty the queue physically; if you want to release memory, please create a new Stack object instead.
func (sk *Stack) Empty() {
	sk.top = -1
}

// Copy makes a deep copy.
func (sk *Stack) Copy() *Stack {
	tmp := common.CopyInterfaces(sk.stack)
	return &Stack{stack: tmp, top: sk.top}
}

func NewStack() *Stack {
	return &Stack{stack: make([]interface{}, defaultSize), top: -1}
}
