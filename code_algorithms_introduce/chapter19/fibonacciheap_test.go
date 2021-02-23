package chapter19

import (
	"fmt"
	"testing"
)

var testA = []int{23, 7, 3, 17, 24}
var testB = []int{43, 5}

func TestFibonacciHeap_FibHeapInsert(t *testing.T) {
	heap := NewFibonacciHeap()
	for _, v := range testA {
		heap.FibHeapInsert(v)
	}
	heap.Print()
}

func TestFibonacciHeap_FibHeapUnion(t *testing.T) {
	heapA := NewFibonacciHeap()
	for _, v := range testA {
		heapA.FibHeapInsert(v)
	}
	heapB := NewFibonacciHeap()
	for _, v := range testB {
		heapB.FibHeapInsert(v)
	}
	heapB.Print()
	heapA.FibHeapUnion(heapB)
	heapA.Print()
}

func TestFibonacciHeap_FibHeapExtractMin(t *testing.T) {
	heap := NewFibonacciHeap()
	for _, v := range testA {
		heap.FibHeapInsert(v)
	}
	min := heap.FibHeapExtractMin()
	fmt.Println(min.data)
	heap.Print()
}

func TestFibonacciHeap_FibHeapDecreaseKey(t *testing.T) {
	heap := NewFibonacciHeap()
	for _, v := range testA {
		heap.FibHeapInsert(v)
	}
	heap.FibHeapExtractMin()
	x := heap.min.child
	heap.FibHeapDecreaseKey(x, 1)
	heap.Print()
}

func TestFibonacciHeap_FibHeapDelete(t *testing.T) {
	heap := NewFibonacciHeap()
	for _, v := range testA {
		heap.FibHeapInsert(v)
	}
	heap.FibHeapExtractMin()
	x := heap.min.child
	heap.FibHeapDelete(x)
	heap.Print()
}
