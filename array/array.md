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

# 双指针
## 两数之和
- 给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出和为目标值 `target` 的那两个整数，并返回它们的数组下标。
- 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
### 暴力解法
- 暴力解法思路比较简单，就是穷举数组
```go
func twoSum1(a []int, target int) (result [2]int) {
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ { // j 下标从 i+1 开始
			if a[i]+a[j] == target {
				result[0], result[1] = i, j
				return
			}
		}
	}
	
	result[0], result[1] = -1, -1
	return
}
```

### 利用 map 提高查询效率
- 算法的复杂度是n方
- 条件 `target=a+b`，在已知 `target` 和 `a` 的条件下，在查找另一个数 `b` ，在数组的数据结构中，只有扫描数组才能达到该目的，时间复杂度比较高
- 通过 `map` 的数据结构能够把查找 `b` 的效率改善到 `o(1)`
```go
func twoSum2(a []int, target int) (result [2]int) {
	look := make(map[int]int)
	// 先把数组放入一个map中方便查找
	for i := 0; i < len(a); i++ {
		look[a[i]] = i
	}

	for i := 0; i < len(a); i++ {
		need := target - a[i]
		j, ok := look[need]
		if ok && i != j { // 保证不是他自己
			result[0], result[1] = i, j
			return
		}
	}

	result[0], result[1] = -1, -1
	return
}
```
- 由于哈希表的查询时间为 `O(1)`，算法的时间复杂度降低到 `O(N)`，但是需要 `O(N)` 的空间复杂度来存储哈希表。不过综合来看，是要比暴力解法高效的。

## 有序数组的两数之和
如果上述问题中的中给的数组是有序的，应该如何编写算法呢？
- 利用双指针法，初始化时，`left`指针在数组的头部指向`a`，`right`在数组的尾部指向`b`
- 那么 `a+b`一定是在中间位置(因为数组有序，默认为升序排列)
- 如果`target > a+b`,那么需要`left++`，让`a+b`变大
- 反之`target < a+b`,那么需要`right--`,让`a+b`变小
- 如果`target == a+b`,那么得到结果返回

```go
func twoSum3(a []int, target int) (result [2]int) {
	i := 0
	j := len(a) - 1
	for i != j {
		sum := a[i] + a[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			result[0], result[1] = i, j
			return
		}
	}

	result[0], result[1] = -1, -1
	return
}
```

## 有序数组去重
- 快慢指针法，一个 `fast` 指针，一个 `slow` 指针。
- 快指针 `fast` 在前面探路，找到一个不重复的元素就告诉 `slow` 指针，并让 `slow` 前进一步。
- 当 `fast` 指针所描完数组中的所有元素，那么`[0~slow]`就是不重复的元素。

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

## 移除数组中的指定元素
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
## 移除零
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

## 反转数组
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
## 回文串判断
- 回文串就是正着读和反着读都一样的字符串。
- 比如说字符串 `aba` 和 `abba` 都是回文串，因为它们对称，反过来还是和本身一样。
- 反之，字符串 `abac` 就不是回文串。
```go
func isPalindrome(a string) bool {
	charArray := []rune(a)
	left, right := 0, len(a)-1
	for left < right {
		if charArray[left] != charArray[right] { // 如果存在左右不想等情况，那么就不是回文串
			return false
		}
		left++
		right--
	}
	return true
}
```

## 最长回文子串
- 找回文串的难点在于，回文串的的长度可能是奇数也可能是偶数，解决该问题的核心是从中心向两端扩散的双指针技巧。
- 
```go
func longestPalindrome(a string) string {
	charArray := []rune(a)
	var res []rune
	for i := 0; i < len(charArray); i++ {
		s1 := palindrome(charArray, i, i)  // 以 charArray[i] 为中心的最长回文子串
		s2 := palindrome(charArray, i, i+1)  // 以 charArray[i] 和 charArray[i+1] 为中心的最长回文子串
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return string(res)
}

// 在 charArray 中寻找以 charArray[l] 和 charArray[r] 为中心的最长回文串
func palindrome(charArray []rune, l, r int) []rune {
	// 防止索引越界
	for l >= 0 && r < len(charArray) && charArray[l] == charArray[r] {
		// 双指针，向两边展开
		l--
		r++
	}
	return charArray[l+1 : r] // 返回以 charArray[l] 和 charArray[r] 为中心的最长回文串
}
```

# 滑动窗口
滑动窗口能解决大部分子串的问题，而且代码都是有固定的框架，掌握了框架代码和算法的核心思想，能解决大部分子串的问题
## 最小覆盖子串
- 题意：在 `S(source)` 中找到包含 `T(target)` 中全部字母的一个子串，且这个子串一定是所有可能子串中最短的
- 在字符串 `S` 中使用双指针中的左右指针技巧，初始化 `left = right = 0`，把索引左闭右开区间 `[left, right)` 称为一个「窗口」
- 先不断地增加 `right` 指针扩大窗口 `[left, right)`，直到窗口中的字符串符合要求（包含了 `T` 中的所有字符）
- 此时，停止增加 `right`，转而不断增加 `left` 指针缩小窗口 `[left, right)`，直到窗口中的字符串不再符合要求（不包含 `T` 中的所有字符了）。同时，每次增加 `left`，我们都要更新一轮结果
- 重复上面的两步，直到 `right` 到达字符串 `S` 的尽头
- 这个思路其实也不难，第 `2` 步相当于在寻找一个「可行解」，然后第 `3` 步在优化这个「可行解」，最终找到最优解，也就是最短的覆盖子串。

```go
func minimumSubstring(a, t string) string {
	// window哈希表，用来记录窗口中包含的T的字符
	window := make(map[byte]int)

	// need哈希表，用来表示T中字符及其出现的次数
	need := make(map[byte]int)
	for _, item := range []byte(t) {
		need[item]++
	}

	// 滑动窗口的左右指针
	left, right := 0, 0

	// valid表示window中满足T的字符串的个数
	valid := 0

	// 原始字符串S转换成字节数组
	chars := []byte(a)

	// 最小覆盖子串的长度和起始位置
	start, length := 0, 0

	for right < len(chars) {
		// 扩大窗口
		char := chars[right]
		right++

		// 如果char是目标字符串T中的字符，更新window哈希表中的数据
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] { // 如果有一个字符满足了条件,valid加1
				valid++
			}
		}

		// 当window中包含了所有的T中的字符，开始移动left指针
		for valid == len(need) {
			// 先更新答案
			length = right - left
			start = left

			// 收缩窗口
			char := chars[left]
			left++

			// 更新窗口中的数据和valid
			if _, ok := need[char]; ok {
				if window[char] == need[char] {
					valid--
				}
				window[char]--
			}
		}
	}
	return string(chars[start : start+length])
}
```

## 字符串排列
给你两个字符串 `s1` 和 `s2` ，写一个函数来判断 `s2` 是否包含 `s1` 的排列。如果是，返回 `true` ；否则，返回 `false`
- 题意：给你一个 `S` 和一个 `T`，请问你 `S` 中是否存在一个子串，包含 `T` 中所有字符且不包含其他字符？
- 算法的框架和第一题一样，不同点在于，当窗口中包含了所有的 `T` 中字符时，如果窗口的长度恰好等于T的长度，那么返回 `true`
- 否则，继续移动 `left` 指针

```go
func checkInclusion(a, t string) bool {
	window := make(map[byte]int)

	// 初始化need集合
	need := make(map[byte]int)
	for _, item := range []byte(t) {
		need[item]++
	}

	left, right := 0, 0
	valid := 0
	chars := []byte(a)

	for right < len(chars) {
		// 扩大窗口
		char := chars[right]
		right++

		// 更新窗口中的数据
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] { // 如果有一个字符满足了条件,valid加1
				valid++
			}
		}

		// 窗口扩大到包含了所有子串的元素
		for valid == len(need) {
			// 如果窗口的大小等于子串的长度,返回true
			if right-left == len(t) {
				return true
			}

			// 否则，就开始缩小窗口
			char := chars[left]
			left++
			if _, ok := need[char]; ok {
				if window[char] == need[char] {
					valid--
				}
				window[char]--
			}
		}
	}
	return false
}
```

## 找所有字母异位词
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
注：异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
- 题意：相当于，输入一个串 S，一个串 T，找到 S 中所有 T 的排列，返回它们的起始索引。
- 这个题和上一题本质上没有区别，只是返回结果的不同

```go
func findAnagrams(a, t string) []int {
	res := make([]int, 0)

	window := make(map[byte]int)

	// 初始化need集合
	need := make(map[byte]int)
	for _, item := range []byte(t) {
		need[item]++
	}

	left, right := 0, 0
	valid := 0
	chars := []byte(a)

	for right < len(chars) {
		// 扩大窗口
		char := chars[right]
		right++
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left == len(t) {
				res = append(res, left)
			}

			char := chars[left]
			left++
			if _, ok := need[char]; ok {
				if window[char] == need[char] {
					valid--
				}
				window[char]--
			}
		}
	}
	return res
}
```
## 最长无重复子串
给定一个字符串 s ，请你找出其中不含有重复字符的`最长子串`的长度。
- 先移动`right`指针，每次进入窗口的字符都存放在window哈希表中
- 当进入的字符重复了的时候，开始移动`left`指针，缩小窗口，使得窗口中不包含重复的字符

```go
func lengthOfLongestSubstring(a string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	chars := []byte(a)

	for right < len(chars) {
		c := chars[right]
		right++
		window[c]++ // 每个进入滑动窗的字符，都先存储在window哈希表中

		for window[c] > 1 { // 当进入的字符重复了，开始移动左边指针，缩小窗口，使得窗口中不包含重复的字符
			d := chars[left]
			left++
			window[d]--
		}
		res = max(right-left, res) //更新结果
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

# 前缀和数组
前缀和数组是通过原始数组推倒出的新的数组
## 定义
`a`为原数组，`sum`为前缀和数组，`sum[i]= a[0]+a[1]+ ...+ a[i-1]` , 特别的`sum[0]=0`
```go
    a:=[]int{1,2,3,4} // 原始数组
    
    // 前缀和数组
    sum := make([]int, len(a)+1)  
    sum[0] = 0
    for i := 0; i < len(a); i++ {
        sum[i+1] = sum[i] + a[i]
    }
```
- 如果原始数组`a`的长度为`n`，那么前缀和数组的长度为`n+1`.
- 关键代码：`sum[i+1] = sum[i] + a[i]` ，如果将 `sum[i]= a[0]+...+a[i-1]` 带入，那么就是`sum[i+1]=a[0]+...+a[i]`符合前面的定义

## 例子
给定一个数组，求解和为 k 的连续子数组的个数
```go
func preSum(a []int, k int) (res int) {
	// 构造前缀和数组
	sum := make([]int, len(a)+1)
	sum[0] = 0
	for i := 0; i < len(a); i++ {
		sum[i+1] = sum[i] + a[i]
	}

	// 求解和为 k 的连续子数组的个数
	for i := 1; i < len(sum); i++ {
		for j := 0; j < i; j++ {
			if sum[i]-sum[j] == k {
				res++
			}
		}
	}
	return
}
```
- 原理：`a[i..j]`的和，只需要 `sum[j+1]-sum[i]`（通过定义的公式带入即可证明，和上面的关键代码证明方法类似）
- 注意遍历子数组的代码结构

# 差分数组
差分数组也是由原始数组推导出的新数组
## 定义
`a`为原数组，`diff`为差分数组,`diff[i] = a[i] - a[i-1]`
```go
	a := []int{1, 2, 3, 4} // 原始数组

	// 差分数组
	diff := make([]int, len(a))
	diff[0] = a[0]
	for i := 1; i < len(a); i++ {
		sum[i] = a[i] - a[i-1]
	}

	//差分数组求前缀和能够得到原始数组
	res := make([]int, 5)
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}
```
- 原始数组和差分数组的长度都是`n`,这一点和前缀和数组是有区别的
- 差分数组的用处：差分数组可以让我们快速进行区间增减的操作，如果你想对区间 nums[i..j] 的元素全部加 3，那么只需要让 `diff[i] += 3`，然后再让 `diff[j+1] -= 3` 即可
- 差分数组求前缀和数组就是原数组（可根据定义带入证明）
## 例子
```
这里有 n 个航班，它们分别从 1 到 n 进行编号
有一份航班预订表`bookings` ，表中第`i`条预订记录`bookings[i] = [firsti, lasti, seatsi]`意味着在从 firsti到 lasti（包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
```
示例 1：
```
输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]
```
### 分析
- 初始化一个数组`a`，长度为5，初始化的时候数组中的元素全部都是`0`，代表每个航班的人数
- 我们要做的事情是，`a[0~1]加10`，`a[1~2]加20`，`a[1~4]加25`，这样就可以得到答案
- 因为是在区间做加法，那么用到差分数组的性质
- `a`的差分数组`diff`也是长度为`5`，元素全部是`0`的数组
- 对原数组`a[0~1]加10`根据差分的性质,需要对差分数组`diff[0]+=10`和`diff[2]-=10`
- 对所有的区间操作用上一步的方法操作
- 最后对得出的差分数组求前缀和即可得到本题答案

### 代码
```go
func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, 5) // 原始数组的差分数组
	for _, booking := range bookings {

		// 区间和加数
		left := booking[0] - 1 //题目中的航班是从0编号的
		right := booking[1]
		k := booking[2]

		diff[left] += k
		if right < n {
			diff[right] -= k
		}
	}

	// 差分数组求前缀和得到原始数组
	res := make([]int, 5)
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}
```


