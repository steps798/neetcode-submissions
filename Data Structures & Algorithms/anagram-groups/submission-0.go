
func groupAnagrams(strs []string) (result [][]string) {
	angrm := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		key := strings.Join(chars, "")
		angrm[key] = append(angrm[key], str)
	}

	for _, list := range angrm {
		result = append(result, list)
	}

	return result
}