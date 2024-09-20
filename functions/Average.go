package mathskills

func Average(data []int) int {
	AVG := 0
	for _, e := range data {
		AVG += e
	}
	return AVG / len(data)
}
