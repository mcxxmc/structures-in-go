package structures

import (
	"some-data-structures/common"
)

// Queue
//
// The FIFO queue structure. Please use NewQueue() or NewQueueWithValues() as the safe constructor.
type Queue struct {
	queue []interface{}
	head int  // points to the head of the queue
	tail int  // points to the first empty slot after the end of the queue
}

func (q *Queue) Len() int {
	return q.tail - q.head
}

func (q *Queue) HasNext() bool {
	return q.tail > q.head
}

// Push pushes v into the queue.
func (q *Queue) Push(v interface{}) {
	if q.tail == len(q.queue) {
		q.queue = append(q.queue, make([]interface{}, defaultSize)...)
	}
	q.queue[q.tail] = v
	q.tail ++
}

// Pop pops out and returns the first element of the queue.
//
// Please check if the queue is empty before using this method.
//
// e.g.,
//
// if queue.HasNext() { queue.Pop() }
func (q *Queue) Pop() interface{} {
	if q.tail <= q.head {
		return nil
	}
	pop := q.queue[q.head]
	q.head ++
	if q.head >= defaultSize {  // resize to save memory
		q.queue = q.queue[q.head:]
		q.tail -= q.head
		q.head = 0
	}
	return pop
}

// Values returns the values in the queue.
//
// This is NOT a safe method, and you should avoid using it.
func (q *Queue) Values() []interface{} {
	return q.queue
}

// Empty completely empties the queue.
//
// Note: it does not empty the queue physically; if you want to release memory, please create a new Queue object instead.
func (q *Queue) Empty() {
	q.head = 0
	q.tail = 0
}

// Copy makes a deep copy of the queue
func (q *Queue) Copy() *Queue {
	tmp := common.CopyInterfaces(q.queue)
	return &Queue{queue: tmp, head: q.head, tail: q.tail}
}

func NewQueue() *Queue {
	return &Queue{queue: make([]interface{}, defaultSize), head: 0, tail: 0}
}
