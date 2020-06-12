package ptr_quote

import "fmt"

func PtrQuote() {
	//var ptrFunc = func(s *string) {
	//	tS := "123"
	//	*s = tS
	//}
	//
	//var quoteFunc = func(s string) {
	//	s = "123"
	//}
	//
	//prtSTest := "miao"
	//println(prtSTest)
	//ptrFunc(&prtSTest)
	//println(prtSTest)
	//
	//quoteTest := "wang"
	//println(quoteTest)
	//quoteFunc(quoteTest)
	//println(quoteTest)

	// 引用是指向其他变量的变量，所以只能修改 不能再指向其他变量，下面的append返回新切片，所以不能影响到原引用变量
	var quoteSFunc = func(s []int) {
		//s[1] = 111
		fmt.Printf("q %p \n", s)
		s = append(s, 444)
		fmt.Printf("q %p \n", s)
	}

	
	var ptrSFunc = func(s *[]int) {
		//ss := *s
		//ss[1] = 111
		fmt.Printf("p %p \n", s)
		*s = append(*s, 444)
		fmt.Printf("p %p \n", s)

	}

	prtSTest := []int{2, 22, 222}
	fmt.Printf("%v \n", prtSTest[:])
	ptrSFunc(&prtSTest)
	fmt.Printf("%v \n", prtSTest[:])

	//quoteTest := []int{2, 22, 222}
	quoteTest := make([]int, 0, 10)
	quoteTest = []int{2, 22, 222}
	fmt.Printf("%v \n", quoteTest[:])
	quoteSFunc(quoteTest)
	fmt.Printf("%v \n", quoteTest[:])
}
