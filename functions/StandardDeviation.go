package mathskills

import "math"

func StandardDeviation(data []int) int {
	return int(math.Sqrt(float64(Variance(data))))
}
