## 两数之和
给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
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
- 算法的复杂度是n方
- 条件target=a+b，在已知target和a的条件下，在查找另一个数b，在数组的数据结构中，只有扫描数组才能达到该目的，时间复杂度比较高
- 通过map的数据结构能够把查找b的效率改善到o(1)

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
由于哈希表的查询时间为 O(1)，算法的时间复杂度降低到 O(N)，但是需要 O(N) 的空间复杂度来存储哈希表。不过综合来看，是要比暴力解法高效的。
### 有序数组的两数之和问题
如果问题1中的中给的数组是有序的，应该如何编写算法呢？
- 利用双指针法，初始化时，`left`指针在数组的头部指向`a`，`right`在数组的尾部指向`b`
- 那么 `a+b`一定是在两数之和的中间位置(因为数组有序，默认为升序排列)
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


