func hasDuplicate(nums []int) bool {
    // tampung pakai map dulu
	// map int bool
	// loop tiap nums
	// input map
	// break ketika ketemu di map

	mapNum := make(map[int]bool)
	for _, num := range nums {
		if isFound := mapNum[num]; isFound {
			return true
		}
		mapNum[num] = true
	}

	return false
}
