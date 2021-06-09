package chapter16

import (
	"container/heap"
	"log"
)

func Huffman(chars char) *tree {
	n := chars.Len()

	pq := make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		newPQ := &tree{
			ch:   chars.ch[i],
			freq: chars.freq[i],
		}
		pq[i] = newPQ
	}
	heap.Init(&pq)

	for i := 1; i <= n-1; i++ {
		z := &tree{}
		x := heap.Pop(&pq).(*tree)
		y := heap.Pop(&pq).(*tree)
		z.left = x
		z.right = y
		z.freq = x.freq + y.freq
		heap.Push(&pq, z)
	}

	return heap.Pop(&pq).(*tree)
}

// char 字符集,包含一个字符和频率
type char struct {
	ch   []byte
	freq []int
}

func newChar() *char {
	return &char{
		ch:   []byte{},
		freq: []int{},
	}
}

func (c *char) Len() int {
	return len(c.ch)
}

func (c *char) Insert(ch byte, f int) {
	c.ch = append(c.ch, ch)
	c.freq = append(c.freq, f)
}

type tree struct {
	ch          byte
	freq        int
	left, right *tree
}

type PriorityQueue []*tree

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	//n := len(*pq)
	item := x.(*tree)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func printTree(t *tree, pre string) {
	if t == nil {
		return
	}
	printTree(t.left, pre+" 0 ")
	if t.left == nil && t.right == nil {
		log.Printf("%c: %s", t.ch, pre)
	}
	printTree(t.right, pre+" 1 ")
}
