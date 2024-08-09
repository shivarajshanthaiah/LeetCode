func stringMatching(words []string) []string {
	res := []string{}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if j != i {
				if strings.Contains(words[j], words[i]) {
					res = append(res, words[i])
					break
				}
			}
		}
	}
	return res
}