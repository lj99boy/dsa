package main

import "fmt"

//func test() *int  {
//	var a = 10
//	return &a
//}

//func run() {
//	t := make([]int, 1000, 1000)
//	s := make([]int, 10000, 10000)
//	for i := 0; i < len(s); i++ {
//		s[i]= i
//	}
//}

//func fibonacci() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}
//
//func main() {
//	f := fibonacci()
//	for i := 0; i < 10; i++ {
//		fmt.Printf("fibonacci:%d\n", f())
//	}
//	return
//}

// 本函数测试入口参数和返回值情况
//func dummy(b int) int {
//	// 声明一个c赋值进入参数并返回
//	var c int
//	c = b
//	return c
//}
//// 空函数, 什么也不做
//func void() {
//}
func printStr(s string) {
	fmt.Println(s)
}
func main() {
	var s string
	s = "Escape"
	printStr(s)
	//print(111)
	//println(123)
	//fmt.Printf("123")
	//// 声明a变量并打印
	//var a int
	//// 调用void()函数
	//void()
	//// 打印a变量的值和dummy()函数返回
	//fmt.Println(a, dummy(0))
}
