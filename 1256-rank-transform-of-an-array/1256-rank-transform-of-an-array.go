func arrayRankTransform(arr []int) []int {
	newArr := append([]int{}, arr...)
	sort.Ints(newArr)

	rm := removeDuplicates(newArr)

	if rm > 0 {
		newArr = newArr[:rm]
	}

	rank := make(map[int]int)
	for i, num := range newArr {
		rank[num] = i + 1
	}

	res := make([]int, len(arr))
	for i, num := range arr {
		res[i] = rank[num]
	}
	return res
}

func removeDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}
	return j + 1
}