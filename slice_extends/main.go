package main

import "fmt"

type stack []interface{}

func (s *stack) f() {
	stack := *s
	fmt.Printf("%p %d %d\n", &s, len(stack), cap(stack))

	*s = stack[:len(stack)-1]

	fmt.Printf("%p %d %d\n", &s, len(*s), cap(*s))

	stack = append(stack, "e", "f")
	fmt.Printf("%p %d %d\n", stack, len(stack), cap(stack))
}

func main() {
	var s stack
	s = append(s, "a", "b", "c", "d")
	s.f()
	fmt.Println(s)
	fmt.Printf("%t\n", false)
}
