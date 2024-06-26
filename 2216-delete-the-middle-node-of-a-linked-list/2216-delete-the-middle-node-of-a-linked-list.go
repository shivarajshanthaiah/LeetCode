/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head

	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	prev.Next = slow.Next
	return head
}