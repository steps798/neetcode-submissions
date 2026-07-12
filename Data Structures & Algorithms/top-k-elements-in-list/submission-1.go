func topKFrequent(nums []int, k int) []int {
	maps := make(map[int]int)
	for _, n := range nums {
		maps[n]++
	}

	// buckets from 0 to len(nums)
	// hence we add+1 to make bucket for zero count
	// 2D array for cater ties bucket
	freq := make([][]int, len(nums)+1)
	for num, count := range maps {
		freq[count] = append(freq[count], num)
	}

	fmt.Printf("%+v", freq)

	var res []int
	// top frequent, walk backward from the highest count bucket
	for i := len(freq) - 1; i > 0; i-- {
		for _, element := range freq[i] {
			res = append(res, element)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
