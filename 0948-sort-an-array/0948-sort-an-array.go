func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	mid := (len(nums) / 2)
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {
	combined := make([]int,len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			combined[i+j] = left[i]
            i++
		} else {
			combined[i+j]= right[j]
            j++
		}
	}

    for i < len(left){
        combined[i+j] = left[i]
            i++
    }
        for j< len(right){
        combined[i+j] = right[j]
            j++
    }
	return combined
}