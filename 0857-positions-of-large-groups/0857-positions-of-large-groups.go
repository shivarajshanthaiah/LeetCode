func largeGroupPositions(s string) [][]int {
	res := [][]int{}
	for i := 0; i < len(s); i++ {
		j := i

		for j < len(s) {
			if s[i] != s[j] {
				break
			}
			j++
		}
		if j-i >= 3 {
			res = append(res, []int{i, j - 1})
		}
		i = j - 1
	}
	return res
}