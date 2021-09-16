package sort

import (
	"fmt"
	"testing"
)

var assertArry = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestAssert(t *testing.T) {
	Assert(assertArry)
	fmt.Printf("%v", assertArry)
}

func Assert(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if a[j] < a[j-1] {
				tmp := a[j-1]
				a[j-1] = a[j]
				a[j] = tmp
			}
		}
	}
}
