package removeDuplicateNodes

import (
	"fmt"
	"testing"
)

func Test_removeDuplicateNodes(t *testing.T) {
	xx := &ListNode{Val: 1,Next: &ListNode{Val: 2,Next: &ListNode{Val: 3,Next: &ListNode{Val: 3,Next: &ListNode{Val: 2,Next: &ListNode{Val: 1,Next: nil}}}}}}

	removeDuplicateNodes(xx)

	fmt.Printf("%v",xx)
}
