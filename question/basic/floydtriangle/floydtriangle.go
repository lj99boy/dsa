//弗洛伊德三角
package floydtriangle

func Floydtriangle(layer int) {
	first := 0
	for i := 1; i <= layer; i++ {
		for blank := 0; blank < layer-i; blank++ {
			print(" ")
		}
		if i == 1 {
			first = 1
		} else {
			first += 1 + (i-2)*1
		}
		tmpFirst := first
		for tmpI := 1; tmpI <= i; tmpI++ {
			print(tmpFirst)
			print(" ")
			tmpFirst++
		}
		println("")
	}
}
