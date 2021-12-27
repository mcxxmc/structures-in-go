package tests

import (
	"some-data-structures/structures"
	"testing"
)

func TestBTree(t *testing.T) {
	nums := []int{1, 18, 2, 5, 19, 6, 7, 20, 21, 25, 12, 26, 10, 11, 22, 24, 13, 14, 15, 16, 17, 3, 4}
	correct := []int{1, 2, 3, 4, 5, 6, 7, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26}
	btree := structures.NewBTree(3, compareInt)

	// 1
	if btree.T() != 3 {
		t.Errorf("wrong t")
	}
	for _, num := range nums {
		btree.Insert(num)
	}
	if n := btree.NumOfElements(); n != len(nums) {
		t.Errorf("wrong number of elements; expecting %d, got %d", len(nums), n)
	}

	// 2
	values := btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree2: wrong value; expecting %d, got %d", v, values[i])
		}
	}
}
