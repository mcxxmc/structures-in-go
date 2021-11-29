package structures

import "some-data-structures/common"

// PriorityQ the priority queue
//
// compare: the function for comparing a and b
//
// for ascending order, compare must return 1 if a > b
//
// for descending order, compare must return 1 if a < b
type PriorityQ struct {
	queue []interface{}
	compare func(a, b interface{}) int
}

func (pq *PriorityQ) Len() int {
	return len(pq.queue)
}

func (pq *PriorityQ) HasNext() bool {
	return len(pq.queue) != 0
}

func (pq *PriorityQ) Less(i, j int) bool {
	if pq.compare(pq.queue[i], pq.queue[j]) == 1 {
		return false
	}
	return true
}

func (pq *PriorityQ) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *PriorityQ) Push(v interface{}) {
	pq.queue = append(pq.queue, v)
}

func (pq *PriorityQ) Pop() interface{} {
	v := pq.queue[len(pq.queue) - 1]
	pq.queue = pq.queue[: len(pq.queue) - 1]
	return v
}

// Reset completely resets the queue
//
// Warning: it will empty the queue
func (pq *PriorityQ) Reset() {
	pq.queue = make([]interface{}, 0)
}

// Copy makes a deep copy
func (pq *PriorityQ) Copy() *PriorityQ {
	cpy := &PriorityQ{queue: make([]interface{}, len(pq.queue)), compare: pq.compare}
	for i := 0; i < len(pq.queue); i ++ {
		cpy.queue[i] = common.Copy(pq.queue[i])
	}
	return cpy
}

// NewPriorityQ creates a new priority queue PriorityQ
func NewPriorityQ(compare func(a, b interface{}) int) *PriorityQ {
	return &PriorityQ{queue: make([]interface{}, 0), compare: compare}
}

