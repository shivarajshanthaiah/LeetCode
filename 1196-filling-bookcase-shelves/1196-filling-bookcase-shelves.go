func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	ans := make([]int, n+1)

	for i := range ans {
		ans[i] = math.MaxInt32
	}
	ans[0] = 0
	for i := 1; i <= n; i++ {
		width := 0
		height := 0

		for j := i; j > 0; j-- {
			width += books[j-1][0]
			if width > shelfWidth {
				break
			}
			if books[j-1][1] > height {
				height = books[j-1][1]
			}
			ans[i] = min(ans[i], ans[j-1]+height)
		}
	}
	return ans[n]
}