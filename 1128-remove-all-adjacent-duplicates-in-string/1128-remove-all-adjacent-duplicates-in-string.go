func removeDuplicates(s string) string {
	stack := []rune{}

	for _, ch := range s {
		n := len(stack)
		if n > 0 && stack[n-1] == ch {
			stack = stack[:n-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}