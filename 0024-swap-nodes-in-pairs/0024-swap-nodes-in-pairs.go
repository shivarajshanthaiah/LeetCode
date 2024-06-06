/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	temp := &ListNode{Next: head}
	prev := temp

	for prev.Next != nil && prev.Next.Next != nil {
		firstNode := prev.Next
		secondNode := prev.Next.Next

		prev.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		prev = firstNode
	}
	return temp.Next
}