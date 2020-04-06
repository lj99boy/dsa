//polynomialadd 多项式加法
package polynomialadd

type chain struct {
	coef  float64
	expon float64
	next  *chain
}

func Polynomialadd(left, right *chain) *chain {
	holder := &chain{}

	for ; left != nil && right != nil; {
		if left.expon > right.expon {
			attach(left.coef, left.expon, holder)
			left = left.next
		} else if left.expon < right.expon {
			attach(right.coef, right.expon, holder)
			right = right.next
		} else {
			expon := left.expon
			coef := left.coef + right.coef
			attach(coef, expon, holder)
			left = left.next
			right = right.next
		}
	}
	for ; left != nil; {
		attach(left.coef, left.expon, holder)
		left = left.next
	}
	for ; right != nil; {
		attach(right.coef, right.expon, holder)
		right = right.next
	}

	return holder.next
}

func attach(coef, expon float64, holder *chain) {
	tmp := &chain{
		coef:  coef,
		expon: expon,
	}
	point := holder
	for ; point.next != nil; {
		point = point.next
	}

	point.next = tmp
}
