package leftrotation

import "testing"

type InputArgs struct {
	a []int32
	d int32
}

func TestRotLeft(t *testing.T) {
	cases := []struct {
		input InputArgs
		want  []int32
	}{
		{InputArgs{[]int32{1}, 1}, []int32{1}},
		{InputArgs{[]int32{1, 1}, 1}, []int32{1, 1}},
		{InputArgs{[]int32{1, 2}, 1}, []int32{2, 1}},
		{InputArgs{[]int32{1, 2, 3}, 1}, []int32{2, 3, 1}},
		{InputArgs{[]int32{1, 2, 3, 4, 5, 6}, 4}, []int32{5, 6, 1, 2, 3, 4}},
		{InputArgs{[]int32{1, 2, 3, 4, 5, 6}, 3}, []int32{4, 5, 6, 1, 2, 3}},
	}
	for _, c := range cases {
		got := LeftRotation(c.input.a, c.input.d)
		if len(got) != len(c.want) {
			t.Errorf("LeftRotation(%v) length: %v, want length: %v", c.input, len(c.input.a), len(c.want))
		}
		for i := range c.want {
			if got[i] != c.want[i] {
				t.Errorf("LeftRotation(%v) == %v, want %v", c.input, got, c.want)
			}
		}
	}
}
