func findLucky(arr []int) int {
	freq := make(map[int]int)

	for _, nums := range arr {
		freq[nums]++
	}
	sort.Ints(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == freq[arr[i]] {
			return arr[i]
		}
	}
	return -1
}