func search(nums []int, target int) int {
	// sorted
	return split(0, len(nums)-1, nums, target)
}

func split(l int, r int, nums []int, target int) int {
	if l == r && nums[l] == target {
		return l
	}
	if l > r {
		return -1
	}

	mid := (l+r)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return split(mid+1, r, nums, target)
	}

	if nums[mid] > target {
		return split(l, mid-1, nums, target)
	}

	return -1
}
