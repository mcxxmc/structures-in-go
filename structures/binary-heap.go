package structures

import (
	"errors"
	"some-data-structures/common"
)

// BinaryHeap
//
// The binary heap structure. Please use NewBinaryHeapWithValues() or NewBinaryHeap() as the safe constructor.
//
// Attributes:
//
// Heap []interface{}. Should start from index 1 as index 0 is reserved.
//
// .
//
// compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b
type BinaryHeap struct {
	Heap    []interface{}
	compare func(a, b interface{}) int
	top     int // the rightmost boundary (inclusive)
}

func (bh *BinaryHeap) Size() int {
	return bh.top
}

func (bh *BinaryHeap) swap(i, j int) {
	bh.Heap[i], bh.Heap[j] = bh.Heap[j], bh.Heap[i]
}

// maintains the order of the binary heap
//
// Parameters:
//
// top: The index of the current root node of the subtree.
func (bh *BinaryHeap) maxHeapify(i int) {
	left := i << 1
	right := left + 1
	var largest int
	if left <= bh.top && bh.compare(bh.Heap[left], bh.Heap[i]) == 1 {
		largest = left
	} else {
		largest = i
	}
	if right <= bh.top && bh.compare(bh.Heap[right], bh.Heap[largest]) == 1 {
		largest = right
	}
	if largest != i {
		bh.swap(i, largest)
		bh.maxHeapify(largest)
	}
}

// sorts the heap and turns it into a binary heap.
func (bh *BinaryHeap) build() {
	for i := bh.top / 2; i > 0; i -- {
		bh.maxHeapify(i)
	}
}

// Heapsort sorts this binary heap into an ordered array. The order is reverse to the order of the binary heap. For
// example, if the binary heap is a maximum binary heap, then the order produce by Heapsort is ascending.
//
// It modifies the BinaryHeap object itself.
func (bh *BinaryHeap) Heapsort() {
	for i := bh.top; i > 1; i -- {
		bh.swap(i, 1)
		bh.top--
		bh.maxHeapify(1)
	}
	bh.top = len(bh.Heap) - 1 // reset
}

// HeapMaximum returns the maximum (or minimum, depending on the compare method) element from the binary heap
func (bh *BinaryHeap) HeapMaximum() (interface{}, error) {
	if bh.top < 1 {
		return nil, errors.New(errorNoElement)
	}
	return bh.Heap[1], nil
}

// ExtractHeapMaximum extracts and returns the maximum (or minimum, depending on the compare method)
// element from the binary heap.
//
// This operation will modify the binary heap.
func (bh *BinaryHeap) ExtractHeapMaximum() (interface{}, error) {
	if bh.top < 1 {
		return nil, errors.New(errorNoElement)
	}
	max := bh.Heap[1]            //todo: consider making a deep copy here
	bh.Heap[1] = bh.Heap[bh.top] // so the old bh.Heap[1] is no longer in the binary heap
	bh.top--
	bh.maxHeapify(1)
	return max, nil
}

// updates the key at position top
func (bh *BinaryHeap) updateKey(i int, key interface{}) error {
	if i > bh.top {
		return errors.New(errorInvalidIndex)
	}
	if bh.compare(bh.Heap[i], key) == 1 {
		return errors.New(errorKeyValue)
	}
	bh.Heap[i] = key
	for i > 1 && bh.compare(bh.Heap[i], bh.Heap[i / 2]) == 1 {
		bh.swap(i, i / 2)
		i = i / 2
	}
	return nil
}

// Insert inserts a key to the proper position in the binary heap.
func (bh *BinaryHeap) Insert(key interface{}) error {
	bh.top++
	if bh.top >= len(bh.Heap) { // needs to extend
		more := make([]interface{}, defaultSize)
		bh.Heap = append(bh.Heap, more...)
	}
	bh.Heap[bh.top] = key
	return bh.updateKey(bh.top, key)
}

// Copy makes a deep copy.
func (bh *BinaryHeap) Copy() *BinaryHeap {
	return &BinaryHeap{Heap: common.CopyInterfaces(bh.Heap), compare: bh.compare, top: bh.top}
}

// NewBinaryHeap returns a new BinaryHeap object with no initial values.
//
// compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b.
func NewBinaryHeap(compare func(a, b interface{}) int) *BinaryHeap {
	return &BinaryHeap{Heap: make([]interface{}, defaultSize), compare: compare, top: 0}
}

// NewBinaryHeapWithValues returns a new BinaryHeap object with initial values.
//
// values must be a slice or an array.
//
// compare is the function for comparing different node values;
// it should return 1 if a > b , 0 if a == b, -1 if a < b.
func NewBinaryHeapWithValues(values interface{}, compare func(a, b interface{}) int) (*BinaryHeap, error) {
	tmp, err := common.ToInterfaces(values)
	if err != nil {
		return nil, err
	}
	newValues := append([]interface{}{0}, tmp...)
	bh := &BinaryHeap{Heap: newValues, compare: compare, top: len(newValues) - 1}
	bh.build()
	return bh, nil
}
