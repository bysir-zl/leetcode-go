package test

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) toArray() (a []int) {
	for head := l; head != nil; head = head.Next {
		a = append(a, head.Val)
	}

	return
}

func (l *ListNode) fromArray(a []int) {
	curr := l
	for _, v := range a {
		curr.Val = v
		next := &ListNode{}
		curr.Next = next
		curr = next
	}

	return
}
