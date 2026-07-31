func majorityElement(nums []int) (majority int) {
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
}
