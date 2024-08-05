func kthDistinct(arr []string, k int) string {
	if len(arr) < k {
		return ""
	}

	unique := make(map[string]int)
	for _, s := range arr {
		unique[s]++
	}

	dCount := 0
	for _, s := range arr {
		if unique[s] == 1 {
			dCount++
		}
		if dCount == k {
			return s
		}
	}
	return ""
}