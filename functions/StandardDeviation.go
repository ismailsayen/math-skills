package mathskills

func StandardDeviation(data []int) int {
	return Sqrt(Variance(data))
}

func Sqrt(n int) int {
	for i := 1; i <= n; i++ {
		if i*i == n {
			return i
		}
	}
	return 0
}
