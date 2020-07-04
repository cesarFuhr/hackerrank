package hourglasssum

type Matrix [][]int32
type Point struct {
	i int32
	j int32
}

func (m Matrix) CalcHourglassSum(p Point) int32 {
	return m[p.i][p.j] + m[p.i][p.j+1] + m[p.i][p.j+2] + m[p.i+1][p.j+1] + m[p.i+2][p.j] + m[p.i+2][p.j+1] + m[p.i+2][p.j+2]
}

func (m Matrix) FindHourglasses() []Point {
	var points []Point
	var i, j int32
	maxI, maxJ := m.GetMaxI(), m.GetMaxJ()
	for i = 0; i < maxI-2; i++ {
		for j = 0; j < maxJ-2; j++ {
			points = append(points, Point{i, j})
		}
	}
	return points
}

func (m Matrix) GetMaxI() int32 {
	return int32(len(m))
}

func (m Matrix) GetMaxJ() int32 {
	return int32(len(m[0]))
}

func HourglassSum(arr [][]int32) int32 {
	matrix := Matrix(arr)
	hourglasses := matrix.FindHourglasses()
	maxSum := (^(int32(1) << 30) + 1) * 2
	for _, p := range hourglasses {
		sum := matrix.CalcHourglassSum(p)
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum
}
