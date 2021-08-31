package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	rever := make([]*ListNode,0)

	for tHead:=head; tHead !=nil;tHead = tHead.Next {
		rever = append(rever,tHead)
	}


	for i:=len(rever)-1;i>=0;i--{
		if head.Val != rever[i].Val{
			return false
		}else {
			head = head.Next
		}
	}

	return true
}