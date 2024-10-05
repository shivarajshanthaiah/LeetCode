func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)

	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}

	compareMaps := func(m1, m2 map[byte]int) bool {
		for k, v := range m1 {
			if m2[k] != v {
				return false
			}
		}
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		if compareMaps(s1Map, s2Map) {
			return true
		}
		s2Map[s2[i]-'a']++
		s2Map[s2[i-len(s1)]-'a']--
	}

	return compareMaps(s1Map, s2Map)
}