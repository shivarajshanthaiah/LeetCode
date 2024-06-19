/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	eHead := head.Next
	oddHead := head
	evenHead := head.Next
	for oddHead.Next != nil && evenHead.Next != nil {
		oddHead.Next = evenHead.Next
		oddHead = oddHead.Next
		evenHead.Next = oddHead.Next
		evenHead = evenHead.Next
	}
	oddHead.Next = eHead
	return head
}