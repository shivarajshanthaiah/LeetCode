func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	var dfs func(keys []int)
	dfs = func(keys []int) {
		for _, key := range keys {
			if visited[key] {
				continue
			}
			visited[key] = true
			dfs(rooms[key])
		}
	}
	dfs([]int{0})
	for _, isVisited := range visited {
		if !isVisited {
			return false
		}
	}
	return true
}