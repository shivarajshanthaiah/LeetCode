/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	temp := head
	arr := []int{}
	for temp != nil {
		arr = append(arr, temp.Val)
		temp = temp.Next
	}
	ans := ispal(arr)
	return ans
}

func ispal(arr []int) bool {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}
	return true
}