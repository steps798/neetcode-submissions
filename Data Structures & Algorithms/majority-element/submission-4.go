func majorityElement(nums []int) (majority int) {
	// 3rd Boyer-Moore Voting Algorithm
	var candidate, count int
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate

	// 2nd submission
	/*
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // Descending order
	})

	if len(nums) > 0 {
		majority =  nums[len(nums)/2]
	}
	return majority
	*/
	// 1st submission
	/*
    count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	fmt.Printf("%+v", count)
	var max int 
	for n, c := range count {
		if max < c {
			majority = n
			max = c
		}
	}
	return majority
	*/
}
