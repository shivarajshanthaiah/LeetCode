func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && len(strs[i]) < len(prefix) {
			prefix = prefix[:len(prefix)-1]
		}

		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}

		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}