package playground

import (
	"fmt"
	"testing"
)

func TestModifySlice(t *testing.T) {
	s1 := []int{11, 22, 33}
	ModifySlice(s1)
	fmt.Printf("%v", s1[:])
}

func TestAddSlice(t *testing.T) {
	s1 := []int{11, 22, 33}
	AddSlice(&s1)
	fmt.Printf("%v", s1[:])
}
