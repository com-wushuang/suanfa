# 二分查找和细节
```go
func binarySearch(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	left, right := 0, len(a)-1
	for left <= right {
		mid := (left + right) / 2
		if a[mid] == target {
			return mid
		}
		if target < a[mid] {
			right = mid - 1
		}
		if target > a[mid] {
			left = mid + 1
		}
	}
	return -1
}
```
## 细节
### 为什么 while 循环的条件中是 <=，而不是 <？
- 因为初始化 `right` 的赋值是 `nums.length - 1`，即最后一个元素的索引，而不是 `nums.length`。
- 这二者可能出现在不同功能的二分查找中，区别是：前者相当于两端都闭区间 `[left, right]`，后者相当于左闭右开区间 `[left, right)`，因为索引大小为 `nums.length` 是越界的。
- 我们这个算法中使用的是前者 `[left, right]` 两端都闭的区间。这个区间其实就是每次进行搜索的区间。
- 什么时候应该停止搜索呢？当然，找到了目标值的时候可以终止:
```go
		if a[mid] == target {
			return mid
		}
```
但如果没找到，就需要 `while` 循环终止，然后返回 `-1`。那 `while` 循环什么时候应该终止？搜索区间为空的时候应该终止，意味着你没得找了，就等于没找到嘛。
- `while(left <= right)` 的终止条件是 `left == right + 1`，写成区间的形式就是 `[right + 1, right]`，或者带个具体的数字进去 `[3, 2]`，可见这时候区间为空，因为没有数字既大于等于 `3` 又小于等于 `2` 的吧。所以这时候 `while` 循环终止是正确的，直接返回 `-1` 即可。
- `while(left < right)` 的终止条件是 `left == right`，写成区间的形式就是 `[right, right]`，或者带个具体的数字进去 `[2, 2]`，这时候区间非空，还有一个数 `2`，但此时 `while` 循环终止了。也就是说这区间 `[2, 2]` 被漏掉了，索引 `2` 没有被搜索，如果这时候直接返回 `-1` 就是错误的。

### 为什么 left = mid + 1，right = mid - 1？我看有的代码是 right = mid 或者 left = mid，没有这些加加减减，到底怎么回事，怎么判断？
- 刚才明确了「搜索区间」这个概念，而且本算法的搜索区间是两端都闭的，即 `[left, right]`。那么当我们发现索引 `mid` 不是要找的 `target` 时，下一步应该去搜索哪里呢？
- 当然是去搜索区间 `[left, mid-1]` 或者区间 `[mid+1, right]` 对不对？因为 `mid` 已经搜索过，应该从搜索区间中去除。

### 此算法有什么缺陷？
- 比如说给你有序数组 `nums = [1,2,2,2,3]`，`target` 为 `2`，此算法返回的索引是 `2`，没错。但是如果我想得到 `target` 的左侧边界，即索引 `1`，或者我想得到 `target` 的右侧边界，即索引 `3`，这样的话此算法是无法处理的。
- 这样的需求很常见，你也许会说，找到一个 target，然后向左或向右线性搜索不行吗？可以，但是不好，因为这样难以保证二分查找对数级的复杂度了。

## 寻找左侧边界的二分查找
核心思想: 找到 `target` 时不要立即返回，而是缩小`「搜索区间」`的上界 `right`，在区间 `[left, mid]` 中继续搜索，即不断向左收缩，达到锁定左侧边界的目的。
```go
func leftBound(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}

	left, right := 0, len(a)-1
	for left <= right {
		mid := (left + right) / 2
		if target == a[mid] {
			right = mid - 1
		}
		if target < a[mid] {
			right = mid - 1
		}
		if target > a[mid] {
			left = mid + 1
		}
	}

	if left > len(a)-1 { // target 比数组中所有的元素都大
		return -1
	}
	if a[left] == target { // 如果前面不加判断直接引用，那么会存在数组越界的问题
		return left
	}
	return -1
}
```
- 因为最后循环终止的条件是 `left == right + 1` ，如果 `right` 一直不变，最后循环终止的时候 `left` 会有数组越界的问题。

## 寻找右侧边界的二分查找
核心思想: 找到 `target` 时，不要立即返回，而是增大「搜索区间」的左边界 `left`，使得区间不断向右靠拢，达到锁定右侧边界的目的。
```go
func rightBound(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	left, right := 0, len(a)-1
	for left <= right {
		mid := (left + right) / 2
		if target == a[mid] {
			left = mid + 1
		}
		if target < a[mid] {
			right = mid - 1
		}
		if target > a[mid] {
			left = mid + 1
		}
	}
	if left == 0 { // 如果 target 比数组中所有的元素都小
		return -1
	}
	if a[left-1] == target {
		return left - 1
	}
	return -1
}
```