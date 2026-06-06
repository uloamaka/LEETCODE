/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    fastP,slowP := head, head
    cnt := 0

    for fastP != nil {
        cnt++
        fastP = fastP.Next
    }

    middle := cnt / 2
    for i := 0; i < middle; i++ {
        slowP = slowP.Next
    }

    return slowP
}