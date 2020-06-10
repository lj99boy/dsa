package playground

//go变量传递分三种，值传递 引用传递 地址（指针）传递
//引用传递包括 map slice channel interface 类型的值，这种传递只能对modify生效，要想记录对传入值的所有变动，需要使用地址传递
//值传递是除以上四种值以外的所有类型值传递

func ModifySlice(s1 []int) {
	s1[0] = 100
}

func AddSlice(s1 *[]int) {
	*s1 = append(*s1, 100)
}
