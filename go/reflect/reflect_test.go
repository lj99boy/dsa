package reflect

import "testing"

func TestExplanation(t *testing.T) {
	Explanation()
}



func TestPrintFieldAndMethod(t *testing.T) {
	var t1 = miao{I1: 123, S1: "ww"}
	PrintFieldAndMethod(t1)
}

func TestModifyByReflect(t *testing.T) {
	ModifyByReflect()
}

func TestCallFunc(t *testing.T) {
	CallFunc()
}
