/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    temp := head
    var prev *ListNode = nil

    for temp != nil{
        if temp.Next != nil && temp.Val == temp.Next.Val{
            val := temp.Val
            for temp != nil && temp.Val == val{
                temp = temp.Next
            }
            if prev != nil{
                prev.Next = temp
            }else{
                head = temp
            }
        }else{
            prev = temp
            temp = temp.Next
        }
    }
    return head
}