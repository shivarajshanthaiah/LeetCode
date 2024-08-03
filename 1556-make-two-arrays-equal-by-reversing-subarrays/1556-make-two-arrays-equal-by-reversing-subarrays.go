func canBeEqual(target []int, arr []int) bool {
	if len(target) == 1 {
		return target[0] == arr[0]
	}
	sort.Ints(target)
	sort.Ints(arr)

	return equal(target, arr)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}