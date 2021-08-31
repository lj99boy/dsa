package detectCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nHead := head.Next
	sHead := head
	for ; nHead.Next != nil; nHead = nHead.Next.Next {
		if nHead == sHead {
			return sHead
		} else {
			sHead = sHead.Next
		}
	}

	return nil
}
