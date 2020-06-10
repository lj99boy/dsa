package decimaltohex

import "fmt"

func DecimalToHex(num int) (str string) {
	codes := []int{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}

	for i := num; i > 0; i /= 16 {
		code := i % 16
		if code > 9 {
			str = fmt.Sprintf("%v%v", string(codes[code]), str)
		}else{
			str = fmt.Sprintf("%v%v", code, str)
		}
	}
	return str
}
