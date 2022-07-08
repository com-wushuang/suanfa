package array

import (
	"fmt"
	"testing"
)

func TestBinSearch(t *testing.T) {
	a := []int{1, 2, 4, 6, 8, 12, 15, 15, 16, 16, 18, 19, 19, 33}
	fmt.Println(binarySearch(a, 13))
	fmt.Println(binarySearch(a, 15))

	fmt.Println(leftBound(a, 16))
	fmt.Println(leftBound(a, 17))
	fmt.Println(leftBound(a, 55))

	fmt.Println(rightBound(a, 16))
	fmt.Println(rightBound(a, 17))
	fmt.Println(rightBound(a, -1))

}

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

	if left > len(a)-1 {
		return -1
	}
	if a[left] == target {
		return left
	}
	return -1
}

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
	if left == 0 {
		return -1
	}
	if a[left-1] == target {
		return left - 1
	}
	return -1
}
