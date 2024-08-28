func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	// for _, ch := range s {
	// 	map1[ch]++
	// }

	// for _, ch := range t {
	// 	map2[ch]++
	// }

	for i := 0; i < len(s); i++ {
		map1[rune(s[i])]++
		map2[rune(t[i])]++
	}

	for key, _ := range map1 {
		if map1[key] != map2[key] {
			return false
		}
	}
	return true
}