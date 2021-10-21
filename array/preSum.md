## 前缀和数组
前缀和数组是通过原始数组推倒出的新的数组
### 定义
`a`为原数组，`sum`为前缀和数组，`sum[i]= a[0]+a[1]+ ...+ a[i-1]` , 特别的`sum[0]=0`
```go
    a:=[]int{1,2,3,4} // 原始数组
    
    // 前缀和数组
    sum := make([]int, len(a)+1)  
    sum[0] = 0
    for i := 0; i < len(a); i++ {
        sum[i+1] = sum[i] + a[i]
    }
```
- 如果原始数组`a`的长度为`n`，那么前缀和数组的长度为`n+1`.
- 关键代码：`sum[i+1] = sum[i] + a[i]` ，如果将 `sum[i]= a[0]+...+a[i-1]` 带入，那么就是`sum[i+1]=a[0]+...+a[i]`符合前面的定义

### 例子
给定一个数组，求解和为 k 的连续子数组的个数
```go
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
```
- 原理：`a[i..j]`的和，只需要 `sum[j+1]-sum[i]`（通过定义的公式带入即可证明，和上面的关键代码证明方法类似）
- 注意遍历子数组的代码结构