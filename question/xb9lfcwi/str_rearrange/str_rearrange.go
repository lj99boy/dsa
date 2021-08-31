package str_rearrange

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	holder := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		holder[s1[i:i+1]]++
		holder[s2[i:i+1]]--
	}

	for _, v := range holder {
		if v != 0 {
			return false
		}
	}

	return true
}
