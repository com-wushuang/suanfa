package link

import (
	"testing"
)

func TestReverse(t *testing.T) {
	li := InitList()
	li = Reverse(li)
	PrintList(li)
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

func NestedReverse(head *ListNode) {

}
