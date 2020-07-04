package countapplesandoranges

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (int32, int32) {
	var insideApples, insideOranges int32
	for _, v := range apples {
		fall := int32(a + v)
		if s < fall && t > fall {
			insideApples++
		}
	}
	for _, v := range oranges {
		fall := int32(b + v)
		if s < fall && t > fall {
			insideOranges++
		}
	}
	return insideApples, insideOranges
}
