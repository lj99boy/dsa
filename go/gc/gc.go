// gc 伪码解释 https://mp.weixin.qq.com/s/G7id1bNt9QpAvLe7tmRAGw
package gc

//表示程序里的各个引用，根引用可能更接近程序的入口
type Refer struct {
	next    *Refer
	str     string
	isRoot  bool
	inStack bool
}

var holder []*Refer

// 扫描整个堆和栈
func ScanAndCollect() {
	//连贯路线1，和程序联通的路线
	path1 := &Refer{
		str: "对象1",

		next: &Refer{
			str: "对象2",
		},
	}
	// 落单路线
	holder = append(holder, path1)

}
