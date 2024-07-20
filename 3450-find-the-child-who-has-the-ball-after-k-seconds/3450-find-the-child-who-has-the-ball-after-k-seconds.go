func numberOfChild(n int, k int) int {
	if k < n {
		return k
	}
	increasing := true
	res := 1
	for i := 1; i <= k; i++ {
		if increasing {
			res++
		} else {
			res--
		}
		if res == n {
			increasing = false
		}else if res == 1 {
			increasing = true
		}
	}
	return res - 1
}