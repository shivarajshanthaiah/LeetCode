func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2freq := make([]int, 1001)

	for _, num := range arr2 {
		arr2freq[num]++
	}

	count := make([]int, 1001)
	remaining := []int{}

	for _, num := range arr1 {
		if arr2freq[num] == 0 {
			remaining = append(remaining, num)
		} else {
			count[num]++
		}
	}

	sort.Ints(remaining)

	res := []int{}
	for _, num := range arr2 {
		for i := 0; i < count[num]; i++ {
			res = append(res, num)
		}
	}
	res = append(res, remaining...)
	return res
}