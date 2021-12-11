package tests

import (
	"fmt"
	"some-data-structures/structures"
	"testing"
)

func TestBST(t *testing.T)  {
	insertions := []int{10, 5, 15, 3, 8, 20, 0, 24}

	// 1
	bTree := structures.NewBSTreeInt()
	for _, num := range insertions {
		bTree.Insert(num)
	}
	if h := bTree.Height(); h != 4 {
		t.Errorf("BST1; expected height 4, got %d", h)
	}
	vals1 := bTree.Values(false)
	correct1 := [][]int{{10}, {5, 15}, {3, 8, 20}, {0, 24}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if v.(int) != correct1[i][j] {
				t.Errorf("BST1; wrong values")
			}
		}
	}
	vals2 := bTree.FlattenedValues(false)
	correct2 := []int{10, 5, 15, 3, 8, 20, 0, 24}
	for i, v := range vals2 {
		if v.(int) != correct2[i] {
			t.Errorf("BST1; wrong flattened values")
		}
	}

	// 2
	copied := bTree.Copy()
	if h := copied.Height(); h != 4 {
		t.Errorf("BST2; expected height 4, got %d", h)
	}

	// 3 rebuild a tree
	newTree := bTree.Rebuild()
	if h := newTree.Height(); h != 4 {
		t.Errorf("BST3; expected height 4, got %d", h)
	}
	vals1 = newTree.Values(false)
	correct1 = [][]int{{8}, {3, 15}, {0, 5, 10, 20}, {24}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if v.(int) != correct1[i][j] {
				t.Errorf("BST3; wrong values")
			}
		}
	}

	// 4 delete
	newTree.Delete(8)
	vals1 = newTree.Values(false)
	correct1 = [][]int{{5}, {3, 15}, {0, 10, 20}, {24}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if v.(int) != correct1[i][j] {
				t.Errorf("BST4.1; wrong values")
			}
		}
	}
	newTree.Delete(20)
	if h := newTree.Height(); h != 3 {
		t.Errorf("BST4.2; expected height 3, got %d", h)
	}
	vals1 = newTree.Values(false)
	correct1 = [][]int{{5}, {3, 15}, {0, 10, 24}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if v.(int) != correct1[i][j] {
				t.Errorf("BST4.2; wrong values")
			}
		}
	}
	newTree.Delete(15)
	vals1 = newTree.Values(false)
	correct1 = [][]int{{5}, {3, 10}, {0, 24}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if v.(int) != correct1[i][j] {
				t.Errorf("BST4.3; wrong values")
			}
		}
	}
	newTree.Delete(5)
	vals1 = newTree.Values(false)
	correct1 = [][]int{{3}, {0, 10}, {24}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if v.(int) != correct1[i][j] {
				fmt.Println(vals1)
				t.Errorf("BST4.4; wrong values")
			}
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
	vals1 = customizedTree.Values(false)
	correct3 := [][]float64{{10.0}, {5.0, 15.0}, {3.0, 8.0, 20.0}, {0.0, 24.0}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if val, _ := v.(*structures.Vector).AtD(1); val != correct3[i][j] {
				t.Errorf("BST5.1; wrong values")
			}
		}
	}
	vector, _ := customizedTree.Search(5.0)
	if !vector.(*structures.Vector).Equal(structures.NewVector([]float64{5.0, 0.0})) {
		t.Errorf("BST5.1; wrong vector")
	}
	vector, _ = customizedTree.Delete(10.0)
	if !vector.(*structures.Vector).Equal(structures.NewVector([]float64{10.0, 0.0})) {
		t.Errorf("BST5.2; wrong vector")
	}
	vals1 = customizedTree.Values(false)
	correct3 = [][]float64{{8.0}, {5.0, 15.0}, {3.0, 20.0}, {0.0, 24.0}}
	for i, slc := range vals1 {
		for j, v := range slc {
			if val, _ := v.(*structures.Vector).AtD(1); val != correct3[i][j] {
				t.Errorf("BST5.2; wrong values")
			}
		}
	}
}
