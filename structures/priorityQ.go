package structures

import (
	"some-data-structures/common"
)

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

func (pq *PriorityQ) swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *PriorityQ) Push(v interface{}) {
	pq.queue = append(pq.queue, v)

	if len(pq.queue) == 1 {
		return
	}

	last := len(pq.queue) - 1

	if pq.compare(v, pq.queue[0]) != 1 {
		for i := 0; i < last; i ++ {
			pq.swap(i, last)
		}
		return
	}
	if pq.compare(pq.queue[last - 1], v) != 1 {
		return
	}

	right := last - 1
	left := 0
	mid := 0

	for {
		mid = (left + right) / 2
		if left >= right {
			break
		}
		c1 := pq.compare(v, pq.queue[mid])
		c2 := pq.compare(v, pq.queue[mid + 1])
		if c1 == 1 && c2 != 1 {
			break
		} else if c1 != 1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	for i := mid + 1; i < last; i ++ {
		pq.swap(i, last)
	}
}

func (pq *PriorityQ) Pop() interface{} {
	v := pq.queue[0]
	pq.queue = pq.queue[1:]
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
//
// compare: the function for comparing a and b
//
// for ascending order, compare must return 1 if a > b
//
// for descending order, compare must return 1 if a < b
func NewPriorityQ(compare func(a, b interface{}) int) *PriorityQ {
	return &PriorityQ{queue: make([]interface{}, 0), compare: compare}
}

