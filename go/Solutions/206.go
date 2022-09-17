package Solutions

func reverseList(head *ListNode) *ListNode {
	t, h := reverse(head)
    if t!=nil{
        t.Next=nil
    }
    
	return h
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next == nil {
		return head, head
	}
	last, h := reverse(head.Next)
	last.Next = head
	return head, h
}
