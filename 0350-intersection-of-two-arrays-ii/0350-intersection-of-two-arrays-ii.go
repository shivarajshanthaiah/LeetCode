func intersect(nums1 []int, nums2 []int) []int {
	duplicate := make(map[int]int)
	result := []int{}

	for _, num := range nums1 {
		duplicate[num]++
	}

	for _, num := range nums2 {
		if count, found := duplicate[num]; found && count > 0 {
			result = append(result, num)
			duplicate[num]--
		}
	}

	return result
}