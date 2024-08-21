/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	l := 0
	for temp != nil {
		l++
		temp = temp.Next
	}

	l = l - n
	if l == 0 {
		return head.Next
	}

	temp = head
	for i := 0; i < l-1; i++ {
		temp = temp.Next
	}
	if temp.Next != nil {
		temp.Next = temp.Next.Next
	}
    
	return head
}