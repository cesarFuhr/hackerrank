package gradingStudents

import "testing"

func TestGradingStudents(t *testing.T) {
	cases := []struct {
		input []int32
		want  []int32
	}{
		{[]int32{1}, []int32{1}},
		{[]int32{1, 33}, []int32{1, 33}},
		{[]int32{1, 43}, []int32{1, 45}},
		{[]int32{1, 57}, []int32{1, 57}},
		{[]int32{77, 65, 64, 99}, []int32{77, 65, 65, 100}},
		{[]int32{40, 39, 38, 26}, []int32{40, 39, 38, 26}},
	}
	for _, c := range cases {
		got := GradingStudents(c.input)
		for i := range c.want {
			if got[i] != c.want[i] {
				t.Errorf("GradingStudents(%v) == %v, want %v", c.input, got, c.want)
			}
		}
	}
}
