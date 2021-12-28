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
		t.Errorf("BTree1: wrong t")
	}
	for _, num := range nums {
		btree.Insert(num)
	}
	if n := btree.NumOfElements(); n != len(nums) {
		t.Errorf("BTree1: wrong number of elements; expecting %d, got %d", len(nums), n)
	}

	// 2
	node, index, b := btree.Search(7)
	if !b {
		t.Errorf("BTree2.1: fail to search")
	}
	if node.Keys[index].(int) != 7 {
		t.Errorf("BTree2.1: wrong search result")
	}
	node, index, b = btree.Search(25)
	if !b {
		t.Errorf("BTree2.2: fail to search")
	}
	if node.Keys[index].(int) != 25 {
		t.Errorf("BTree2.2: wrong search result")
	}

	values := btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree2: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	//3
	correct = []int{1, 2, 3, 4, 5, 6, 7, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20, 21, 22, 24, 25, 26}
	btree.Delete(17)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.1: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	correct = []int{1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20, 21, 22, 24, 25, 26}
	btree.Delete(7)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.2: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	node, index, b = btree.Search(7)
	if b || node != nil || index != -1 {
		t.Errorf("BTree3: wrong search result")
	}

	if n := btree.NumOfElements(); n != len(correct) {
		t.Errorf("BTree3: wrong number of elements; expecting %d, got %d", len(correct), n)
	}

	correct = []int{1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16, 18, 19, 21, 22, 24, 25, 26}
	btree.Delete(20)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.3: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	correct = []int{1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16, 18, 19, 21, 22, 24, 25}
	btree.Delete(26)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.3: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	correct = []int{1, 2, 3, 5, 6, 10, 11, 12, 13, 14, 15, 16, 18, 19, 21, 22, 24, 25}
	btree.Delete(4)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.3: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	correct = []int{1, 2, 3, 5, 6, 10, 11, 12, 13, 14, 15, 16, 18, 19, 21, 22, 24, 25, 1000}
	btree.Insert(1000)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.3: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	correct = []int{1, 2, 3, 5, 6, 10, 11, 12, 13, 14, 15, 16, 18, 19, 21, 22, 24, 25, 25, 1000}
	btree.Insert(25)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.3: wrong value; expecting %d, got %d", v, values[i])
		}
	}

	correct = []int{1, 2, 3, 5, 6, 10, 11, 12, 13, 14, 15, 16, 18, 19, 21, 22, 24, 25, 1000}
	btree.Delete(25)
	values = btree.Values()
	for i, v := range correct {
		if v != values[i] {
			t.Errorf("BTree3.3: wrong value; expecting %d, got %d", v, values[i])
		}
	}
}
