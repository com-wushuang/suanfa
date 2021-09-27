package link

import "testing"

func TestReverseGroup(t *testing.T) {
	li:=InitList()
	li=ReverseGroup(li,3)
	PrintList(li)
}

func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	for ; cur != b; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	a.Next = b
	return pre
}

func ReverseGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	head.Next = ReverseGroup(b, k)
	return newHead
}
