func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)

	return equal(target, arr)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 1 {
		return a[0] == b[0]
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}