package structures

import "some-data-structures/common"

// Queue a queue that first in first out
//
// IMPORTANT: use NewQueue() to construct a new Queue object!
type Queue struct {
	queue []interface{}
}

func (q *Queue) Len() int {
	return len(q.queue)
}

func (q *Queue) HasNext() bool {
	return len(q.queue) != 0
}

// Push adds a new element to the end of the queue
func (q *Queue) Push(val interface{}) {
	q.queue = append(q.queue, val)
}

// Pop pops out and returns the first element of the queue
//
// will return nil if the queue is empty; however, a nil does not mean there is no more values in this queue,
// since you can push nil as a value into this queue
//
// a safe way to check if the queue is empty is using HasNext
func (q *Queue) Pop() interface{} {
	if len(q.queue) == 0 {
		return nil
	}

	pop := q.queue[0]
	q.queue = q.queue[1:]

	return pop
}

// Reset completely resets the queue
//
// Warning: it will empty the queue
func (q *Queue) Reset() {
	q.queue = make([]interface{}, 0)
}

// Copy makes a deep copy of the queue
func (q *Queue) Copy() *Queue {
	cpy := &Queue{queue: make([]interface{}, len(q.queue))}
	for i := 0; i < len(q.queue); i ++ {
		cpy.queue[i] = common.Copy(q.queue[i])
	}
	return cpy
}

func NewQueue() *Queue {
	return &Queue{queue: make([]interface{}, 0)}
}

// NewQueueWithValues creates a Queue with initial values
//
// WARNING: modifying the values will also modify the Queue
func NewQueueWithValues(values []interface{}) *Queue {
	return &Queue{queue: values}
}
