package birthdayCakeCandles

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	cases := []struct {
		input []int32
		want  int32
	}{
		{[]int32{1}, 1},
		{[]int32{1, 1}, 2},
		{[]int32{2, 2}, 2},
		{[]int32{1, 3, 2, 2}, 1},
		{[]int32{2, 20, 1000, 100, 1000, 1000, 1000}, 4},
		{[]int32{5, 5, 6, 10, 15, 10, 10, 5, 5, 15, 15, 20}, 1},
	}
	for _, c := range cases {
		got := BirthdayCakeCandles(c.input)
		if got != c.want {
			t.Errorf("BirthdayCakeCandles(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}
