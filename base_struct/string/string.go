package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func bytes2str(s []byte) (p string) {
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)

	return p
}




func main() {
	//a := []int{1, 2, 3, 51, 7, 5, 98}
	//fmt.Printf("%#v", append(a[:2], append([]int{0}, a[2:]...)...))
	//a = append(a, 0)
	//copy(a[3:], a[2:])
	//a[2] = 0
	//fmt.Printf("%#v", a)
	a := make([]int, 1)
	g := new([]int)
	g = nil
	g = append(g, 123)
	//ww := *g
	b := &a
	//xx := reflect.TypeOf(b)
	//xx.NumField()
	//yy := reflect.ValueOf(b)
	fmt.Printf("%T---%v\n", nil, nil)
	fmt.Printf("%T---%v\n", new([]int), new([]int))
	fmt.Printf("%T---%v\n", b, b)
	fmt.Printf("%T---%v\n", g, g)
	//fmt.Printf("%T---%v\n", ww[0], ww[0])
	fmt.Printf("%T---%v\n", a[0], a[0])
	//fmt.Printf("%T---%v\n", yy, yy)
	//fmt.Printf("%T---%v\n", reflect.TypeOf(nil).Name(), new([]int))
}
