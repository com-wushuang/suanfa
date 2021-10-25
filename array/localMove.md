## 原地修改数组
### 有序数组去重
- 快慢指针法，一个fast指针，一个slow指针。
- 快指针fast在前面探路，找到一个不重复的元素就告诉slow指针，并让slow前进一步。
- 当fast指针所描完数组中的所有元素，那么`[0~slow]`就是不重复的元素。

```go
// 移除有序数组中重复的元素,返回新数组的长度
func moveDuplicate(a []int) int {
	slow, fast := 0, 0
	for fast < len(a) {
		if a[slow] != a[fast] {
			slow++
			a[slow] = a[fast]
		}
		fast++ //一直向前 ，直到遇到不重复的元素
	}
	return slow + 1
}
```

### 移除数组中的指定元素
- 和上述思路完全一致，稍微在代码实现上有些许不同。
```go
// 将指定的元素移除，返回新数组的长度，不必考虑新数组超出部分的元素
func moveElement(a []int, val int) int {
	slow, fast := 0, 0
	for fast < len(a) {
		if a[fast] != val { // 
			a[slow] = a[fast]
			slow++
		}
		fast++
	}

	return slow
}
```
### 移除零

```go
// 将数组中所有的0移动到数组的末尾,上述函数的特殊形式，并需要保留移除的0
func moveZero(a []int) {
	slow,fast := 0,0
	for fast < len(a) {
		if a[fast] != 0 {
			a[slow] = a[fast]
			slow++
		}
		fast++

	}
	// slow后面的所有元素赋值为0
	for slow < len(a) {
		a[slow] = 0
		slow++
	}
}
```

 