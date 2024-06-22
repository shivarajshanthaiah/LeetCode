func numSubarraysWithSum(nums []int, goal int) int {
	countFreq := make(map[int]int)
	countFreq[0] = 1
	sum, count := 0, 0
	for _, num := range nums {
		sum += num
		count += countFreq[sum-goal]
		countFreq[sum]++
	}
	return count
}