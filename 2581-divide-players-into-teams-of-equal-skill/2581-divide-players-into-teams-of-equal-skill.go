func dividePlayers(skill []int) int64 {
	sum := 0
	for _, v := range skill {
		sum += v
	}

	if (sum*2) % len(skill) != 0 {
		return -1
	}

	target := (sum * 2) / len(skill)
	sort.Ints(skill)

	ans := 0
	for i := 0; i < len(skill)/2; i++ {
		if skill[i]+skill[len(skill)-1 -i] != target {
			return -1
		} else {
			ans += skill[i] * skill[len(skill)-1-i]
		}
	}
	return int64(ans)
}