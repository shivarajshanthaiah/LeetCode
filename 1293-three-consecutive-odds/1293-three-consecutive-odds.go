func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			count++
		} else {
			count = 0
		}
		if count > 2 {
			return true
		}
	}
	return false
}