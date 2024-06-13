func mergeAlternately(word1 string, word2 string) string {
	res := []byte{}
	i := 0

	for i < len(word1) && i < len(word2) {
		res = append(res, word1[i], word2[i])
		i++
	}
	res = append(res, word1[i:]...)
	res = append(res, word2[i:]...)
	return string(res)
}