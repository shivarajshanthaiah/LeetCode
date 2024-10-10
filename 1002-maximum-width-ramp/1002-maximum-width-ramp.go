func maxWidthRamp(nums []int) int {
	stack := []int{}

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	ans := 0
	for j := len(nums) - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
			ans = max(ans, j-stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}