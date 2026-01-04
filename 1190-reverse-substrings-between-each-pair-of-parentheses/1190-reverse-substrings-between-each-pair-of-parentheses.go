func reverseParentheses(s string) string {
	stack := []byte{}
	cur := []byte{}
	for i := range s {
		if s[i] != ')' {
			stack = append(stack, s[i])
		} else {
			for stack[len(stack)-1] != '(' {
				cur = append(cur, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, cur...)
			cur = cur[:0]
		}
	}
	return string(stack)
}