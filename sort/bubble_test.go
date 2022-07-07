package sort

import (
	"fmt"
	"testing"
)

var bubbleArray = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestBubble(t *testing.T) {
	Bubble(bubbleArray)
	fmt.Printf("%v", bubbleArray)
}

func Bubble(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
}
