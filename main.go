package main

import "fmt"

func main() {
	type x struct {
		x1 interface{}
	}

	var miao x = x{x1: nil}

	fmt.Printf("%#v\n",miao)

	mm := []int{7}
	fmt.Println(len(mm[1:]))

}
