package sort

import (
	"fmt"
	"testing"
)

var selectArray = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestSelect(t *testing.T) {
	Select(selectArray)
	fmt.Printf("%v", selectArray)
}

func Select(a []int) {
	for i := 0; i < len(a); i++ {
		minIndex := i
		for j := i; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i],a[minIndex]=a[minIndex],a[i]
	}
}
