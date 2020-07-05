package leftrotation

func LeftRotation(a []int32, d int32) []int32 {
	t := a[0:d]
	r := a[d:]
	return append(r, t...)
}
