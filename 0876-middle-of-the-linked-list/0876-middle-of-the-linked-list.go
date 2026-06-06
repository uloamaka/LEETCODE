/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    fastP,slowP := head, head

    for fastP != nil && fastP.Next != nil {
        fastP = fastP.Next.Next
        slowP = slowP.Next
    }

    return slowP
}