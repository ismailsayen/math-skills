package mathskills

func StandardDeviation(data []int) int {
	return int(Sqrt(Variance(data)))
}

func Sqrt(n int) float64 {
	for i := float64(0); i <= float64(n); i += 0.001 {

		if int(i*i) == n {
			return i
		}
	}
	return 0
}
