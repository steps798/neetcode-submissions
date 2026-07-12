func validPalindrome(s string) bool {

	// result = true
	// left = 0
	// right = len-1
	// loop while left != right
	
	// if s[left] != s[right]
	//   if s[left] == s[right-1], right = right-1
	//   else if s[left+1] == s[right], left = left+1
	//   else result = false, break
	// if right - left == 1, break
	
	// end of loop: left = left+1, right = right-1
	// return result

	isPalindrome := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	left := 0
	right := len(s)-1

	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}

		left = left + 1
		right = right - 1
	}

	return true
}
