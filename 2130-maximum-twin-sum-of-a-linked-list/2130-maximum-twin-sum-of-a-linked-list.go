/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    stack := []int{}
    current := head
    future := head

    for future != nil && future.Next != nil {
        stack = append(stack, current.Val)
        current = current.Next
        future = future.Next.Next
    }
    
    maxSum := 0
    for current != nil {
        val := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        sum := current.Val + val

        if sum > maxSum {
            maxSum = sum
        }
        current = current.Next
    }
   
    return maxSum
}