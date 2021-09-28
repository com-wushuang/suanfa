package link

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitList() *ListNode {
	head := &ListNode{
		Val: 1,
	}
	i := 2
	p := head
	for ; i < 10; i++ {
		newNode := &ListNode{
			Val: i,
		}
		p.Next = newNode
		p = p.Next
	}
	return head
}

func InitListWithArray(a []int) *ListNode {
	head := &ListNode{
		Val: a[0],
	}
	p := head
	for i := 1; i < len(a); i++ {
		newNode := &ListNode{
			Val: a[i],
		}
		p.Next = newNode
		p = p.Next
	}
	return head
}

func PrintList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}