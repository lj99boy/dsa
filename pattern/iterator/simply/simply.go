//iterator 最简实现，这样简单但是迭代过程里函数的功能表达不清晰
//https://yourbasic.org/golang/iterator-generator-pattern/
package simply

type Ints []int

func (ints Ints) Do(fn func(n int)) {
	for _, v := range ints {
		fn(v)
	}
}

func (ints Ints) DoWithEscape(fn func(n int) bool) {
	for _, v := range ints {
		if fn(v) {
			break
		}
	}
}
