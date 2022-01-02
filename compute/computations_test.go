package compute

import "testing"

func TestAdd(t *testing.T) {
	x, y := 1, 2
	if Add(x, y) != x+y {
		t.Errorf("Error in Add Computation")
	}
}

func TestTableAdd(t *testing.T) {
	table := []struct {
		x, y, expected int
	}{
		{1, 2, 3},
		{10, 2, 12},
		{3, 3, 6},
		{11, 22, 33},
	}
	for _, v := range table {
		if Add(v.x, v.y) != v.expected {
			t.Errorf("Error in Add Computation")
		}
	}
}

func TestMulti(t *testing.T) {
	x, y := 1, 2
	if Multi(x, y) != x*y {
		t.Errorf("Error in Multi Computation")
	}
}
