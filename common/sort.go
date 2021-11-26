package common

import "fmt"

// Sorter a helper struct for sorting interfaces
//
// Compare: the compare method. Must return -1 when a < b.
//
// IMPORTANT: use NewSorter to create a new Sorter object
type Sorter struct {
	Compare func(a, b interface{}) int
}

// Sort sorts the slice of interfaces; notes that it modifies the original slice
//
// uses quicksort
//
// the returned bool shows if the sorting is successful
func (sorter *Sorter) Sort(interfaces []interface{}) bool {
	if sorter.Compare == nil {
		fmt.Println("Error: no compare method is defined.")
		return false
	}

	sorter.QuickSort(interfaces, 0, len(interfaces) - 1)
	return true
}

func (sorter *Sorter) QuickSort(a []interface{}, left, right int) {
	if right > left {
		newPivot := (left + right) / 2
		newPivot = sorter.partition(a, left, right, newPivot)
		sorter.QuickSort(a, left, newPivot - 1)
		sorter.QuickSort(a, newPivot + 1, right)
	}
}

func (sorter *Sorter) partition(a []interface{}, left, right, pivotIndex int) int {
	pivotValue := Copy(a[pivotIndex])
	a[pivotIndex], a[right] = a[right], a[pivotIndex]  // move a[pivotIndex] to the end
	storeIndex := left

	for i := left; i < right; i ++ {
		if sorter.Compare(pivotValue, a[i]) != -1 {  // pivotValue > a[i]
			a[storeIndex], a[i] = a[i], a[storeIndex]
			storeIndex += 1
		}
	}

	a[right], a[storeIndex] = a[storeIndex], a[right]
	return storeIndex
}

func NewSorter(compare func(a, b interface{}) int) *Sorter {
	return &Sorter{Compare: compare}
}

