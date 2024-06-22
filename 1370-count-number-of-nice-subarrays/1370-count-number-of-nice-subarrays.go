func numberOfSubarrays(nums []int, k int) int {
	sumFreq := map[int]int{0:1}
	sum, res := 0, 0
	for _, num := range nums {
		sum += num % 2
		res += sumFreq[sum-k]
		sumFreq[sum]++
	}
	return res
}