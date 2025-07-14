/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    current := head
    var concatenatedString string
    for current != nil {
        var builder strings.Builder
    	builder.WriteString(strconv.Itoa(current.Val))
    	concatenatedString += builder.String()
        current = current.Next
    }

    if concatenatedString == "0" {
        return 0
    }

    decimalInt, _ := strconv.ParseInt(concatenatedString, 2, 64)
    return int(decimalInt)
}