## 滑动窗口
滑动窗口能解决大部分子串的问题，而且代码都是有固定的框架，掌握了框架代码和算法的核心思想，能解决大部分子串的问题
### 最小覆盖子串
- 题意：在 S(source) 中找到包含 T(target) 中全部字母的一个子串，且这个子串一定是所有可能子串中最短的
- 在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引左闭右开区间 [left, right) 称为一个「窗口」
- 先不断地增加 right 指针扩大窗口 [left, right)，直到窗口中的字符串符合要求（包含了 T 中的所有字符）
- 此时，停止增加 right，转而不断增加 left 指针缩小窗口 [left, right)，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。同时，每次增加 left，我们都要更新一轮结果
- 重复上面的两步，直到 right 到达字符串 S 的尽头
这个思路其实也不难，第 2 步相当于在寻找一个「可行解」，然后第 3 步在优化这个「可行解」，最终找到最优解，也就是最短的覆盖子串。

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

### 字符串排列
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false
- 题意：给你一个 S 和一个 T，请问你 S 中是否存在一个子串，包含 T 中所有字符且不包含其他字符？
- 算法的框架和第一题一样，不同点在于，当窗口中包含了所有的T中字符时，如果窗口的长度恰好等于T的长度，那么返回true
- 否则，继续移动left指针

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

### 找所有字母异位词
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
### 最长无重复子串
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
