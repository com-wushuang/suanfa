## 快慢指针
### 判断链表是否有环
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
### 寻找环的起始节点
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
### 链表的中点
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
### 删除链表倒数第n个节点
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
## 左右指针
左右指针在数组中实际是指两个索引值，一般初始化为 left = 0, right = len(nums) - 1 
### 二分查找
略
### 两数之和
略
### 反转数组
- 将一个数组反转，初始化左右数组指针，然后左右元素互相交换
```go
func reverseArray(a []int) []int {
	left := 0
	right := len(a) - 1
	for left < right {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
```
### 滑动窗口算法
略