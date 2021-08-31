package slide_windows

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	fmt.Printf("%s", MinWindow("ADOBECODEBANC", "ABC"))
	//fmt.Printf("%s", MinWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
	//fmt.Printf("%s", MinWindow("a", "aa"))
}

func TestCheckInclusion(t *testing.T) {
	println(CheckInclusion("ab","eidbaooo"))
}
