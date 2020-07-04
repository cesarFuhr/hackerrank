package birthdayCakeCandles

var m map[int32]int32

func BirthdayCakeCandles(ar []int32) int32 {
	m = make(map[int32]int32)
	for _, v := range ar {
		m[v]++
	}
	var h, r int32
	for k, v := range m {
		if k > h {
			r, h = v, k
		}
	}
	return r
}
