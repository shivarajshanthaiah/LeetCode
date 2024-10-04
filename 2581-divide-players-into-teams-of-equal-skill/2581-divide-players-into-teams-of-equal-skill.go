func dividePlayers(skill []int) int64 {

    sort.Ints(skill)
    n := len(skill)
	target := skill[0] + skill[n-1]

	ans := 0
	for i := 0; i < n/2; i++ {
		if skill[i]+skill[n-1 -i] != target {
			return -1
		} else {
			ans += skill[i] * skill[n-1-i]
		}
	}
	return int64(ans)
}