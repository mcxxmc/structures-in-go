package tests

import (
	"some-data-structures/structures"
	"testing"
)

func TestFibonacciHeap(t *testing.T) {
	nums := []int{7, 18, 38, 24, 17, 23, 21, 39, 41, 26, 46, 30, 52, 35}
	sorted := []int{7, 17, 18, 21, 23, 24, 26, 30, 35, 38, 39, 41, 46, 52}
	fh := structures.NewFibonacciHeap(compareInt)
	m := make(map[int]*structures.FibNode)
	index := 0

	// 1
	for _, num := range nums {
		node := fh.Insert(num)
		m[num] = node
	}
	if n := fh.NumOfElements(); n != len(nums) {
		t.Errorf("TestFibonacciHeap1.1: wrong element numbers; expected %d, got %d", len(nums), n)
	}
	if n := fh.Minimum().Val.(int); n != sorted[index] {
		t.Errorf("TestFibonacciHeap1: wrong minimun; expected %d, got %d", sorted[index], n)
	}

	for i := 0; i < 10; i ++ {
		n := fh.ExtractMin().Val.(int)
		if n != sorted[index] {
			t.Errorf("TestFibonacciHeap1.2: wrong extraction; expected %d, got %d", sorted[index], n)
		}
		m[n] = nil
		index ++
	}
	if n := fh.NumOfElements(); n != len(nums) - index {
		t.Errorf("TestFibonacciHeap1.2: wrong element numbers; expected %d, got %d", len(nums) - index, n)
	}

	for i := 0; i < 10; i ++ {
		index --
		node := fh.Insert(sorted[index])
		m[sorted[index]] = node
	}
	if n := fh.NumOfElements(); n != len(nums) {
		t.Errorf("TestFibonacciHeap1.3: wrong element numbers; expected %d, got %d", len(nums), n)
	}
	// check if m is good
	for _, v := range sorted {
		if node := m[v]; node == nil {
			t.Errorf("TestFibonacciHeap1.1: nil in map; key %d", v)
		} else if n := node.Val.(int); n != v {
			t.Errorf("TestFibonacciHeap1.1: wrong map; key %d", v)
		}
	}

	for i := 0; i < 10; i ++ {
		node := m[sorted[index]]
		err := fh.Delete(node, -1)
		if err != nil {
			t.Errorf("TestFibonacciHeap1: wrong deletion")
		}
		m[sorted[index]] = nil
		index ++
	}
	if n := fh.NumOfElements(); n != len(nums) - index {
		t.Errorf("TestFibonacciHeap1.4: wrong element numbers; expected %d, got %d", len(nums) - index, n)
	}
	if n := fh.Minimum().Val.(int); n != sorted[index] {
		t.Errorf("TestFibonacciHeap1: wrong minimun; expected %d, got %d", sorted[index], n)
	}

	fh2 := structures.NewFibonacciHeap(compareInt)
	for i := 0; i < index; i ++ {
		node := fh2.Insert(sorted[i])
		m[sorted[i]] = node
	}
	// check if m is good
	for _, v := range sorted {
		if node := m[v]; node == nil {
			t.Errorf("TestFibonacciHeap1.1: nil in map; key %d", v)
		} else if n := node.Val.(int); n != v {
			t.Errorf("TestFibonacciHeap1.1: wrong map; key %d", v)
		}
	}

	fh3 := fh.Union(fh2)
	if n := fh3.NumOfElements(); n != len(nums) {
		t.Errorf("TestFibonacciHeap1.5: wrong element numbers; expected %d, got %d", len(nums), n)
	}

	sorted2 := []int{1, 7, 17, 18, 21, 23, 24, 26, 30, 35, 36, 39, 41, 46}
	_ = fh3.DecreaseKey(m[52], 1)
	m[1] = m[52]
	m[52] = nil
	_ = fh3.DecreaseKey(m[38], 36)
	m[36] = m[38]
	m[38] = nil
	if n := fh3.NumOfElements(); n != len(nums) {
		t.Errorf("TestFibonacciHeap1.6: wrong element numbers; expected %d, got %d", len(nums), n)
	}

	for _, v := range sorted2 {
		n := fh3.ExtractMin().Val.(int)
		if n != v {
			t.Errorf("TestFibonacciHeap1.2: wrong extraction; expected %d, got %d", v, n)
		}
	}
}
