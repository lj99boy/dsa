//打印金字塔
package pyramid

func printPyramid(num int) {

	for row := 1; row <= num; row++ {
		for i := 1; i <= num-row; i++ {
			print(" ")
		}

		for i := 1; i <= (1 + (row-1)*2); i++ {
			print("*")
		}
		println(" ")
	}
}
