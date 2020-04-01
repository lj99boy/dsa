package tstruct

import (
	"fmt"
	"testing"
)

func TestSet_Insert(t *testing.T) {
	x := make(map[string]interface{}, 0)
	x["ww"] =12
	x["ww"] =33
	fmt.Printf("%#v",x)
}
