func searchMatrix(matrix [][]int, target int) bool {
	var n int
	n = len(matrix)
	// if n != 0 {
	// 	m = len(matrix[0])
	// }
	
	flatten := []int{}
	for y := 0; y < n; y++ {
		flatten = append(flatten, matrix[y]...)
	}

	return search(0, len(flatten)-1, flatten, target)
}

func search(l int, r int, arr []int, target int) bool {
	if l == r && arr[l] == target {
		return true
	}
	if l > r {
		return false
	}

	mid := (l+r) / 2

	if arr[mid] == target {
		return true
	}
	if arr[mid] < target {
		return search(mid+1, r, arr, target)
	}
	if arr[mid] > target {
		return search(l, mid-1, arr, target)
	}

	return false
}
