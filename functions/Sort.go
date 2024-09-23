package mathskills

func Sort(sl []int) []int {
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
	return sl
}
