package polynomialadd

import "testing"

func TestPolynomialadd(t *testing.T) {
	left := &chain{
		coef:  3,
		expon: 5,
		next: &chain{
			coef:  -2,
			expon: 3,
			next: &chain{
				coef:  4,
				expon: 0,
				next:  nil,
			},
		},
	}

	right := &chain{
		coef:  3,
		expon: 6,
		next: &chain{
			coef:  4,
			expon: 3,
			next: &chain{
				coef:  2,
				expon: 0,
				next:  nil,
			},
		},
	}

	res := Polynomialadd(left, right)

	for ; res != nil; {
		print(res.coef)
		println(res.expon)
		res = res.next
	}
}
