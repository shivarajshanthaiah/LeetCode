func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	var jobs [][]int
	for i := 0; i < len(difficulty); i++ {
		jobs = append(jobs, []int{difficulty[i], profit[i]})
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	sort.Ints(worker)
	count, temp, maxProf := 0, 0, 0

	for i := 0; i < len(worker); i++ {
		for temp < len(jobs) && jobs[temp][0] <= worker[i] {
			if jobs[temp][1] > maxProf {
				maxProf = jobs[temp][1]
			}
			temp++
		}
		count += maxProf
	}
	return count
}