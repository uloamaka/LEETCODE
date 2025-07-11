/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Create a dummy node to simplify edge cases (e.g., removing head)
    dummy := &ListNode{0, head}
    fast := dummy
    slow := dummy

    // Step 1: Move fast pointer n+1 steps ahead
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // Step 2: Move both pointers until fast hits the end
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    // Step 3: Remove the nth node from the end
    slow.Next = slow.Next.Next

    return dummy.Next
}