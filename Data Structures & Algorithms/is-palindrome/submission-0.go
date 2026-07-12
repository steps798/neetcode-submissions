func isPalindrome(s string) bool {
	// s dipecah jadi bytes of char
	// di loop tiap bytes
	// tampung array 
	// clean up non-alphanumeric, space, lower case
	// int len(array)/2
	// loop 0 - len(array)/2 // 3
	// compare array[0] == bytes[len(array)-1]
	// compare array[1] == bytes[len(array)-1-1]
	// compare array[i] == bytes[len(array)-1-i]
	// if array[i] != bytes[len(array)-1-i] return false

	// tabacat
	// 0123456

	// return true

	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	cleanup := regex.ReplaceAllString(s, "")
	cleanup = strings.ToLower(cleanup)

	for i := 0; i < len(cleanup)/2; i++ {
		if cleanup[i] != cleanup[len(cleanup)-1-i] {
			return false
		}
	}

	return true
}
