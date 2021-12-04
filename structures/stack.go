package structures

import "some-data-structures/common"

type Stack struct {
	stack []interface{}
}

func (st *Stack) Len() int {
	return len(st.stack)
}

func (st *Stack) HasNext() bool {
	return len(st.stack) != 0
}

// Pop pops out the element on the top of the stack
//
// e.g.,
//
// if stack.HasNext() {
//
//     stack.Pop()
//
// }
func (st *Stack) Pop() interface{} {
	v := st.stack[0]
	st.stack = st.stack[1:]
	return v
}

func (st *Stack) Push(v interface{}) {
	st.stack = append([]interface{}{v}, st.stack...)
}

// Reset completely resets the queue
//
// Warning: it will empty the queue
func (st *Stack) Reset() {
	st.stack = make([]interface{}, 0)
}

// Copy makes a deep copy
func (st *Stack) Copy() *Stack {
	cpy := &Stack{stack: make([]interface{}, len(st.stack))}
	for i := 0; i < len(st.stack); i ++ {
		cpy.stack[i] = common.Copy(st.stack[i])
	}
	return cpy
}

func NewStack() *Stack {
	return &Stack{stack: make([]interface{}, 0)}
}
