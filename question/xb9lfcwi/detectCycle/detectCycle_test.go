package detectCycle

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
	tL := &ListNode{Val: 3,Next: &ListNode{Val: 2,Next: &ListNode{Val: 0,Next: &ListNode{Val: 4,Next: nil}}}}
	tL.Next.Next.Next.Next = tL.Next

	println(detectCycle(tL))
}
