func reverseParentheses(s string) string {
	stack := []rune{}
	cur := []rune{}
	for _, rn := range s {
		if rn != ')' {
			stack = append(stack, rn)
		} else {
			for stack[len(stack)-1] != '(' {
				cur = append(cur, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, cur...)
			cur = cur[:0] // clean up reversal stack
		}
	}
	return string(stack)
}