package tests

import (
	"some-data-structures/structures"
	"testing"
)

func TestRedBlackTree(t *testing.T) {
	insertions := []int{16, 3, 7, 11, 9, 26, 18, 14, 15}

	// 1
	tree := structures.NewRedBlackTree(compareInt)
	for _, num := range insertions {
		tree.Insert(num)
	}
	if h := tree.Height(); h != 4 {
		t.Errorf("RBT1; expected height 4, got %d", h)
	}
	values1 := tree.InOrderTreeWalk()
	correct1 := []int{3, 7, 9, 11, 14, 15, 16, 18, 26}
	for i, val := range values1 {
		if val.(int) != correct1[i]{
			t.Errorf("RBT1; wrong values")
		}
	}

	// 2
	node, b := tree.Search(9)
	if !b {
		t.Errorf("RBT2: fail to search")
	}
	successor := tree.Successor(node)
	if successor == nil || successor.Val.(int) != 11 {
		t.Errorf("RBT2: wrong successor")
	}
	predecessor := tree.Predecessor(node)
	if predecessor == nil || predecessor.Val.(int) != 7 {
		t.Errorf("RBT2: wrong predecessor")
	}
	node, b = tree.Search(26)
	if !b {
		t.Errorf("RBT2.2: fail to search")
	}
	successor = tree.Successor(node)
	if successor != nil {
		t.Errorf("RBT2.2: wrong successor")
	}

	// 3
	tree.Delete(14)
	values1 = tree.InOrderTreeWalk()
	correct1 = []int{3, 7, 9, 11, 15, 16, 18, 26}
	for i, val := range values1 {
		if val.(int) != correct1[i]{
			t.Errorf("RBT3.1; wrong values")
		}
	}
	if h := tree.Height(); h != 4 {
		t.Errorf("RBT3.1; expected height 4, got %d", h)
	}

	tree.Delete(15)
	values1 = tree.InOrderTreeWalk()
	correct1 = []int{3, 7, 9, 11, 16, 18, 26}
	for i, val := range values1 {
		if val.(int) != correct1[i]{
			t.Errorf("RBT3.2; wrong values")
		}
	}
	if h := tree.Height(); h != 3 {
		t.Errorf("RBT3.2; expected height 3, got %d", h)
	}

	tree.Delete(7)
	values1 = tree.InOrderTreeWalk()
	correct1 = []int{3, 9, 11, 16, 18, 26}
	for i, val := range values1 {
		if val.(int) != correct1[i]{
			t.Errorf("RBT3.3; wrong values")
		}
	}
	if h := tree.Height(); h != 3 {
		t.Errorf("RBT3.3; expected height 3, got %d", h)
	}

	tree.Delete(11)
	values1 = tree.InOrderTreeWalk()
	correct1 = []int{3, 9, 16, 18, 26}
	for i, val := range values1 {
		if val.(int) != correct1[i]{
			t.Errorf("RBT3.4; wrong values")
		}
	}
	if h := tree.Height(); h != 3 {
		t.Errorf("RBT3.4; expected height 3, got %d", h)
	}
}
