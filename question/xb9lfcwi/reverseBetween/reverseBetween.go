package reverseBetween

type ListNode struct {
	     Val int
	     Next *ListNode
}

var lastNext *ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}

	if left == 1 {
		return innerReverse(head,right)
	}

	head.Next = reverseBetween(head.Next,left-1,right-1)
	return head
}

func innerReverse(head *ListNode,n int)  *ListNode  {
	if n == 1 {
		lastNext = head.Next
		return head
	}

	last := innerReverse(head.Next,n-1)

	head.Next.Next = head
	head.Next = lastNext

	return last
}