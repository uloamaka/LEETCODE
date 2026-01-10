func countCollisions(directions string) int {
	stack := []rune{}
	collisions := 0

	for _, car := range directions {
		if car == 'R' {
			stack = append(stack, 'R')
			continue
		}

		if car == 'S' {
			for len(stack) > 0 && stack[len(stack)-1] == 'R' {
				stack = stack[:len(stack)-1]
				collisions++
			}
			if len(stack) == 0 || stack[len(stack)-1] != 'S' {
				stack = append(stack, 'S')
			}
			continue
		}

		if car == 'L' {
			if len(stack) > 0 && stack[len(stack)-1] == 'R' {
				stack = stack[:len(stack)-1]
				collisions += 2

				for len(stack) > 0 && stack[len(stack)-1] == 'R' {
					stack = stack[:len(stack)-1]
					collisions++
				}
				stack = append(stack, 'S')
			} else if len(stack) > 0 && stack[len(stack)-1] == 'S' {
				// L hits stopped car
				collisions++
			}
		}
	}

	return collisions
}
