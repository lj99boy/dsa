package tstruct

import "testing"

func TestMiaoStruct_Echo(t *testing.T) {

	var tmp interface{}
	var xx miaoStruct
	//xx = miaoStruct{}
	tmp = xx
	_,e := tmp.(miaoin)
	println(e)
	println(xx.ss)
}