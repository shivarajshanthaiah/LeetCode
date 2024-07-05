func findMaxAverage(nums []int, k int) float64 {
    maxSum := 0
    kSum := 0

    for i:=0;i<k;i++{
        maxSum += nums[i]
    }

    kSum = maxSum
    left := k
    for left < len(nums){
        kSum += nums[left] - nums[left-k]
        if kSum > maxSum{
            maxSum = kSum
        }
        left++
    }
    return float64(maxSum)/ float64(k)
}