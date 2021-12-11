package structures

// BinaryIndexedTree the binary indexed tree
//
// this struct accepts float64 values as input; indexes start from 0
type BinaryIndexedTree struct {
	values []float64
}

func (bit *BinaryIndexedTree) lowbit(x int) int {
	return x & (-x)
}

// Update updates / increases the value for the element at index i
func (bit *BinaryIndexedTree) Update(i int, v float64) {
	for t := i + 1; t < len(bit.values); t += bit.lowbit(t) {
		bit.values[t] += v
	}
}

// Query returns the prefix sum up to the element at index i (inclusive)
func (bit *BinaryIndexedTree) Query(i int) float64 {
	ans := float64(0)
	for t := i + 1; t > 0; t -= bit.lowbit(t) {
		ans += bit.values[t]
	}
	return ans
}

// Range returns the prefix sum from the element at index i (inclusive) to the element at index j (inclusive)
func (bit *BinaryIndexedTree) Range(i, j int) float64 {
	return bit.Query(j) - bit.Query(i - 1)
}

// Size returns the size of this tree
func (bit *BinaryIndexedTree) Size() int {
	return len(bit.values) - 1
}

// Snapshot returns a deep copy of the current values in this tree
func (bit *BinaryIndexedTree) Snapshot() []float64 {
	tmp := make([]float64, len(bit.values))
	copy(tmp, bit.values)
	return tmp
}

// Copy returns a deep copy of this tree
func (bit *BinaryIndexedTree) Copy() *BinaryIndexedTree {
	tmp := bit.Snapshot()
	return &BinaryIndexedTree{values: tmp}
}

func NewBinaryIndexedTree(n int) *BinaryIndexedTree {
	return &BinaryIndexedTree{values: make([]float64, n + 1)}
	// the real size is n + 1; user index 0 corresponds to actual index 1; values[0] will be forever 0
}
