package structures

// Queue a queue that first in first out
//
// IMPORTANT: use NewQueue() to construct a new Queue object!
type Queue struct {
	queue []interface{}
	length int
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

// Add adds a new element to the end of the queue
func (q *Queue) Add(val interface{}) {
	q.queue = append(q.queue, val)
	q.length += 1
}

// Pop pops out and returns the first element of the queue
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}

	pop := q.queue[0]
	q.queue = q.queue[1:]
	q.length -= 1

	return pop
}

// Reset completely resets the queue
//
// Warning: it will empty the queue
func (q *Queue) Reset() {
	q.queue = make([]interface{}, 0)
	q.length = 0
}

// Copy makes a deep copy of the queue
func (q *Queue) Copy() *Queue {
	tmp := make([]interface{}, q.length)
	copy(tmp, q.queue)
	return &Queue{queue: tmp, length: q.length}
}

func NewQueue() *Queue {
	return &Queue{queue: make([]interface{}, 0), length: 0}
}

func NewQueueWithValues(values []interface{}) *Queue {
	return &Queue{queue: values, length: len(values)}
}
