package removeDuplicateNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	for tHead := head; tHead.Next != nil; tHead = tHead.Next {
		for tTHead := tHead; tTHead.Next != nil; {
			if tTHead.Next.Val == tHead.Val {
				tTHead.Next = tTHead.Next.Next
			} else {
				tTHead = tTHead.Next
			}
		}
		if tHead.Next == nil {
			break
		}
	}

	return head
}
