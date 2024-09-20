package mathskills

func Median(data []int) int {
	var Date_sorted []int
	Date_sorted = append(Date_sorted, data...)
	Date_sorted = Sort(Date_sorted)
	MD := 0
	if len(Date_sorted)%2 != 0 {
		MD = Date_sorted[len(data)/2]
	} else {
		MD = (Date_sorted[len(Date_sorted)/2] + Date_sorted[((len(Date_sorted)/2)-1)]) / 2
	}
	
	return MD
}
