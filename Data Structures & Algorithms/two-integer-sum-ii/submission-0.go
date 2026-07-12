func twoSum(numbers []int, target int) []int {
	mapNumIdx := make(map[int]int)

	for i, n := range numbers {
		if j, ok := mapNumIdx[target - n]; ok && i != j {
			if i > j {
				return []int{j+1,i+1}
			}
			return []int{i+1,j+1}
		}
		mapNumIdx[n] = i
	}

	return []int{}
}
