package reverseKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	t := head
	for i:=0;i<k;i++{
		if t==nil {
			return head
		}
		t=t.Next
	}


	part:=reverse(head,t)
	head.Next = reverseKGroup(t,k)

	return part
}

func reverse(left *ListNode, right *ListNode) *ListNode {
	var pre, cur, next *ListNode
	cur = left
	next = left
	pre = nil
	for cur != right {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
