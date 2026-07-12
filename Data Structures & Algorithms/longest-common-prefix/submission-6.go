func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

    stack := []rune(strs[0])
	 
	for i := 1; i < len(strs); i++ {
		if len(stack) == 0 || len(strs[i]) == 0 {
			return ""
		}
		for j, r := range stack {
			if j >= len(strs[i]) || r != []rune(strs[i])[j] {
				stack = stack[:j]
				break
			}
		}
	}

	return string(stack)
}
