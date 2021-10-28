package array

import (
	"fmt"
	"testing"
)

func TestSlideWindow(t *testing.T) {
	fmt.Println(minimumSubstring("ADOBECODEBANC", "ABC"))
	fmt.Println(checkInclusion("eidbaooo", "ab"))
	fmt.Println(checkInclusion("eidboaoo", "ab"))
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// 最小覆盖子串
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

// 字符串排列
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

// 找所有字母异位词
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

//最长子串的长度
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
