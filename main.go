package main

import "fmt"

func main() {
	type x struct {
		x1 interface{}
	}

	var miao x = x{x1: nil}

	fmt.Printf("%#v",miao)


}
