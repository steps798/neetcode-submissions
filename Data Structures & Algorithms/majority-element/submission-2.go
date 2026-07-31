func majorityElement(nums []int) (majority int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // Descending order
	})

	if len(nums) > 0 {
		majority =  nums[len(nums)/2]
	}

	return majority
}
