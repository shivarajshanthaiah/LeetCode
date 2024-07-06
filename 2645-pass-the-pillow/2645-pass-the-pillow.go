func passThePillow(n int, time int) int {
	if time < n {
		return time+1
	}
    
	increasing := true
	res := 1
	for i := 1; i <= time; i++ {
		if increasing {
			res++
		} else {
			res--
		}
		if res == n {
			increasing = false
		} else if res == 1 {
			increasing = true
		}
	}
	return res
}