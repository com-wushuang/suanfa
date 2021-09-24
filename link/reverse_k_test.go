package link

import "testing"

func TestReverseK(t *testing.T) {
	li := InitList()
	li = ReverseK(li, 3)
	PrintList(li)
}

func ReverseK(head *ListNode, k int) *ListNode {
	var pre *ListNode
	cur := head
	for i := 0; i < k; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	head.Next = cur
	return pre
}
