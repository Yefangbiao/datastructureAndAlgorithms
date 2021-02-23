package chapter19

import (
	"fmt"
)

type fibonacciHeap struct {
	min *node //指向具有最小关键字的树的结点
	n   int   //H中当前的结点数目
}

// NewFibonacciHeap 创建一个新的斐波那契堆
// 返回一个新的斐波那契堆，其中H.n=0  H.min=nil
func NewFibonacciHeap() *fibonacciHeap {
	return &fibonacciHeap{}
}

// !+
// FibHeapInsert 插入一个新的数据到斐波那契堆中
func (H *fibonacciHeap) FibHeapInsert(x int) {
	n := newNode(x)
	if H.min == nil {
		// case1: H.min==nil
		n.left = n
		n.right = n
		H.min = n
	} else {
		//case2:
		H.min.insertNodeLeft(n)
		if n.data < H.min.data {
			H.min = n
		}
	}
	H.n++
}

// FibHeapMinimum 返回堆中的最小结点的指针
func (H *fibonacciHeap) FibHeapMinimum() *node {
	return H.min
}

// FibHeapUnion 合并两个斐波那契堆H1和H2
// 简单的将根链表连接，然后确定新的最小结点
func (H1 *fibonacciHeap) FibHeapUnion(H2 *fibonacciHeap) {
	if H1.min == nil {
		H1.min = H2.min
	} else {
		t := H2.min
		for i := 0; i < H2.n; i++ {
			next := t.right
			t.right = nil
			t.left = nil
			H1.min.insertNodeLeft(t)
			t = next
		}
	}
	if H1.min == nil || (H2.min != nil && H2.min.data < H1.min.data) {
		H1.min = H2.min
	}

	H1.n += H2.n
}

// FibHeapExtractMin 抽取最小结点，返回最小结点的指针
func (H *fibonacciHeap) FibHeapExtractMin() *node {
	z := H.min
	if z != nil {
		H.n--
		t := z.child
		for i := 0; i < z.degree; i++ {
			next := t.right
			t.right = nil
			t.left = nil
			t.p = nil
			H.min.insertNodeLeft(t)
			t = next
		}

		z.remove()
		if z == z.right {
			H.min = nil
		} else {
			H.min = z.right
			H.consolidate()
		}

	}
	return z
}

func (H *fibonacciHeap) FibHeapDecreaseKey(x *node, k int) {
	if k > x.data {
		return
	}
	x.data = k
	y := x.p
	if y != nil && x.data < y.data {
		H.cut(x, y)
		H.cascadingCut(y)
	}
	if x.data < H.min.data {
		H.min = x
	}
}

func (H *fibonacciHeap) FibHeapDelete(x *node) {
	H.FibHeapDecreaseKey(x, MINNODENUM)
	H.FibHeapExtractMin()
}

func (H *fibonacciHeap) Print() {
	fmt.Printf("Print fibonacci heap, %d nodes\n", H.n)
	H.min.print("-")
}

// consolidate 合并根链表
func (H *fibonacciHeap) consolidate() {
	// 应该是某个对数函数，但在math中找不到以自己设置底的函数，就用最大来代替
	A := make([]*node, H.n+1)
	m := make(map[*node]bool)
	t := H.min
	for i := 0; i < H.n; i++ {
		m[t] = true
		next := t.right
		x := t
		d := x.degree
		for A[d] != nil {
			y := A[d]
			if x.data > y.data {
				x, y = y, x
			}
			H.fibHeapLink(y, x)
			A[d] = nil
			d += 1
		}
		t = next
		A[d] = x
		if m[t] {
			break
		}
	}

	H.min = nil
	for i := 0; i <= H.n; i++ {
		if A[i] != nil {
			if H.min == nil {
				H.min = A[i]
			} else {
				if A[i].data < H.min.data {
					H.min = A[i]
				}
			}
		}
	}
}

// fibHeapLink 让y成为x的孩子，将.degree加1
func (H *fibonacciHeap) fibHeapLink(y *node, x *node) {
	y.remove()
	if x.child == nil {
		x.child = y
		y.left = y
		y.right = y
	} else {
		x.child.insertNodeLeft(y)
	}
	y.p = x
	x.degree++
	y.mark = false
}

func (H *fibonacciHeap) cut(x *node, y *node) {
	otherChild := x.left
	x.remove()
	if x == x.left {
		y.child = nil
	} else {
		y.child = otherChild
	}
	y.degree--
	x.p = nil
	H.min.insertNodeLeft(x)
	x.mark = false

}

func (H *fibonacciHeap) cascadingCut(y *node) {
	z := y.p
	if z != nil {
		if y.mark == false {
			y.mark = true
		} else {
			H.cut(y, z)
			H.cascadingCut(z)
		}
	}
}

// !-
