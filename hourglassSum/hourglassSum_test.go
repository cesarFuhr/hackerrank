package hourglasssum

import "testing"

func TestHourglassSum(t *testing.T) {
	cases := []struct {
		input [][]int32
		want  int32
	}{
		{[][]int32{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
		}, 7},
		{[][]int32{
			{2, 1, 5},
			{8, 1, 5},
			{1, -1, -1},
		}, 8},
	}
	for _, c := range cases {
		got := HourglassSum(c.input)
		if got != c.want {
			t.Errorf("HourglassSum(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}

func TestGetMaxI(t *testing.T) {
	cases := []struct {
		input [][]int32
		want  int32
	}{
		{[][]int32{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
		}, 3},
		{[][]int32{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
		}, 3},
		{[][]int32{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
		}, 5},
		{[][]int32{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
		}, 8},
	}
	for _, c := range cases {
		m := Matrix(c.input)
		got := m.GetMaxI()
		if got != c.want {
			t.Errorf("Matrix.GetMaxI(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}

func TestGetMaxJ(t *testing.T) {
	cases := []struct {
		input [][]int32
		want  int32
	}{
		{[][]int32{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
		}, 3},
		{[][]int32{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
		}, 7},
	}
	for _, c := range cases {
		m := Matrix(c.input)
		got := m.GetMaxJ()
		if got != c.want {
			t.Errorf("Matrix.GetMaxJ(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}

func TestFindHourglasses(t *testing.T) {
	cases := []struct {
		input [][]int32
		want  []Point
	}{
		{[][]int32{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
		}, []Point{{0, 0}},
		},
		{[][]int32{
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 1},
			{0, 1, 0, 1, 1},
		}, []Point{{0, 0}, {0, 1}, {0, 2}},
		},
		{[][]int32{
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
		}, []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}},
		},
		{[][]int32{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}, []Point{{0, 0}, {1, 0}, {2, 0}},
		},
		{[][]int32{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}, []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
		},
		{[][]int32{
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1},
		}, []Point{
			{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6},
			{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
			{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6},
			{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5}, {3, 6},
			{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5}, {4, 6},
		},
		},
	}
	for _, c := range cases {
		m := Matrix(c.input)
		got := m.FindHourglasses()
		for i := range c.want {
			if got[i] != c.want[i] {
				t.Errorf("Matrix.FindHourglasses(%v) == %v, want %v", c.input, got, c.want)
			}
		}
	}
}