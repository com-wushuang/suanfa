package sort

import (
	"fmt"
	"testing"
)

var shellArry = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestShell(t *testing.T) {
	Shell(shellArry)
	fmt.Printf("%v", shellArry)
}

func Shell(a []int) {
	length := len(a)
	for inc := length / 2; inc > 0; inc = inc / 2 {
		for i := inc; i < length; i++ {
			for j := i - inc; j >= 0; j = j - inc {
				if a[j] > a[i] {
					a[j], a[i] = a[i], a[j]
				}
			}
		}
	}
}
