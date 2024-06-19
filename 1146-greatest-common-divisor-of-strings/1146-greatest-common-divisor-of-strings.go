func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:pre(len(str1), len(str2))]
}

func pre(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}