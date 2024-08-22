func maximum69Number(num int) int {
	str := strconv.Itoa(num)

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '6' {
			runes[i] = '9'
			break
		}
	}
	updated := string(runes)
	res, _ := strconv.Atoi(updated)
	return res
}