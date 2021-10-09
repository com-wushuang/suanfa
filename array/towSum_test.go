package array

import "testing"

func TestTwoSum(t *testing.T) {

}

// 暴力解法 时间复杂度n方
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

// 条件target=a+b，在已知target和a的条件下，在查找另一个数b，在数组的数据结构中，只有扫描数组才能达到该目的，时间复杂度比较高
// 通过map的数据结构能够把查找的效率改善到o(1)
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

// 双指针法，对于有序数组而言
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
