# 合并\分解链表
## 合并两个有序链表
- 题目:给你输入两个有序链表，请你把他俩合并成一个新的有序链表。
- 循环每次比较 `p1` 和 `p2` 的大小，把较小的节点接到结果链表上。
- 这个算法的逻辑类似于拉拉链，`l1`, `l2` 类似于拉链两侧的锯齿，指针 `p` 就好像拉链的拉索，将两个有序链表合并。
- 代码中还用到一个链表的算法题中是很常见的「虚拟头结点」技巧，也就是 `dummy` 节点。如果不使用 `dummy` 虚拟节点，代码会复杂很多，而有了 `dummy` 节点这个占位符，可以避免处理空指针的情况，降低代码的复杂性。
```go
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
```
## 单链表的分解
![divide_link](https://github.com/com-wushuang/suanfa/blob/main/image/divide_link.jpeg)
- 在合并两个有序链表时让你合二为一，而这里需要分解让你把原链表一分为二。
- 具体来说，我们可以把原链表分成两个小链表，一个链表中的元素大小都小于 `x`，另一个链表中的元素都大于等于 `x`，最后再把这两条链表接到一起，就得到了题目想要的结果。
```go
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
```
# 快慢指针
## 判断链表是否有环
- 如果链表存在环，那么扫描链表不会停止下来
- 利用快慢指针，那么一定会有快慢指针相遇的时候
- 就像是在环形的赛道中赛跑，跑的快的人一定会超跑的慢的人一圈
- 他们相遇的时候，快的人超了慢的人整整一圈
```go
func isLoop(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { // 注意这个循环终止条件,因为fast一次走两步
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
```
## 寻找环的起始节点
- 相遇时，fast走了`2k`步，slow走了`k`步
- 设环的起始点`loopStart`和相遇点`meetPoint`的距离为`m`
- 那么链表头节点`head`到`loopStart`的距离为`k-m`
- 因为fast超了slow一圈，所以圈的大小为k
- 推导出，相遇点`meetPoint`到环的起始点`loopStart`的距离也为`k-m`
- 因此利用这个，我们先让快慢指针第一次相遇
- 相遇后，让慢指针从头开始，两个指针一次走一步，再一次相遇的时候，就是环的起点
```go
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
```
## 链表的中点
- 快指针一次前进两步，慢指针一次前进一步，当快指针到达链表尽头时，慢指针就处于链表的中间位置
- 当链表的长度是奇数时，slow 恰巧停在中点位置；如果长度是偶数，slow 最终的位置是中间偏右
```go
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
```
## 删除链表倒数第n个节点
- 使用快慢指针，让快指针先走 n 步，然后快慢指针开始同速前进。
- 这样当快指针走到链表末尾 null 时，慢指针所在的位置就是倒数第 n 个链表节点（n 不会超过链表长度）。
```go
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
```

# 反转链表
## 逆序打印链表,链表的后续遍历
- 其实链表可以看成是一种特殊的二叉树
```go
func reversePrint(head *ListNode) {
	if head == nil {
		return
	}
	reversePrint(head.Next)
	fmt.Print(head.Val)
}
```
## 反转整个链表
### 非递归
```go
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
```

### 递归
```go
func NestedReverse(head *ListNode) *ListNode {
	if head.Next != nil {
		return head
	}
	newHead := NestedReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
```

## 反转链表前 N 个节点
### 非递归
```go
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
```

### 递归
```go
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
```

## 反转链表的一部分
### 非递归
```go
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
```
### 递归
```go
func NestedReverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return NestedReverseK(head, m)
	}

	newHead := NestedReverseBetween(head.Next, n-1, m-1)
	head.Next = newHead
	return head
}
```

##  K 个一组反转链表
```go
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	for ; cur != b; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	a.Next = b
	return pre
}

func ReverseGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	head.Next = ReverseGroup(b, k)
	return newHead
}
```
# 回文链表
```go
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
```