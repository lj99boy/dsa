//iterator 标准实现
package regular

type Ints []int

func (ints Ints) Iterator() func() (int, bool) {
	index := 0
	return func() (val int, ok bool) {
		if index < len(ints) {
			val = ints[index]
			ok = true
			index++
			return
		} else {
			return
		}
	}
}
