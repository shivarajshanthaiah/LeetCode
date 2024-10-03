func kthCharacter(k int) byte {
	str := []byte{'a'}

	for len(str) < k {
		nxt := []byte{}

		for _, ch := range str {
			nxt = append(nxt, ch)
			nxt = append(nxt, ch+1)
		}
		str = nxt
	}
	return str[k-1]
}