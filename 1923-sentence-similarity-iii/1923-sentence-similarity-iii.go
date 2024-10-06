func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence1) < len(sentence2) {
		sentence1, sentence2 = sentence2, sentence1
	}

	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	i, j := 0, 0

	for i < len(words2) && words1[i] == words2[i] {
		i++
	}

	for j < len(words2) && words1[len(words1)-j-1] == words2[len(words2)-j-1] {
		j++
	}

	return i+j >= len(words2)
}