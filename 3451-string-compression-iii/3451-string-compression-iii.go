func compressedString(word string) string {
	comp := []string{}
	count := 1

	for i := 1; i <= len(word); i++ {
		if i == len(word) || word[i] != word[i-1] || count == 9 {
			comp = append(comp, strconv.Itoa(count)+string(word[i-1]))
			count = 1
		} else {
			count++
		}
	}
	return strings.Join(comp, "")
}