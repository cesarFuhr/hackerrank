package gradingStudents

func round(n int32) int32 {
	if n < 38 {
		return n
	}
	diff := n%5 - 5
	if diff >= -2 {
		return n - diff
	}
	return n
}

func GradingStudents(ar []int32) []int32 {
	var grades = make([]int32, len(ar))
	for i, v := range ar {
		grades[i] = round(v)
	}
	return grades
}
