func occurrencesOfElement(nums []int, queries []int, x int) []int {
	occurence := []int{}
	count := 0
	for i, num := range nums {
		if num == x {
			occurence = append(occurence, i)
			count++
		}
	}
	ans := make([]int, len(queries))
	for idx, query := range queries {
		if query > count {
			ans[idx] = -1
		} else {
			ans[idx] = occurence[query-1]
		}
	}
	return ans
}