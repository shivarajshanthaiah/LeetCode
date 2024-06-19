func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	left := 1
	right := math.MaxInt32

	for left < right {
		mid := left + (right-left)/2
		if canMake(bloomDay, m, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canMake(bloomday []int, m, k, maxDays int) bool {
	bCount := 0
	fCount := 0

	for _, bloom := range bloomday {
		if bloom <= maxDays {
			fCount++
			if fCount == k {
				bCount++
				fCount = 0
				if bCount == m {
					return true
				}
			}
		} else {
			fCount = 0
		}
	}
	return false
}