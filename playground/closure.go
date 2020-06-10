package playground

func TClosure() func() int {
	var res = 2
	return func() int {
		res = res * 2
		return res
	}
}
