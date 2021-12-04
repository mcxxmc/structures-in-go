package tests

import (
	"fmt"
	"some-data-structures/structures"
	"testing"
)

func TestVector(t *testing.T) {
	floats := []float64{3.0, 4.0}
	vector1 := structures.NewVector(floats)

	floats[0] = 5.0

	if v, _ := vector1.AtD(1); v != 3.0 {
		t.Errorf("changed vector value")
	}

	if vector1.Magnitude() != 5.0 {
		t.Errorf("wrong magnitude")
	}

	floats[0] = 3.0

	zero := structures.ZeroVector(vector1.D())
	tmp, _ := vector1.Add(zero)

	if !vector1.Equal(tmp) {
		t.Errorf("fail to add zero")
	}

	tmp = vector1.Negative()
	if vs := tmp.V(); vs[0] != -3.0 || vs[1] != -4.0 {
		t.Errorf("wrong negative vector")
	}

	tmp = vector1.Unit()
	if vs := tmp.V(); vs[0] != 0.6 || vs[1] != 0.8 {
		t.Errorf("wrong unit vector")
	}

	tmp, _ = vector1.Minus(vector1)
	if !tmp.Equal(zero) {
		t.Errorf("wrong minus result")
	}

	tmp = vector1.Multiple(2.0)
	if tmp.Magnitude() != 10.0 {
		t.Errorf("wrong scaling")
	}

	tmp, _ = vector1.ReDimension(1)
	if tmp.Magnitude() != 3.0 {
		t.Errorf("fail to redimension")
	}

	tmp, _ = vector1.Expand(3)
	if vs := tmp.V(); vs[0] != 3.0 || vs[1] != 4.0 || vs[2] != 0.0 {
		t.Errorf("fail to expand")
	}

	vector2 := structures.NewVector([]float64{5.0, 12.0})

	if dot, b := vector1.Dot(vector2); !b || dot != 63.0 {
		t.Errorf("wrong dot product; expected 63.0, got %f", dot)
	}

	projection, b := vector1.Projection(vector2)
	if !b {
		t.Errorf("fail to project")
	} else {
		if vs := projection.V(); vs[0] - 315.0 / 169.0 > 10e-6 || vs[1] - 756.0 / 169.0 > 10e-6 {
			fmt.Println(vs)
			t.Errorf("wrong projection result")
		}
	}

	vector1, _ = vector1.Expand(3)
	vector2, _ = vector2.Expand(3)
	tmp, _ = vector1.CrossProduct(vector2)
	if vs := tmp.V(); vs[0] != 0.0 || vs[1] != 0.0 || vs[2] != 16.0 {
		t.Errorf("wrong crossproduct result")
	}
}
