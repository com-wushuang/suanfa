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
	for ; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil { // 链表长度是奇数
		fast = fast.Next
	}

	right := Reverse(slow)
	left := head

	for ; right != nil; {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
