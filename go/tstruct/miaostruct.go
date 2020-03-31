package tstruct

type miaoin interface {
	Echo()
}

type miaoStruct struct {
	ss int
}

func (m *miaoStruct) Echo() {
	println("miao")
}

func (m *miaoStruct) EEE() {
	println(m.ss)
}
