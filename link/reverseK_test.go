package link

import (
	"fmt"
	"testing"
)

func TestReverseK(t *testing.T) {
	li := InitList()
	li = ReverseK(li, 3)
	PrintList(li)

	fmt.Println()

	li2 := InitList()
	li2 = NestedReverseK(li2, 3)
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

func NestedReverseK(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	newHead := NestedReverseK(head, k-1)
	sucK := head.Next.Next // 第 k+1 个节点的指针
	head.Next.Next = newHead
	indexK := head.Next // 第 k 个节点的指针
	head.Next = sucK    // head 链接 k+1 节点

	return indexK
}
