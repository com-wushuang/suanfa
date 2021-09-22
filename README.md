### 堆排序
堆是一种近似完全二叉树,下标为i的节点的(从0开始),有如下的性质:
- 父节点是 (i-1)/2
- 左孩子的节点是 i*2 + 1
- 右孩子的节点是 i*2 + 2
- 大(小)顶堆的父节点总是大(小)于子节点

首先需要调整初始化数组中的元素位置来构造一个大顶堆(以大顶堆为例)，在构造的过程中，我们就是需要用到上面的4个性质。
- heapify 是用来调整下标为i的节点，使得以其为根节点的树是一个大顶堆
- 对于整个数组，最后一个节点是`len(a)-1`，他的父节点是`len(a)/2-1`(计算方法参考性质1)，整个二叉树的调整就是从该节点开始的
```go
package sort

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
```
当我们把原始的数组构造成一个大顶堆之后，那我们就可以进行堆排序，排序的过程是
- 将堆顶元素(最大的元素，第0个节点)，和最后一个节点交换
- 排除最后一个元素(已经选出的最大元素)，然后调整第0个节点
- 循环上述过程，就完成了排序

### 快速排序
- 快速排序最重要的函数就是分区函数`partition`
- 分区函数的目的是，找到`pivot`在序列中的正确位置，使得左边的元素都小于它，右边的元素都大于它
- 选出最后一个元素作为`pivot`
- 先利用一个左指针从左向右扫描序列，当左指针指向的元素`a[i]`大于`pivot`时停止扫描(这意味着`a[i]`元素的位置不正确需要调换)
- 再利用一个右指针从右向左扫描序列，当右指针指向的元素`a[j]`小于`pivot`时停止扫描(这意味着`a[j]`元素的位置不正确需要调换)
- 调换`a[i]`,`a[j]`
- 继续按照上述逻辑进行扫描，终止的条件是i=j，此时就到了`pivot`再序列中的位置
- 将`pivot`元素放置在i的位置

```go
package sort

func partition(a []int, low, high int) int {
	pivot := high
	i := low  // 左指针
	j := high // 右指针

	for i != j {
		for i < j && a[i] <= a[pivot] { // 从左边开始扫描
			i++
		}
		for i < j && a[j] >= a[pivot] {
			j--
		}
		a[i], a[j] = a[j], a[i] // 交换元素
	}
	a[i], a[pivot] = a[pivot], a[i] // pivot放置在最终的位置上
	return i
}
```

### 归并排序
归并排序最重要的函数是merge函数,相较于堆排序和快速排序,归并排序是相对简单的。
- 合并左右序列到临时数组b中
- 左右序列各一个扫描指针i,j
- 如果a[i]小于a[j],则b[t]=a[i],i++;否则b[t]=a[j],j++
- 移动b数组的指针,t++
- 当左序列已经扫描完毕，那么将右序列直接拷贝到临时数组
- 当右序列已经扫描完毕，那么将左序列直接拷贝到临时数组
- 将排序完的元素拷贝到原来的数组a中
```go
package sort

func merge(a, b []int, left, mid, right int) {
	i := left
	j := mid + 1
	t := 0
	for i <= mid && j <= right { // 左右序列扫描
		if a[i] < a[j] {
			b[t] = a[i]
			i++
		} else {
			b[t] = a[j]
			j++
		}
		t++
	}

	for i <= mid { // 当右序列已经扫描完毕，那么将左序列直接拷贝到临时数组
		b[t] = a[i]
		t++
		i++
	}

	for j <= right { // 当左序列已经扫描完毕，那么将右序列直接拷贝到临时数组
		b[t] = a[j]
		t++
		j++
	}

	t = 0
	for left <= right { // 将排序完的元素拷贝到原来的数组a中 ,注意拷贝到a数组中的位置
		a[left] = b[t]
		t++
		left++
	}
}

```

### 希尔排序
希尔排序是插入排序算法的一种改进,将序列按照一定的间隔分组,然后对每一组使用简单的插入排序算法
- 间隔每次缩小一半(通常情况下，不是最优的)
- 间隔越大原序列被分成的子序列越多
- 比如原数组有10个数,间隔如果是5,那么原序列会被分成5个子序列
- 固定间隔的一次循环中,直接对所有的子序列进行插入排序
```go
package sort

func Shell(a []int) {
	length := len(a)
    
    // 间隔每次缩小一半
	for inc := length / 2; inc > 0; inc = inc / 2 {
        
		// 对所有的子序列进行插入排序
		for i := inc; i < length; i++ {
            
            // 单个子序列插入排序
			for j := i - inc; j >= 0; j = j - inc {
				if a[j] > a[i] {
					a[j], a[i] = a[i], a[j]
				}
			}
		}
	}
}
```