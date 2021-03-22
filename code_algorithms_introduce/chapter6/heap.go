package chapter6

// heap 堆
type heap struct {
	nums     []int //存储堆的数组
	heapSize int   //堆的大小
	compare  bool  //false:最小堆,true:最大堆,默认最小
}

// 堆排序
func SortByHeap(nums []int, heapSize int, compare bool) []int {
	hp := newHeap(nums, heapSize, compare)
	for i := len(nums) - 1; i >= 1; i-- {
		hp.nums[0], hp.nums[i] = hp.nums[i], hp.nums[0]
		hp.heapSize--
		hp.down(0)
	}
	return hp.nums
}

func newHeap(nums []int, heapSize int, compare bool) *heap {
	return &heap{nums: nums, heapSize: heapSize, compare: compare}
}

func (hp *heap) Push(x int) {
	hp.nums = append(hp.nums, x)
	hp.heapSize++
	hp.up(hp.heapSize - 1)
}

func (hp *heap) Pop() int {
	x := hp.nums[0]
	hp.nums[0], hp.nums[hp.heapSize-1] = hp.nums[hp.heapSize-1], hp.nums[0]
	hp.heapSize--
	hp.down(0)
	return x
}

// i: 下标
func (hp *heap) down(i int) {
	left := i*2 + 1
	right := i*2 + 2
	index := i
	if left < hp.heapSize && hp.less(left, index) {
		index = left
	}
	if right < hp.heapSize && hp.less(right, index) {
		index = right
	}
	if index != i {
		hp.nums[index], hp.nums[i] = hp.nums[i], hp.nums[index]
		hp.down(index)
	}
}

func (hp *heap) up(i int) {
	for i > 0 && hp.less(i, i/2) {
		hp.nums[i], hp.nums[i/2] = hp.nums[i/2], hp.nums[i]
		i = i / 2
	}
}

func (hp *heap) less(i, j int) bool {
	if hp.compare {
		return hp.nums[i] > hp.nums[j]
	}
	return hp.nums[i] < hp.nums[j]
}
