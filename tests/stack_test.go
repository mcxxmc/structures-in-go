package tests

import (
	"some-data-structures/structures"
	"testing"
)

func TestStack(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	stk1 := structures.NewStack()
	for _, num := range nums {
		stk1.Push(num)
	}
	if l := stk1.Len(); l != 4 {
		t.Errorf("expected length 4, but got %d", l)
	}
	stk2 := stk1.Copy()
	if p := stk2.Pop(); p.(int) != 4 {
		t.Errorf("expected 4, but got %d", p)
	}
	stk2.Empty()
	if l := stk2.Len(); l != 0 {
		t.Errorf("expected length 0, but got %d", l)
	}
	l := len(nums)
	for i := 0; i < l; i ++ {
		if stk1.HasNext() {
			if p := stk1.Pop(); p != nums[l - i - 1] {
				t.Errorf("expected %d, but got %d", nums[l - i - 1], p)
			}
		}
	}
}
