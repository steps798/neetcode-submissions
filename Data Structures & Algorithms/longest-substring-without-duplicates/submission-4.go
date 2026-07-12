func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	l, longest := 0, 0
	subs := make(map[rune]int)
	for i, r := range s {
		if prev, ok := subs[r]; ok {
			l = max(prev + 1, l)
		}
		subs[r] = i
		if longest < i - l + 1 {
			longest = i - l + 1 
		}
	}

	return longest
}
