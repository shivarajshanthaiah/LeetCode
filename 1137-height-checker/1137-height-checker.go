func heightChecker(heights []int) int {
    newArr := []int{}
    newArr = append(newArr, heights...)
    sort.Ints(newArr)

    i, j := 0, 0
    count := 0
    for i < len(heights) && j < len(newArr){
        if heights[i] != newArr[j]{
            count++
        }
        i++
        j++
    }
    return count
}