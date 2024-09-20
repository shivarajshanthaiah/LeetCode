func largestNumber(nums []int) string {
	numStr := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		numStr[i] = strconv.Itoa(nums[i])
	}

	sort := BubbleSort(numStr)
	if sort[0] == "0" {
		return "0"
	}

	ans := strings.Join(sort, "")
	return ans
}

func BubbleSort(str []string) []string {
	n := len(str)
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if str[j]+str[j+1] < str[j+1]+str[j] {
				str[j], str[j+1] = str[j+1], str[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return str
}