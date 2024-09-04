func reverseStr(s string, k int) string {
	conv := []rune(s)
	n := len(conv)

	for i := 0; i < n; i += 2 * k {
		l, r := i, min(i+k-1, n-1)
		for l < r {
			conv[l], conv[r] = conv[r], conv[l]
			l++
			r--
		}
	}
	return string(conv)
}