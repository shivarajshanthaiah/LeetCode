func singleNumber(nums []int) []int {

	hashMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]]++
	}

	res := []int{}
	for key, val := range hashMap {
		if val == 1 {
			res = append(res, key)
		}
	}
	return res
}