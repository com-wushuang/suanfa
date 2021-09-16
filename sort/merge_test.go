package sort

import (
	"fmt"
	"testing"
)

var mergeArry = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestMerge(t *testing.T) {

	tmp := make([]int, len(mergeArry))
	MergeSort(mergeArry, tmp, 0, len(mergeArry)-1)
	fmt.Printf("%v", mergeArry)
}

func MergeSort(a, b []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(a, b, left, mid)
		MergeSort(a, b, mid+1, right)
		merge(a, b, left, mid, right)
	}
}

func merge(a, b []int, left, mid, right int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right {
		if a[i] < a[j] {
			b[t] = a[i]
			i++
		} else {
			b[t] = a[j]
			j++
		}
		t++
	}

	for i <= mid {
		b[t] = a[i]
		t++
		i++
	}

	for j <= right {
		b[t] = a[j]
		t++
		j++
	}

	t = 0
	for left <= right {
		a[left] = b[t]
		t++
		left++
	}
}
