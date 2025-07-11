/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

    // start from the head node of both list and transverse 
    // check if the currentSum > 10 if yes 
        // take the return the modolus of currentSum,
        // and add 1 to the next currentSum
    // then check the currentSum again if it is > 10.
    // then check the tail node to delete trailing zeros
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{Val: 0}
    current := dummyHead
    carry := 0

    for l1 != nil || l2 != nil || carry > 0 {
        var x, y int
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }

        sum := x + y + carry
        carry = sum / 10

        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }

    return dummyHead.Next
}


  // currentSum := 0
    // newList :=
    // start from the head node of both list and transverse
    // for l1 != nil   {
    //     left := 0
    //     left = l1.Val
    //     fmt.Println(left)
    //     l1 = l1.Next
    // }

    // for l2 != nil   {
    //     right := 0
    //     right = l2.Val
    //     fmt.Println(right)
    //     l2 = l2.Next
    // }
