package tests

import (
	"some-data-structures/structures"
	"sort"
	"testing"
)

// used to find the number of reverse pairs in the array
func reversePairs(nums []float64) int {
	l := len(nums)
	tmp := make([]float64, l)
	copy(tmp, nums)
	sort.Float64s(tmp)

	indexes := make([]int, l)

	for i := 0; i < l; i ++ {
		indexes[i] = sort.SearchFloat64s(tmp, nums[i])
	}

	bit := structures.NewBinaryIndexedTree(l)

	ans := 0.0
	for i := l - 1; i >= 0; i -- {
		ans += bit.Query(indexes[i])
		bit.Update(indexes[i], 1.0)
	}

	return int(ans)
}

// tests the BinaryIndexedTree by finding the number of reverse pairs in arrays
func TestBIT(t *testing.T) {
	nums := []float64{7.0, 5.0, 6.0, 4.0}
	if rp := reversePairs(nums); rp != 5 {
		t.Errorf("TestBIT 1; expected 5, but got %d", rp)
	}
	nums = []float64{7.0, 5.0, 6.0, 4.0, 10.0, 15.0, 8.0, -7.0}
	if rp := reversePairs(nums); rp != 14 {
		t.Errorf("TestBIT 1; expected 14, but got %d", rp)
	}
}
