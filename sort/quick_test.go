package sort

import (
	"fmt"
	"testing"
)

var quickArry = []int{8, 1, 5, 7, 2, 3, 9, 6, 4, 10, 19, 18, 12, 11, 15, 13, 14, 17, 16}
var quickArry2 = []int{1, 1, 3, 1, 1, 2, 3, 2, 2}

func TestQucik(t *testing.T) {
	QuickSort(quickArry2, 0, len(quickArry2)-1)
	fmt.Printf("%v", quickArry2)
}

func QuickSort(a []int, low, high int) {
	if low < high {
		mid := partition(a, low, high)
		QuickSort(a, low, mid-1)
		QuickSort(a, mid+1, high)
	}
}

func partition(a []int, low, high int) int {
	pivot := high
	i := low  // 左指针
	j := high // 右指针

	for i != j {
		for i < j && a[i] <= a[pivot] { // 从左边开始扫描
			i++
		}
		for i < j && a[j] >= a[pivot] {
			j--
		}
		a[i], a[j] = a[j], a[i] // 交换元素
	}
	a[i], a[pivot] = a[pivot], a[i] // pivot放置在最终的位置上
	return i
}
