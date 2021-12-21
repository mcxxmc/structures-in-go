package tests

import (
	"some-data-structures/structures"
	"testing"
)

func TestBST(t *testing.T)  {
	insertions := []int{10, 5, 15, 3, 8, 20, 0, 24}

	// 1
	bTree := structures.NewIntBSTree()
	for _, num := range insertions {
		bTree.Insert(num)
	}
	if h := bTree.Height(); h != 4 {
		t.Errorf("BST1; expected height 4, got %d", h)
	}
	values1 := bTree.InOrderTreeWalk()
	correct1 := []int{0, 3, 5, 8, 10, 15, 20, 24}
	for i, val := range values1 {
		if val.(int) != correct1[i]{
			t.Errorf("BST1; wrong values")
		}
	}

	// 2
	node, b := bTree.Search(8)
	if !b {
		t.Errorf("BST2: fail to search")
	}
	successor := bTree.Successor(node)
	if successor == nil || successor.Val.(int) != 10 {
		t.Errorf("BST2: wrong successor")
	}
	predecessor := bTree.Predecessor(node)
	if predecessor == nil || predecessor.Val.(int) != 5 {
		t.Errorf("BST2: wrong predecessor")
	}
	node, b = bTree.Search(24)
	if !b {
		t.Errorf("BST2.2: fail to search")
	}
	successor = bTree.Successor(node)
	if successor != nil {
		t.Errorf("BST2.2: wrong successor")
	}

	// 3 rebuild a tree
	newTree := bTree.Rebuild()
	if h := newTree.Height(); h != 4 {
		t.Errorf("BST3; expected height 4, got %d", h)
	}
	values1 = newTree.InOrderTreeWalk()
	for i, val := range values1 {
		if val.(int) != correct1[i] {
			t.Errorf("BST3; wrong values")
		}
	}

	// 4 delete
	newTree.Delete(8)
	values1 = newTree.InOrderTreeWalk()
	correct1 = []int{0, 3, 5, 10, 15, 20, 24}
	for i, val := range values1 {
		if val.(int) != correct1[i] {
			t.Errorf("BST4.1; wrong values")
		}
	}

	newTree.Delete(20)
	if h := newTree.Height(); h != 3 {
		t.Errorf("BST4.2; expected height 3, got %d", h)
	}
	values1 = newTree.InOrderTreeWalk()
	correct1 = []int{0, 3, 5, 10, 15, 24}
	for i, val := range values1 {
		if val.(int) != correct1[i]{
			t.Errorf("BST4.2; wrong values")
		}
	}

	newTree.Delete(15)
	values1 = newTree.InOrderTreeWalk()
	correct1 = []int{0, 3, 5, 10, 24}
	for i, val := range values1 {
		if val.(int) != correct1[i] {
			t.Errorf("BST4.3; wrong values")
		}
	}

	newTree.Delete(5)
	values1 = newTree.InOrderTreeWalk()
	correct1 = []int{0, 3, 10, 24}
	for i, val := range values1 {
		if val.(int) != correct1[i] {
			t.Errorf("BST4.4; wrong values")
		}
	}

	// 5 customized
	// with a tricky compare
	compare := func(a, b interface{}) int {
		ra, _ := a.(*structures.Vector).AtD(1)
		var rb float64
		switch b.(type) {
		case float64:
			rb = b.(float64)
		case *structures.Vector:
			rb, _ = b.(*structures.Vector).AtD(1)
		}
		if ra > rb {
			return 1
		} else if ra == rb {
			return 0
		}
		return -1
	}
	customizedTree := structures.NewBSTree(compare)
	for _, num := range insertions {
		customizedTree.Insert(structures.NewVector([]float64{float64(num), float64(0)}))
	}
	if h := customizedTree.Height(); h != 4 {
		t.Errorf("BST5; expected height 4, got %d", h)
	}
	values1 = customizedTree.InOrderTreeWalk()
	correct2 := []float64{0.0, 3.0, 5.0, 8.0, 10.0, 15.0, 20.0, 24.0}
	for i, v := range values1 {
		if val, _ := v.(*structures.Vector).AtD(1); val != correct2[i] {
			t.Errorf("BST5.1; wrong values")
		}
	}
	vector, _ := customizedTree.Search(5.0)
	if !vector.Val.(*structures.Vector).Equal(structures.NewVector([]float64{5.0, 0.0})) {
		t.Errorf("BST5.1; wrong vector")
	}
	b = customizedTree.Delete(10.0)
	if !b {
		t.Errorf("BST5.1: fail to delete")
	}
	values1 = customizedTree.InOrderTreeWalk()
	correct2 = []float64{0.0, 3.0, 5.0, 8.0, 15.0, 20.0, 24.0}
	for i, v := range values1 {
		if val, _ := v.(*structures.Vector).AtD(1); val != correct2[i] {
			t.Errorf("BST5.2; wrong values")
		}
	}
}
