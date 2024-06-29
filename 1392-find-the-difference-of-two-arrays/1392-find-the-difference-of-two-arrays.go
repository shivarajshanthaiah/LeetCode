func findDifference(nums1 []int, nums2 []int) [][]int {
	unique1 := make(map[int]bool)
	unique2 := make(map[int]bool)
	uniqueNums1 := []int{}
	uniqueNums2 := []int{}

	for _, num := range nums1 {
		unique1[num] = true
	}

	for _, num := range nums2 {
		unique2[num] = true
	}

	for key, _ := range unique1 {
		_, ok := unique2[key]
		if !ok {
			uniqueNums1 = append(uniqueNums1, key)
		}
	}

	for key, _ := range unique2 {
		_, ok := unique1[key]
		if !ok {
			uniqueNums2 = append(uniqueNums2, key)
		}
	}
	return [][]int{uniqueNums1, uniqueNums2}
}