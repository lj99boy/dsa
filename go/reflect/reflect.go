package reflect

import (
	"fmt"
	"reflect"
)

type tInt int

type miao struct {
	I1 int
	S1 string
}

func (m miao) TMiao() {
	println("this TMiao")
}

func (m miao) MiaoWithArgs(xx string) {
	println(xx)
}

func Explanation() {
	// 变量初始化的时候获得pair(type,value)，在连续断言定义里不会改变
	//file := &os.File{}
	//var r io.Reader
	//r = file
	//var w io.Writer
	//w = r.(io.Writer)
	//var c io.Closer
	//c = w.(io.Closer)
	//var e interface{}
	//e = c
	//tType := reflect.TypeOf(e)
	//tValue := reflect.ValueOf(e)
	//fmt.Printf("t:%v , v:%v", tType, tValue)

	/////////////////////////

	t := tInt(1)
	tType := reflect.TypeOf(t)
	//获得底层类型 reflect.Kind
	tKind := reflect.TypeOf(t).Kind()
	tValue := reflect.ValueOf(t)
	fmt.Printf("type:%v ,kind:%v ,value:%v\n", tType, tKind, tValue)
}

func ModifyByReflect() {
	var xx = 1.2
	pointer := reflect.ValueOf(&xx)
	newVal := pointer.Elem()
	fmt.Printf("%v\n", newVal)

	newVal.SetFloat(55)
	fmt.Printf("%v\n", newVal)

}

func CallFunc() {
	m1 := &miao{}
	m1Val := reflect.ValueOf(m1)
	m1ArgFunc := m1Val.MethodByName("MiaoWithArgs")
	args := []reflect.Value{reflect.ValueOf("test string")}
	if len(args) != m1ArgFunc.Type().NumIn() {
		panic("args num error")
	}
	m1ArgFunc.Call(args)

}

func PrintFieldAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	getValue := reflect.ValueOf(input)

	for i := 0; i < getValue.NumField(); i++ {
		field := getType.Field(i)
		fieldVal := getValue.Field(i).Interface()
		fmt.Printf("name:%s,type:%s,val:%v\n", field.Name, field.Type, fieldVal)
	}

	for i := 0; i < getValue.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}
