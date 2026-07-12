func isValid(s string) bool {
	mapBracePairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}
	for _, r := range s {
		if closeBrace, isOpenBrace := mapBracePairs[r]; isOpenBrace {
			stack = append(stack, closeBrace)
		} else {
			if len(stack) == 0 || r != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}