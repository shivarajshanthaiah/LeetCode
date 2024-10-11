func findLHS(nums []int) int {
	res := 0
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	for key, val := range count {
		if count[key+1] > 0 {
			res = max(res, val+count[key+1])
		}
	}
	return res
}