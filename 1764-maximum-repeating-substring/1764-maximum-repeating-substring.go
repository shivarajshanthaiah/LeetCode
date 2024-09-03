func maxRepeating(sequence string, word string) int {

	count := 0
	n := len(sequence)
	m := len(word)
	for i := 0; i <= n-m; i++ {
		currCount := 0
		for j := i; j <= n-m; j +=m  {
			if sequence[j:j+m] == word {
				currCount++
			} else {
				break
			}
		}
		if count < currCount {
			count = currCount
		}
	}
	return count
}