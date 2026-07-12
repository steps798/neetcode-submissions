type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var sizes string
	var joinStr string
	for _, s := range strs {
		if len(sizes) > 0 {
			sizes += ","
		}
		sizes += strconv.Itoa(len(s))
		joinStr += s
		fmt.Println(s)
	}

	result := sizes + "#" + joinStr
	fmt.Println(result)
	return result
}

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 {
		return []string{}
	}

	encodedSplits := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(encodedSplits[0], ",")

	var results []string
	var latestIdx int
	for _, s := range sizes {
		size, _ := strconv.Atoi(s)
		result := encodedSplits[1][latestIdx:latestIdx+size]
		results = append(results, result)
		latestIdx += size
	}

	return results
}
