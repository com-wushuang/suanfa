package link

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	li := InitListWithArray([]int{1, 2, 3, 2, 1})
	fmt.Println(isPalindromeNested(li))
	fmt.Println(isPalindrome(li))
}

func TestMergeOrderLink(t *testing.T) {
	l1 := InitListWithArray([]int{1, 2, 2, 5, 8})
	l2 := InitListWithArray([]int{3, 7, 9, 9, 12, 12, 14, 15})
	l3 := mergeOrderLink(l1, l2)
	PrintList(l3)
}

// 合并有序链表
func mergeOrderLink(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		// 比较 p1 和 p2 两个指针
		// 将值较小的的节点接到 p 指针
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		// p 指针不断前进
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

// 拆分链表
func divideLink(l *ListNode, x int) *ListNode {
	// p1, p2 指针负责生成结果链表
	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -1}
	// p 负责遍历原链表，类似合并两个有序链表的逻辑，这里是将一个链表分解成两个链表
	p1, p2 := dummy1, dummy2
	p := l
	for p != nil {
		if p.Val <= x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 断开原链表中的每个节点的 next 指针
		tmp := p.Next
		p.Next = nil
		p = tmp
	}
	// 连接两个链表
	p1.Next = dummy2.Next
	return dummy1.Next
}

// 逆序打印链表,链表的后续遍历
func reversePrint(head *ListNode) {
	if head == nil {
		return
	}
	reversePrint(head.Next)
	fmt.Print(head.Val)
}

// 递归判断回文列表
var left *ListNode

func isPalindromeNested(head *ListNode) bool {
	left = head
	return traverse(head.Next)
}
func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)

	// 后序遍历代码
	res = res && right.Val == left.Val
	left = left.Next
	return res
}

// 快慢指针判断回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil { // 链表长度是奇数
		fast = fast.Next
	}

	right := Reverse(slow)
	left := head

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
