//isarmstrongnum 是否阿姆斯特朗数
package isarmstrongnum

import "math"

func IsArmStrongNum(num int) bool {
	rawNum := num
	sum := 0
	for ; num != 0; {
		tmp := num % 10
		sum += int(math.Pow(float64(tmp), float64(3)))
		num = (num - tmp) / 10
	}

	return sum == rawNum
}
