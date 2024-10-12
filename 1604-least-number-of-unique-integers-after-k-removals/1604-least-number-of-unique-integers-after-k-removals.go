func findLeastNumOfUniqueInts(arr []int, k int) int {
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}

	counts := []int{}
	for _, val := range count {
		counts = append(counts, val)
	}
	sort.Ints(counts)

	unique := len(counts)
	for _, c := range counts {
		if k >= c {
			k -= c
			unique--
		} else {
			break
		}
	}
	return unique
}