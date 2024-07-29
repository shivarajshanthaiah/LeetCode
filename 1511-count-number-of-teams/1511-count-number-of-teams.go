func numTeams(rating []int) int {
	count := 0
	fir, sec, th := 0, 0, 0
	for i := 0; i < len(rating); i++ {
		fir = rating[i]
		for j := i + 1; j < len(rating); j++ {
			sec = rating[j]
			for k := j + 1; k < len(rating); k++ {
				th = rating[k]
				if fir > sec && sec > th || fir < sec && sec < th {
					count++
				}
			}
		}
	}
	return count
}