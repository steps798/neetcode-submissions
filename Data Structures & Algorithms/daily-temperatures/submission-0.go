func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		count := 0
		warmerFound := false
		if i != len(temperatures) - 1 {
			for j := i+1; j < len(temperatures); j++ {
				count++
				if temperatures[j] > temperatures[i] {
					warmerFound = true
					break
				}
			}
		}
		if warmerFound {
			result[i] = count
			warmerFound = false
		} else {
			result[i] = 0
		}

	}

	return result
}
