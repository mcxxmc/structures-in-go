package tests

import (
	"some-data-structures/structures"
	"testing"
)

func TestLinkedList(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	l := len(nums)
	compare := func(a, b interface{}) int {
		if a.(int) == b.(int) {
			return 0
		}
		return 1
	}

	// 1
	ll := structures.NewLinkedList(compare)
	for i := 0; i < l; i ++ {
		ll.Insert(nums[l - 1 - i])
	}
	s1 := ll.Search(5)
	if s1 == nil || s1.Val.(int) != 5 {
		t.Errorf("TestLinkedList1: wrong search")
	}
	for i := 0; i < l; i ++ {
		s1 = ll.Delete(nums[i])
		if s1 == nil || s1.Val.(int) != nums[i] {
			t.Errorf("TestLinkedList1: wrong delete")
		}
	}

	// 2
	dll := structures.NewDoubleLinkedList(compare)
	for i := 0; i < l; i ++ {
		dll.Insert(nums[l - 1 - i])
	}
	s2 := dll.Search(5)
	if s2 == nil || s2.Val.(int) != 5 {
		t.Errorf("TestLinkedList2: wrong search")
	}
	for i := 0; i < l; i ++ {
		s2 = dll.Delete(nums[i])
		if s2 == nil || s2.Val.(int) != nums[i] {
			t.Errorf("TestLinkedList2: wrong delete")
		}
	}
}
