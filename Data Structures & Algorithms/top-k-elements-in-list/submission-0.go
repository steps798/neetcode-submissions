type Element struct {
	num int
	count int 
}
func topKFrequent(nums []int, k int) []int {
	maps := make(map[int]*Element)
	for _, n := range nums {
		if el, ok := maps[n]; ok {
			el.count = el.count + 1 
		} else {
			maps[n] = &Element{num: n, count: 1}
		}
	}

	var els []Element
	for _, el := range maps {
		fmt.Printf("%+v", *el)
		els = append(els, *el)
	}

	sort.Slice(els, func(i, j int) bool{
		return els[i].count > els[j].count
	})

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, els[i].num)
	}

	return res
}
