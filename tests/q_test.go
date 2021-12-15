package tests

import (
	"some-data-structures/structures"
	"testing"
)

func TestQueue(t *testing.T)  {
	nums := make([]int, 100)
	for i := 0; i < 100; i ++ {
		nums[i] = i
	}

	// 1
	q := structures.NewQueue()
	for i := 0; i < 100; i ++ {
		q.Push(nums[i])
	}
	if q.Len() != 100 {
		t.Errorf("TestQueue1: wrong queue size")
	}
	q = q.Copy()
	for i := 0; i < 100; i ++ {
		tmp := q.Pop().(int)
		if tmp != nums[i] {
			t.Errorf("TestQueue1: wrong pop")
		}
	}
	if q.HasNext() || q.Len() != 0 {
		t.Errorf("TestQueue1.2: wrong queue size")
	}

	// 2
	q.Empty()
	if q.HasNext() || q.Len() != 0 {
		t.Errorf("TestQueue2: wrong queue size")
	}
	for i := 0; i < 50; i ++ {
		q.Push(nums[i])
	}
	for i := 0; i < 50; i ++ {
		tmp := q.Pop()
		if tmp != nums[i] {
			t.Errorf("TestQueue2.1: wrong pop")
		}
	}
	for i := 50; i < 100; i ++ {
		q.Push(nums[i])
	}
	for i := 50; i < 100; i ++ {
		tmp := q.Pop()
		if tmp != nums[i] {
			t.Errorf("TestQueue2.2: wrong pop")
		}
	}
	if q.HasNext() || q.Len() != 0 {
		t.Errorf("TestQueue2.2: wrong queue size")
	}
}
