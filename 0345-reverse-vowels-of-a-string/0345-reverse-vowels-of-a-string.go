func reverseVowels(s string) string {
	convert := []rune(s)
	vow := make(map[rune]bool)
	vow['a'], vow['A'] = true, true
	vow['e'], vow['E'] = true, true
	vow['i'], vow['I'] = true, true
	vow['o'], vow['O'] = true, true
	vow['u'], vow['U'] = true, true

	start, end := 0, len(convert)-1
	for start < end {
		for start < end && !vow[convert[start]] {
			start++
		}
		for start < end && !vow[convert[end]] {
			end--
		}
		convert[start], convert[end] = convert[end], convert[start]
		start++
		end--
	}
	return string(convert)
}