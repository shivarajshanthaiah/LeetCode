func kthDistinct(arr []string, k int) string {
	if len(arr) < k {
		return ""
	}

	unique := make(map[string]int)
	for _, s := range arr {
		unique[s]++
	}

	for _, s := range arr {
		if unique[s] == 1 {
			k--
		}
		if k == 0 {
			return s
		}
	}
	return ""
}