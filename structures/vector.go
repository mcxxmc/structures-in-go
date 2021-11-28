package structures

import (
	"fmt"
	"math"
	"strconv"
)

// Vector a vector
//
// should not be modified manually after it is created; please use method calls instead
//
// Note that it implements the Template interface
type Vector struct {
	v []float64
	d int  // dimension
}

// Negative returns the negative vector
//
// e.g., x + (-x) = 0, then -x is the negative vector of x
func (v *Vector) Negative() *Vector {
	tmp := make([]float64, v.d)
	for i := 0; i < v.d; i ++ {
		tmp[i] = - v.v[i]
	}
	return NewVector(tmp)
}

// Magnitude returns the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	tmp := float64(0)
	for i := 0; i < v.d; i ++ {
		tmp += math.Pow(v.v[i], float64(2))
	}
	return math.Sqrt(tmp)
}

// D returns the dimension of the vector
func (v *Vector) D() int {
	return v.d
}

// V returns the vector in []float64 form
//
// note that it returns a (deep) copy of the value of the vector
func (v *Vector) V() []float64 {
	tmp := make([]float64, v.d)
	copy(tmp, v.v)
	return tmp
}

// AtD returns the value at a given dimension
//
// n : the nth dimension (starting from 1)
func (v *Vector) AtD(n int) (float64, bool) {
	if n > v.d {
		return float64(0), false
	}
	return v.v[n - 1], true
}

// Multiple multiplies this Vector with a number
//
// NOTE: this operation does not change the vector; it returns a new Vector instead
func (v *Vector) Multiple(n float64) *Vector {
	tmp := make([]float64, v.d)
	for i, val := range v.v {
		tmp[i] = val * n
	}
	return NewVector(tmp)
}

// Unit returns a unit Vector of the same direction
func (v *Vector) Unit() *Vector {
	mag := v.Magnitude()
	tmp := make([]float64, v.d)
	for i, val := range v.v {
		tmp[i] = val / mag
	}
	return NewVector(tmp)
}

// Add returns a new Vector as the addition of 2 Vectors
//
// the 2 Vectors must have the same dimension
func (v *Vector) Add(v2 *Vector) (*Vector, bool) {
	if v.d != v2.D() {
		fmt.Println("Dimensions of vectors do not match " + strconv.Itoa(v.d) + " vs " + strconv.Itoa(v2.D()))
		return nil, false
	}
	tmp := make([]float64, v.d)
	vs2 := v2.V()
	for i, val := range v.v {
		tmp[i] = val + vs2[i]
	}
	return NewVector(tmp), true
}

// ReDimension returns a NEW vector of a different dimension
//
// will return nil if n <= 0
func (v *Vector) ReDimension(n int) (*Vector, bool) {
	if n <= 0 {
		return nil, false
	} else {
		tmp := make([]float64, n)
		copy(tmp, v.v)
		return NewVector(tmp), true
	}
}

// Expand returns a NEW Vector of a higher dimension
//
// will return a copy of the old Vector itself if n is smaller than or equal to the current dimension
//
// WARNING: you should not use this function to make a copy of the Vector; use Copy() instead
func (v *Vector) Expand(n int) (*Vector, bool) {
	if n <= v.d {
		return v.Copy().(*Vector), false
	} else {
		tmp := make([]float64, n)
		copy(tmp, v.v)
		return NewVector(tmp), true
	}
}

// Minus returns a new Vector as the difference between this Vector v1 and the other Vector v2 ( = v1 - v2)
//
// the 2 Vectors must have the same dimension
func (v *Vector) Minus(v2 *Vector) (*Vector, bool) {
	if v.d != v2.D() {
		fmt.Println("Dimensions of vectors do not match " + strconv.Itoa(v.d) + " vs " + strconv.Itoa(v2.D()))
		return nil, false
	} else {
		tmp := make([]float64, v.d)
		vs2 := v2.V()
		for i, val := range v.v {
			tmp[i] = val - vs2[i]
		}
		return NewVector(tmp), true
	}
}

// Dot returns a number as the dot product of 2 Vectors
//
// the 2 Vectors must have the same dimension
func (v *Vector) Dot(v2 *Vector) (float64, bool) {
	if v.d != v2.D() {
		fmt.Println("Dimensions of vectors do not match " + strconv.Itoa(v.d) + " vs " + strconv.Itoa(v2.D()))
		return float64(0), false
	} else {
		total := float64(0)
		vs2 := v2.V()
		for i, val := range v.v {
			total += val * vs2[i]
		}
		return total, true
	}
}

// Projection returns a NEW Vector as the projection onto another Vector v2
//
// the projection is parallel with v2
//
// the 2 Vectors must have the same dimension
func (v *Vector) Projection(v2 *Vector) (*Vector, bool) {
	if v.d != v2.D() {
		fmt.Println("Dimensions of vectors do not match " + strconv.Itoa(v.d) + " vs " + strconv.Itoa(v2.D()))
		return nil, false
	}
	mul, _ := v.Dot(v2)
	magV2 := v2.Magnitude()
	return v2.Multiple(mul / (magV2 * magV2)), true
}

// CrossProduct returns a NEW Vector as the cross product of Vectors v and v2
//
// Vectors v and v2 must both have a dimension of 3
func (v *Vector) CrossProduct(v2 *Vector) (*Vector, bool) {
	if v.d != 3 || v2.D() != 3 {
		fmt.Println("Dimensions for cross product must be 3 but " + strconv.Itoa(v.d) + " vs " + strconv.Itoa(v2.D()))
		return nil, false
	}
	vs2 := v2.V()
	x1, y1, z1 := v.v[0], v.v[1], v.v[2]
	x2, y2, z2 := vs2[0], vs2[1], vs2[2]
	return NewVector([]float64{y1 * z2 - z1 * y2, z1 * x2 - x1 * z2, x1 * y2 - y1 * x2}), true
}


// Copy makes a deep copy
func (v *Vector) Copy() interface{} {
	tmp := make([]float64, v.d)
	copy(tmp, v.v)
	return NewVector(tmp)
}

// String stringify
func (v *Vector) String() string {
	s := "Dimension " + strconv.Itoa(v.d) + " ["
	for i, f := range v.v {
		s += strconv.FormatFloat(f, 'f', -1, 64)
		if i != v.d - 1 {
			s += ", "
		} else {
			s += "]"
		}
	}
	return s
}


// NewVector returns a new Vector
func NewVector(v []float64) *Vector {
	return &Vector{v: v, d: len(v)}
}

// ZeroVector returns a zero Vector of dimension d
func ZeroVector(d int) *Vector {
	return &Vector{v: make([]float64, d), d: d}
}
