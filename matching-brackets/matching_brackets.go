package brackets

func Bracket(input string) bool {
	var stack []rune

	for _, r := range input {
		switch r {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case ')', '}', ']':
			n := len(stack) - 1
			if n < 0 || stack[n] != r {
				return false
			}
			stack = stack[:n]
		default:
			continue
		}
	}

	return len(stack) == 0
}
