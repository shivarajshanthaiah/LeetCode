func canArrange(arr []int, k int) bool {
	reminder := make(map[int]int)

	for _, num := range arr {
		rem := ((num % k) + k) % k
		reminder[rem]++
	}

	for rem, count := range reminder {
		if rem == 0 {
			if count%2 != 0 {
				return false
			}
		} else {
			comp := k - rem
			if reminder[comp] != count {
				return false
			}
		}
	}
	return true
}