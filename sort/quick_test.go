package sort

import (
	"fmt"
	"testing"
)

var quickArry = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestQucik(t *testing.T) {
	QuickSort(quickArry, 0, len(quickArry)-1)
	fmt.Printf("%v", quickArry)
}

func QuickSort(a []int, low, high int) {
	if low < high {
		mid := partition(a, low, high)
		QuickSort(a, low, mid-1)
		QuickSort(a, mid+1, high)
	}
}

func partition(a []int, low, high int) int {
	pivot := a[low]

	for low != high {
		for low < high && pivot <= a[high] {
			high--
		}
		a[low], a[high] = a[high], a[low]

		for low < high && pivot >= a[low] {
			low++
		}
		a[low], a[high] = a[high], a[low]
	}

	a[low] = pivot
	return low
}
