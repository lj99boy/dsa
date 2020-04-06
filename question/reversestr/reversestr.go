package reversestr

func ReverseStr(s string) (string, bool) {
	str := []rune(s)
	length := len(str)
	if length > 5000 {
		return string(str), false
	}

	for i := 0; i < length/2; i++ {
		str[i], str[length-i-1] = str[length-i-1], str[i]
	}
	return string(str), true
}
