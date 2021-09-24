package link

import "testing"

func TestReverseBetween(t *testing.T) {
	li := InitList()
	li = ReverseBetween(li, 3, 5)
	PrintList(li)
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
