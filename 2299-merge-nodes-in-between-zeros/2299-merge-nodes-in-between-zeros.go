/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	if head == nil || head.Val == 0 && head.Next == nil {
		return nil
	}
	curr := head
	sum := &ListNode{}
	for curr.Next != nil {
		if curr.Val == 0 {
			sum = curr
			curr = curr.Next
		}
		if curr.Val != 0 {
			sum.Val += curr.Val
			*curr = *curr.Next
		}
	}
	sum.Next = nil
	return head
}