package timeconversion

import "testing"

func TestTimeConversion(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"07:05:45PM", "19:05:45"},
		{"11:05:45AM", "11:05:45"},
		{"12:40:22AM", "00:40:22"},
		{"12:40:22PM", "12:40:22"},
	}
	for _, c := range cases {
		got := TimeConversion(c.input)
		if got != c.want {
			t.Errorf("TimeConversion(%v) == %v, want %v", c.input, got, c.want)
		}
	}
}
