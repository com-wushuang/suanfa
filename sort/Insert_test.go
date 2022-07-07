package sort

import (
	"fmt"
	"testing"
)

var insertArray = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestAssert(t *testing.T) {
	Insert(insertArray)
	fmt.Printf("%v", insertArray)
}

func Insert(a []int) {
	for i := 0; i < len(a)-1; i++ {  // 因为下一行的起始条件是j:=i+1,所以这里的终止条件是len(a)-1
		for j := i + 1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j-1],a[j]=a[j],a[j-1]
			}
		}
	}
}