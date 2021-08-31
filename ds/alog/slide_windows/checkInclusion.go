package slide_windows

func CheckInclusion(s1 string, s2 string) bool {
	left := 0
	right := 0
	tUnique := make(map[string]int)
	tNeed := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		tNeed[string(s1[i])]++
	}
	tNum := len(tNeed)
	tExists := 0
	for ; right < len(s2); right++ {
		if tNeed[string(s2[right])] > 0 {
			tUnique[string(s2[right])]++
			if tUnique[string(s2[right])] == tNeed[string(s2[right])] {
				tExists++
			}
		}

		for ; tExists == tNum; {

			if tNeed[string(s2[left])] > 0 {
				if len(s2[left:right+1]) == len(s1) {
					return true
				}
				tUnique[string(s2[left])]--
				if tUnique[string(s2[left])] < tNeed[string(s2[left])] {
					tExists--
				}
			}
			left++
		}
	}

	return false
}
