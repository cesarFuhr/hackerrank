package sockmerchant

import "testing"

type InputArgs struct {
	n  int32
	ar []int32
}

func TestSockMerchant(t *testing.T) {
	cases := []struct {
		input InputArgs
		want  int32
	}{
		{InputArgs{1, []int32{1}}, 0},
		{InputArgs{2, []int32{1, 1}}, 1},
		{InputArgs{3, []int32{1, 1, 1}}, 1},
		{InputArgs{4, []int32{1, 1, 1, 1}}, 2},
		{InputArgs{3, []int32{1, 1, 2}}, 1},
		{InputArgs{3, []int32{1, 2, 3}}, 0},
		{InputArgs{7, []int32{100, 200, 100, 300, 1, 1, 1}}, 2},
		{InputArgs{9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}}, 3},
	}
	for _, c := range cases {
		got := SockMerchant(c.input.n, c.input.ar)
		if got != c.want {
			t.Errorf("SockMerchant(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}
