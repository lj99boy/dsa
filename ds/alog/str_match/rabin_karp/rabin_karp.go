package rabin_karp

import "math"

//const base = 16777619

//ascii字符数量
const base = 128

//func Search(txt string, patterns []string) []string {
//	in := indices(txt, patterns)
//	matches := make([]string, len(in))
//	i := 0
//	for j, p := range patterns {
//		if _, ok := in[j]; ok {
//			matches[i] = p
//			i++
//		}
//	}
//
//	return matches
//}

func IsContain(haystack string, needle string) bool {
	tH := []rune(haystack)
	tN := []rune(needle)

	tNHash := getHash(tN)
	tHHash := getHash(tH[0:len(tN)])

	for i := 0; i <= len(tH)-len(tN)+1; i++ {

		if i != 0 {
			//tHHash -= getHash(append(make([]rune, 0), tH[i-1]))
			tHHash -= getTargetBaseHash(tH[i-1], len(tN)-1)
			tHHash *= base
			//tHHash += getHash(append(make([]rune, 0), tH[i+len(tN)]))
			tHHash += getTargetBaseHash(tH[i-1+len(tN)], 0)
		}

		if tHHash == tNHash {
			return true
		}
	}

	return false
}

//
//func rollingGetHash(txt []rune, start int, length int, oldHash float64) float64 {
//	if oldHash == 0 {
//
//	}
//}

func getTargetBaseHash(char rune, base int) float64 {
	return float64(char) * math.Pow(128, float64(base))
}

func getHash(txt []rune) (res float64) {
	for i := 0; i < len(txt); i++ {
		//这里省略了一个把字符映射到数字的步骤，这里是因为rune默认找到的就是unicode数字，所以可以直接用
		res += float64(txt[len(txt)-i-1]) * math.Pow(128, float64(i))
	}

	return
}
