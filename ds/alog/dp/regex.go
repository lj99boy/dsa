package dp

import "strconv"

// 实现可以使用.和*的正则匹配
// 如regex(xxy,xxy) || regex(aab,a*b) true  || regex(djwqidowq , .*) true

//dp表 没写完，go实在是写着麻烦
var cache map[string]bool

func init() {
	cache = make(map[string]bool, 0)
}

func Regex(str []rune, pattern []rune) bool {
	ans := false
	i := len(str)
	j := len(pattern)

	if len(str) == 0 {
		return len(pattern) == 2 && pattern[1] == '*'
	}

	currentMatch := false
	if pattern[0] == str[0] || pattern[0] == '.' {
		currentMatch = true
	}

	if len(pattern) == 1 {
		return currentMatch && len(str) == 1
	}

	if pattern[1] == '*' {
		if len(pattern) == 2 {
			if v, ok := cache[strconv.Itoa(i+1)+"_"+strconv.Itoa(j)]; ok {
				return v
			}
			ans = currentMatch && Regex(str[1:], pattern)
			cache[strconv.Itoa(i+1)+"_"+strconv.Itoa(len(pattern))] = ans
		} else {
			//dp表 没写完，go实在是写着麻烦
			//if v, ok := cache[strconv.Itoa(i+1)+"_"+strconv.Itoa(j+2)]; ok {
			//	return
			//}
			ans = Regex(str[1:], pattern[2:]) || (currentMatch && Regex(str[1:], pattern))
		}
	} else {
		ans = currentMatch && Regex(str[1:], pattern[1:])
	}
	return ans
}
