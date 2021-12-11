package tests

import (
	"reflect"
	"some-data-structures/structures"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	testSlice := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	compare := func(a, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) == b.(int) {
			return 0
		} else {
			return -1
		}
	}

	// 1
	bh, err := structures.NewBinaryHeap(testSlice, compare)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(bh.Heap, []interface{}{0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1}) {
		t.Errorf("TestBinaryHeap1: wrong heap")
	}

	// 2
	tmp := bh.Copy()
	tmp.Heapsort()
	if !reflect.DeepEqual(tmp.Heap, []interface{}{0, 1, 2, 3, 4, 7, 8, 9, 10, 14, 16}) {
		t.Errorf("TestBinaryHeap2: wrong heap")
	}

	// 3
	val, err := bh.HeapMaximum()
	if err != nil {
		t.Error(err)
	}
	if val.(int) != 16 {
		t.Errorf("TestBinaryHeap3: wrong maximum")
	}
	val, err = bh.ExtractHeapMaximum()
	if err != nil {
		t.Error(err)
	}
	if val.(int) != 16 {
		t.Errorf("TestBinaryHeap3: wrong extracted maximum")
	}

	// 4
	err = bh.Insert(20)
	if err != nil {
		t.Error(err)
	}
	err = bh.Insert(6)
	if err != nil {
		t.Error(err)
	}
	sorted := []int{20, 14, 10, 9, 8, 7, 6, 4, 3, 2, 1}
	for i := 0; i < len(sorted); i ++ {
		v, err := bh.ExtractHeapMaximum()
		if err != nil {
			t.Error(err)
		}
		if v != sorted[i] {
			t.Errorf("TestBinaryHeap4: wrong extracted maximum")
		}
	}
	_, err = bh.ExtractHeapMaximum()
	if err == nil {
		t.Errorf("TestBinaryHeap4: more values extracted than expected")
	}
}