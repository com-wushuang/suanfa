package link

import "testing"

func TestReverseBetween(t *testing.T) {
	li := InitList()
	li = ReverseBetween(li, 3, 5)
	PrintList(li)

	li2 := InitList()
	li2 = NestedReverseBetween(li2, 3, 5)
	PrintList(li2)
}

func ReverseBetween(head *ListNode, m, n int) *ListNode {
	var pre *ListNode
	cur := head
	for i := 0; i < m-1; i++ {
		pre = cur
		cur = cur.Next
	}
	startPre := pre
	startCur := cur

	for i := 0; i < n-m+1; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	startPre.Next = pre
	startCur.Next = cur
	return head
}

func NestedReverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return NestedReverseK(head, m)
	}

	newHead := NestedReverseBetween(head.Next, n-1, m-1)
	head.Next = newHead
	return head
}
