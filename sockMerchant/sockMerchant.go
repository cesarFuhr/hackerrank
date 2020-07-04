package sockmerchant

func SockMerchant(n int32, ar []int32) int32 {
	colorAcc := make(map[int32]int32)
	var pairs int32
	for _, c := range ar {
		colorAcc[c]++
		if colorAcc[c]%2 == 0 {
			pairs++
		}
	}
	return pairs
}
