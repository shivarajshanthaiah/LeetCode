func subarraySum(nums []int, k int) int {
	countFreq := make(map[int]int)
	sum := 0
	count := 0

	for _, num := range nums {
		sum += num
		if sum == k {
			count++
		}
		count += countFreq[sum-k]
		countFreq[sum]++
	}
	return count
}