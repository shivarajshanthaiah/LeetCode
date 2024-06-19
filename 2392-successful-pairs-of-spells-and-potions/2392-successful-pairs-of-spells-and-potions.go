func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := []int{}
	for i := 0; i < len(spells); i++ {
		left, right := 0, len(potions)-1
		for left <= right {
			mid := (left + right) / 2
			if int64(spells[i]* potions[mid]) < success {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		res = append(res, len(potions)-left)
	}
	return res
}