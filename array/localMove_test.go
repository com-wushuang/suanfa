package array

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	sortArry := []int{0, 0, 1, 1, 2, 2, 2, 3, 4, 5, 5}
	newLen := moveDuplicate(sortArry)
	fmt.Printf("%d,%v", newLen, sortArry)
	fmt.Println()

	a := []int{1, 2, 3, 0, 4, 5, 0, 6, 0, 7, 8, 9, 0, 0}
	newLength := moveElement(a, 0)
	fmt.Printf("%d,%v", newLength, a)
	fmt.Println()

	b := []int{1, 2, 3, 0, 4, 5, 0, 6, 0, 7, 8, 9, 0, 0}
	moveZero(b)
	fmt.Printf("%v", b)
}

// 移除有序数组中重复的元素,返回新数组的长度
func moveDuplicate(a []int) int {
	slow, fast := 0, 0
	for fast < len(a) {
		if a[slow] != a[fast] {
			slow++
			a[slow] = a[fast]
		}
		fast++
	}
	return slow + 1
}

// 将指定的元素移除，返回新数组的长度，不必考虑新数组超出部分的元素
func moveElement(a []int, val int) int {
	slow, fast := 0, 0
	for fast < len(a) {
		if a[fast] != val { // 不遇到val时,快慢指针一起移动
			a[slow] = a[fast]
			slow++
		}
		fast++
	}

	return slow
}

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
