func construct2DArray(original []int, m int, n int) [][]int {

	if len(original) != m*n {
		return [][]int{}
	}

	res := make([][]int, m)

	for i := 0; i < m; i++ {
		start := i * n
		end := start + n
		res[i] = original[start:end]
	}
	return res
}