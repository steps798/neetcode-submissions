func getConcatenation(nums []int) (ans []int) {
    ans = nums
	ans = append(ans, nums...)
	return ans
}
