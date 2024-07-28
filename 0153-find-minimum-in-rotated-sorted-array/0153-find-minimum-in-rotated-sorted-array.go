func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1

	if nums[left] < nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := (left + right) / 2

		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}