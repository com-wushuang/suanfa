package link

import (
	"testing"
)

func TestReverse(t *testing.T) {
	li := InitList()
	li = Reverse(li)
	PrintList(li)

	li2 := InitList()
	li2 = NestedReverse(li2)
	print(li2)
}

func Reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func NestedReverse(head *ListNode) *ListNode {
	if head.Next != nil {
		return head
	}
	newHead := NestedReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
