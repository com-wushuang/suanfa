package array

import (
	"fmt"
	"testing"
)

func TestPreSum(t *testing.T) {
	a := []int{1, 1, 1}
	fmt.Printf("%d", preSum(a, 2))

	fmt.Println()

	b := []int{1, 2, 3}
	fmt.Printf("%d", preSum(b, 3))

}

func preSum(a []int, k int) (res int) {
	// 构造前缀和数组
	sum := make([]int, len(a)+1)
	sum[0] = 0
	for i := 0; i < len(a); i++ {
		sum[i+1] = sum[i] + a[i]
	}

	// 求解和为 k 的连续子数组的个数
	for i := 1; i < len(sum); i++ {
		for j := 0; j < i; j++ {
			if sum[i]-sum[j] == k {
				res++
			}
		}
	}
	return
}
