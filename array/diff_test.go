package array

import (
	"fmt"
	"testing"
)

func TestDiff(t *testing.T) {
	bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	n := 5
	fmt.Printf("%v", corpFlightBookings(bookings, n))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, 5) // 原始数组的差分数组
	for _, booking := range bookings {

		// 区间和加数
		left := booking[0] - 1 //题目中的航班是从0编号的
		right := booking[1]
		k := booking[2]

		diff[left] += k
		if right < n {
			diff[right] -= k
		}
	}

	// 差分数组求前缀和得到原始数组
	res := make([]int, 5)
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}
