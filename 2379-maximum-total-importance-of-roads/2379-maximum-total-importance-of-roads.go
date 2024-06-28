func maximumImportance(n int, roads [][]int) int64 {
	val := make([]int, n)

	for i := 0; i < len(roads); i++ {
		val[roads[i][0]]++
		val[roads[i][1]]++
	}

	sort.Ints(val)
	var totImp int64 = 0
	for i := 0; i < n; i++ {
		totImp += int64(i+1) * int64(val[i])
	}
	return totImp
}