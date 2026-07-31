func majorityElement(nums []int) (majority int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // Descending order
	})

	if len(nums) > 0 {
		majority =  nums[len(nums)/2]
	}

	return majority

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
