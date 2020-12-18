package sorts

// 我们说堆默认是最大堆

// MaxHeap 最大堆
type maxHeap struct {
	slice    []int
	heapSize int
}

func (heap *maxHeap) buildHeap() {
	for i := heap.heapSize / 2; i >= 0; i-- {
		heap.heapify(i)
	}
}

// 对堆十分重要的调整函数
func (heap *maxHeap) heapify(i int) {
	l := 2*i + 1
	r := 2*i + 2
	max := i

	if l < heap.heapSize && heap.slice[max] < heap.slice[l] {
		max = l
	}
	if r < heap.heapSize && heap.slice[max] < heap.slice[r] {
		max = r
	}
	if max != i {
		heap.slice[i], heap.slice[max] = heap.slice[max], heap.slice[i]
		heap.heapify(max)
	}
}

func HeapSort(arr []int) []int {
	h := maxHeap{slice: arr, heapSize: len(arr)}
	h.buildHeap()
	for i := len(arr) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.heapify(0)
	}
	return h.slice
}
