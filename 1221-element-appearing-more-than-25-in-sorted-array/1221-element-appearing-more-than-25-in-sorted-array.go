func findSpecialInteger(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	curr25 := len(arr) / 4

	for i := 0; i < len(arr); {
		value := arr[i]
		j := i
		for j < len(arr) && arr[j] == value {
			j++
		}

		count := j - i

		if count > curr25 {
			return value
		}
		i = j
	}
	return -1
}