package sortfuncapply

import (
	"fmt"
	"sort"
	"strings"
)

func sortFuncApply() {
	num := []int{50, 90, 30, 10, 50}
	fmt.Println(num)
	if sort.IntsAreSorted(num) == false {
		sort.Ints(num)
	}
	println(num)
	println(sort.SearchInts(num, 30))

	text := []string{"US", "UK", "Germany", "Australia", "Japan"}
	println(text)
	if sort.StringsAreSorted(text) == false {
		sort.Strings(text)
	}
	println(text)
	println(sort.SearchStrings(text, "japan"))

	val := []float64{2.5, 6.3, 1.5, 9.8, 4.7}
	fmt.Println(val)
	if sort.Float64sAreSorted(val) == false {
		sort.Float64s(val)
	}
	println(val)
	println(sort.SearchFloat64s(val, 2.5))

	str := "djwqidjowqdw"
	//strings.Replace()
}
