func calPoints(operations []string) int {
	score := make([]int, 0, len(operations))
	sum := 0

	for _, operation := range operations {
		switch operation {
		case "C":
			sum -= score[len(score)-1]
			score = score[:len(score)-1]
		case "D":
			sum += score[len(score)-1] * 2
			score = append(score, score[len(score)-1]*2)
		case "+":
			sum += score[len(score)-1] + score[len(score)-2]
			score = append(score, score[len(score)-1]+score[len(score)-2])
		default:
			i, _ := strconv.Atoi(operation)
			sum += i
			score = append(score, i)
		}
	}
	return sum
}