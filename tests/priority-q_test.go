package tests

import (
	"some-data-structures/structures"
	"sort"
	"testing"
)

func TestPriorityQ(t *testing.T) {
	nums := []int{9, 8, 7, 1, 14, 12, -8, 10, 6, 5}
	compare := func(a, b interface{}) int {
		if a.(int) > b.(int) {
			return 1
		}
		return 0
	}
	pq := structures.NewPriorityQ(compare)
	for _, num := range nums {
		pq.Push(num)
	}
	if l := pq.Len(); l != len(nums) {
		t.Errorf("expected priority length %d, got %d", len(nums), l)
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i ++ {
		if tmp := pq.Pop().(int); nums[i] != tmp {
			t.Errorf("expected %d th element %d, but got %d", i, nums[i], tmp)
		}
	}
	if l := pq.Len(); l != 0 {
		t.Errorf("expected priority length 0, got %d", l)
	}
}
