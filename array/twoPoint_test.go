package array

import (
	"fmt"
	"testing"
)

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

func InitLoopList() *ListNode {
	var loopBegin *ListNode
	var loopEnd *ListNode
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
		if i == 5 {
			loopBegin = newNode
		}
		if i == 9 {
			loopEnd = newNode
		}
	}
	loopEnd.Next = loopBegin

	return head
}

func PrintList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}

func TestTwoPoint(t *testing.T) {
	li1 := InitLoopList()
	li2 := InitList()
	fmt.Println(isLoop(li1))
	fmt.Println(isLoop(li2))

	fmt.Println(loopIndex(li1).Val)

	li3 := InitList()
	PrintList(li3)
	fmt.Println(midOfLink(li3).Val)

	li4 := InitList()
	PrintList(removeNthFromEnd(li4, 5))
}

// 判断链表是否有环
func isLoop(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { // 注意这个循环终止条件
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 返回链表环的起始点
func loopIndex(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	// 上面的代码类似 isLoop 函数
	if fast == nil || fast.Next == nil {
		// fast 遇到空指针说明没有环
		return nil
	}

	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// 寻找链表的中点
func midOfLink(head *ListNode) *ListNode {
	slow := head
	fast := head
	if fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 如果是奇数个节点，那么slow是链表的中点，如果是偶数个节点，那slow中间靠右的节点
	return slow
}

// 删除链表的倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head

	// 先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 如果fast走到了nil,说明倒数第n个节点就是第一个节点
	if fast == nil {
		head = head.Next
		return head
	}

	// 删除节点
	for fast.Next != nil { // 让fast指向最后一个节点，而不是nil这个跟后面的操作有关
		fast = fast.Next
		slow = slow.Next
	}

	// slow.Next 就是倒数第 n 个节点，删除它
	slow.Next = slow.Next.Next
	return head
}

// 反转数组
func reverseArray(a []int) []int {
	left := 0
	right := len(a) - 1
	for left < right {
		a[left], a[right] = a[right], a[left]
	}
	return a
}

func TestIsPalindrome(t *testing.T) {
	a := "abc"
	fmt.Println(isPalindrome(a))

	b := "aba"
	c := "abba"
	fmt.Println(isPalindrome(b))
	fmt.Println(isPalindrome(c))
}

// 回文串
func isPalindrome(a string) bool {
	charArray := []rune(a)
	left, right := 0, len(a)-1
	for left < right {
		if charArray[left] != charArray[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {
	a := "babad"
	fmt.Println(longestPalindrome(a))
	b := "cbbd"
	fmt.Println(longestPalindrome(b))
}

// 最长回文串
func longestPalindrome(a string) string {
	charArray := []rune(a)
	var res []rune
	for i := 0; i < len(charArray); i++ {
		s1 := palindrome(charArray, i, i)
		s2 := palindrome(charArray, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return string(res)
}

func palindrome(charArray []rune, l, r int) []rune {
	// 防止索引越界
	for l >= 0 && r < len(charArray) && charArray[l] == charArray[r] {
		// 双指针，向两边展开
		l--
		r++
	}
	return charArray[l+1 : r]
}
