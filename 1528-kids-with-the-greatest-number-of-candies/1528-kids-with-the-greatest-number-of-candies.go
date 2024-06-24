func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatest := candies[0]
	for _, candy := range candies {
		if candy > greatest {
			greatest = candy
		}
	}
	res := []bool{}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= greatest {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}