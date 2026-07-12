func twoSum(nums []int, target int) []int {
    mapNumIdx := make(map[int]int, 0)
    for i, n := range nums {
        mapNumIdx[n] = i
    }

    for i, n := range nums {
        diff := target-n
        if j, ok := mapNumIdx[diff]; ok {
            if i == j {
                continue
            }
            return []int{i, j}
        }
    }

    return []int{}
}
