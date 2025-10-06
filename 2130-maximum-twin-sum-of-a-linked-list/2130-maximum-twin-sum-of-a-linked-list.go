/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    list := []int{}
    curr := head
    for curr != nil {
        list = append(list, curr.Val)
        curr = curr.Next
    } 
    return getMax(list)
}
func getMax(list []int) int {
    max := 0
    l, r := 0, len(list)-1
    for l < r {
        sum := list[l]+list[r]
        if sum > max {
            max = sum
        }
        r--
        l++
    }
    return max
}