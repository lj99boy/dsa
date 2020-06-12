package writealternately

//func findStr(needle string, haystack string) int {
//	tNeedle := []rune(needle)
//	tHaystack := []rune(haystack)
//	tag := -1
//
//	inRes := false
//	for i := 0; i <= len(tNeedle); i++ {
//		tNeedle = tNeedle[i:]
//		tag = i
//		for j := 0; j <= len(tHaystack); j++ {
//			tHaystack = tHaystack[j:]
//			if len(tNeedle) > len(tHaystack) {
//				return -1
//			}
//
//			if tNeedle[i] != tHaystack[j] {
//				break
//			}
//		}
//	}
//
//	return tag
//}

func findStr(needle string, haystack string) int {
	tNeedle := []rune(needle)
	tHaystack := []rune(haystack)
	tag := -1
	for i := 0; i <= len(haystack); i++ {
		tag = i
		if inMatch(tNeedle, tHaystack) {
			return tag
		}
	}
	return -1
}

func inMatch(left []rune, right []rune) bool {
	if len(left) > len(right) {
		return false
	}
	for i := 0; i <= len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
