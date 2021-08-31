package slide_windows

func MinWindow(s string, t string) string {
	left := 0
	right := 0
	tUnique := make(map[string]int)
	tNeed := make(map[string]int)
	for i := 0; i < len(t); i++ {
		tNeed[string(t[i])]++
	}
	tNum := len(tNeed)
	tExists := 0
	tStart := 0
	minWindowLen := len(s) + len(t) + 1

	for ; right < len(s); right++ {
		if tNeed[string(s[right])] > 0 {
			tUnique[string(s[right])]++
			if tUnique[string(s[right])] == tNeed[string(s[right])] {
				tExists++
			}
		}

		for ; tExists == tNum; {
			if len(s[left:right]) < minWindowLen {
				minWindowLen = len(s[left:right])
				tStart = left
			}

			if tNeed[string(s[left])] > 0 {
				tUnique[string(s[left])]--
				if tUnique[string(s[left])] < tNeed[string(s[left])] {
					tExists--
				}
			}
			left++
		}
	}

	if minWindowLen != len(s)+len(t)+1 {
		return s[tStart : tStart+minWindowLen+1]
	} else {
		return ""
	}
}
