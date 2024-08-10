func reformat(s string) string {
	letters := []rune{}
	digits := []rune{}

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			letters = append(letters, ch)
		} else {
			digits = append(digits, ch)
		}
	}

	if abs(len(letters)-len(digits)) > 1 {
		return ""
	}

	res := []rune{}
	if len(letters) > len(digits) {
		res = merge(letters, digits)
	} else {
		res = merge(digits, letters)
	}
	return string(res)
}

func merge(left, right []rune) []rune {
	res := []rune{}
	for i := 0; i < len(left) || i < len(right); i++ {
		if i < len(left) {
			res = append(res, left[i])
		}
		if i < len(right) {
			res = append(res, right[i])
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}