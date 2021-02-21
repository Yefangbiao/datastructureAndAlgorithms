package binomialheap

import (
	"fmt"
	"testing"
)

var testA = []int{12, 7, 25, 15, 28, 33, 41}
var testB = []int{18, 35, 20, 42, 9, 31, 23, 6, 48, 11, 24, 52, 13}

func TestBinomialNode_BinomialHeapInsert(t *testing.T) {
	heapA := NewBinomialHeap()
	fmt.Println("print heapA")
	for _, v := range testA {
		heapA.BinomialHeapInsert(v)
	}
	heapA.Print()

	heapB := NewBinomialHeap()
	fmt.Println("print heapB")
	for _, v := range testB {
		heapB.BinomialHeapInsert(v)
	}
	heapB.Print()
}

func TestBinomialHeap_BinomialHeapMinimum(t *testing.T) {
	heapA := NewBinomialHeap()
	for _, v := range testA {
		heapA.BinomialHeapInsert(v)
	}
	minimum := heapA.BinomialHeapMinimum()
	fmt.Println(minimum.key)
}

func TestBinomialHeap_BinomialHeapExtractMin(t *testing.T) {
	heapA := NewBinomialHeap()
	for _, v := range testA {
		heapA.BinomialHeapInsert(v)
	}
	minimum := heapA.BinomialHeapExtractMin()
	fmt.Println(minimum.key)

	fmt.Println("print heapA after extract minimum")
	heapA.Print()
}

func TestBinomialHeap_BinomialHeapDeCreaseKey(t *testing.T) {
	heapA := NewBinomialHeap()
	for _, v := range testA {
		heapA.BinomialHeapInsert(v)
	}
	minimum := heapA.BinomialHeapMinimum()
	de := minimum.child.child
	heapA.BinomialHeapDeCreaseKey(de, 1)

	fmt.Println("print heapA after decrease a key")
	heapA.Print()
	fmt.Printf("the miminum after decrease a key:%v\n", heapA.BinomialHeapMinimum().key)
}

func TestBinomialHeap_BinomialHeapDelete(t *testing.T) {
	heapA := NewBinomialHeap()
	for _, v := range testA {
		heapA.BinomialHeapInsert(v)
	}
	minimum := heapA.BinomialHeapMinimum()
	de := minimum.child.child
	heapA.BinomialHeapDelete(de)

	fmt.Println("print heapA after delete a key")
	heapA.Print()
}
