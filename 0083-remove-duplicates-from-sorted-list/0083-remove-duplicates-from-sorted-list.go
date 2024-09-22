/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    temp := head
    for temp != nil{
        for temp.Next != nil && temp.Next.Val == temp.Val{
            temp.Next = temp.Next.Next
        }
        temp = temp.Next
    }
    return head
}