package countapplesandoranges

import "testing"

func TestCountApplesAndOranges(t *testing.T) {
	cases := []struct {
		input struct {
			s       int32
			t       int32
			a       int32
			b       int32
			apples  []int32
			oranges []int32
		}
		want struct {
			apples  int32
			oranges int32
		}
	}{
		{7, 10, 4, 12, []int{2, 3, -4}, []int{3, -2, -4}},
	}
	for _, c := range cases {
		got := CountApplesAndOranges(c.input...)
		if got != c.want {
			t.Errorf("CountApplesAndOranges(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}
