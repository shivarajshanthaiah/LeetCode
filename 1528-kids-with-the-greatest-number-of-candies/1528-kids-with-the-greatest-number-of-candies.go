func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := candies[0]
	for _, candy := range candies {
		if candy > greatest {
			greatest = candy
		}
	}
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= greatest {
			res[i] = true
		} else {
			res[i] = false
		}
	}
	return res
}