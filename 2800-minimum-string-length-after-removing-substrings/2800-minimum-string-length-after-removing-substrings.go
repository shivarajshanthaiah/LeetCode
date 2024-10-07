func minLength(s string) int {
	stack := []rune{}

	for _, r := range s {
		if len(stack) > 0 {
			last := stack[len(stack)-1]

			if (last == 'A' && r == 'B') || (last == 'C' && r == 'D') {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, r)
	}
	return len(stack)
}