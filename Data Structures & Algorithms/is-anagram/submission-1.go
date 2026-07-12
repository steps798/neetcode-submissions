func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapCharCountS := make(map[byte]int)
	mapCharCountT := make(map[byte]int)

	for _, c := range []byte(s) {
		mapCharCountS[c] = mapCharCountS[c] + 1
	}

	for _, c := range []byte(t) {
		mapCharCountT[c] = mapCharCountT[c] + 1
	}

	isAnagram := true
	for charS, countS := range mapCharCountS {
		if countT, ok := mapCharCountT[charS]; !ok || countT != countS {
			isAnagram = false
		}
	}

	return isAnagram
}
