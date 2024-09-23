package mathskills

func Variance(data []int) int {
	VR := 0
	for _, e := range data {
		x := e - Average(data)
		VR += Pow(x)
	}
	return VR / len(data)
}

func Pow(n int) int {
	p := 1
	for i := 1; i <= 2; i++ {
		p = p * n
	}
	return p
}
