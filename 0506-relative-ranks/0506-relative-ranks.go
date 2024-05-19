func findRelativeRanks(score []int) []string {
scoreIndex := make(map[int]int)
    for i := 0;i<len(score);i++{
        scoreIndex[score[i]]=i
    }

sortedScore := make([]int, len(score))
copy(sortedScore, score)

sort.Slice(sortedScore, func(i, j int) bool{
    return sortedScore[i] > sortedScore[j]
})

result := make([]string, len(score))

for i, s := range sortedScore{
    idx := scoreIndex[s]
    switch i {
        case 0:
        result[idx] = "Gold Medal"
        case 1 :
        result[idx] = "Silver Medal"
        case 2 :
        result[idx] = "Bronze Medal"
        default :
        result[idx] = strconv.Itoa(i+1)
    }
}

return result

}