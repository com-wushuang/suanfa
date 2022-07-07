package sort

import (
	"fmt"
	"testing"
)

var heapArry = []int{8, 1, 5, 7, 2, 3, 9, 6, 4}

func TestHeap(t *testing.T) {
	HeapSort(heapArry, len(heapArry))
	fmt.Printf("%v", heapArry)
}

func HeapSort(a []int, n int) {
	for i := n/2 - 1; i >= 0; i-- { // 从最后一个节点(n-1)的父节点开始
		heapify(a, i, n)
	}

	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, 0, i) // 因为第0个节点被调换了位置，那么需要调整，保持为大顶堆
	}
}

func heapify(a []int, i, n int) {
	largest := i
	lson := 2*i + 1 // 堆的性质
	rson := 2*i + 2 // 堆的性质

	// 选择出i节点和左右子节点中哪个位置的值最大
	if lson < n && a[largest] < a[lson] {
		largest = lson
	}

	if rson < n && a[largest] < a[rson] {
		largest = rson
	}

	// 如果最大的节点不是i节点，那么需要将i节点的值和largest节点的值做调换
	if largest != i {
		a[largest], a[i] = a[i], a[largest] // go快速交换元素的语法糖
		// 递归调用，使得largest节点也满足大顶堆的性质(因为调整过后largest节点可能不满足大顶堆性质)
		heapify(a, largest, n)
	}
}
