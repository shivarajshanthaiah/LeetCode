func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combination := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ans := []string{""}

	for len(digits) > 0 {
		curNum := digits[0]
		lenRes := len(ans)

		for _, val := range combination[rune(curNum)] {
			for i := 0; i < lenRes; i++ {
				ans = append(ans, ans[i]+string(val))
			}
		}
		ans = ans[lenRes:]
		digits = digits[1:]
	}
	return ans
}