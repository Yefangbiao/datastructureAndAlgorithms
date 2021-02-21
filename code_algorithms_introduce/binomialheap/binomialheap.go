package binomialheap

import "fmt"

type binomialHeap struct {
	*binomialNode
}

// !+ 对外暴露的方法

// newBinomialHeap 一个新的Heap
func NewBinomialHeap() *binomialHeap {
	return &binomialHeap{}
}

// BinomialHeapInsert 插入一个值
func (b *binomialHeap) BinomialHeapInsert(x int) {
	insertNode := &binomialNode{key: NodeType(x)}
	b.binomialNode = b.binomialNodeInsert(insertNode)
}

// BinomialHeapMinimum 寻找最小关键字
// 返回一个指针，指向二项堆的最小关键字结点
func (b *binomialHeap) BinomialHeapMinimum() *binomialNode {
	return b.binomialNodeMinimum()
}

// BinomialHeapExtractMin 抽取最小值
// 返回一个结点的指针
func (b *binomialHeap) BinomialHeapExtractMin() *binomialNode {
	// 获取min的子结点
	min := b.binomialNodeMinimum()
	minChild := min.binomialGetChild()

	b.removeMinfromH(min)
	// 新的H
	b.binomialNode = b.binomialNode.binomialNodeUnion(minChild)

	min.child = nil
	min.sibling = nil
	return min
}

func (b *binomialHeap) removeMinfromH(min *binomialNode) {
	// 把min从H中移除
	var prev *binomialNode
	t := b.binomialNode
	for t != min {
		prev = t
		t = t.sibling
	}
	if prev == nil {
		b.binomialNode = b.sibling
	} else {
		prev.sibling = min.sibling
	}
}

// BinomialHeapDeCreaseKey 减小关键字的值
func (b *binomialHeap) BinomialHeapDeCreaseKey(x *binomialNode, k int) {
	b.binomialNodeDecreaseKey(x, NodeType(k))
}

// BinomialHeapDelete 删除某个结点
func (b *binomialHeap) BinomialHeapDelete(x *binomialNode) {
	b.BinomialHeapDeCreaseKey(x, int(Minimum))
	b.BinomialHeapExtractMin()
}

// 打印
func (b *binomialHeap) Print() {
	binomialPrint(b.binomialNode)
}

// !-

// !+ 打印辅助方法

func binomialPrint(heap *binomialNode) {
	if heap == nil {
		return
	}

	p := heap
	fmt.Printf("== 二项堆( ")
	for p != nil {
		fmt.Printf("B%d ", p.degree)
		p = p.sibling
	}

	fmt.Printf(")的详细信息：\n")

	i := 0
	for heap != nil {
		i++
		fmt.Printf("%d. 二项树B%d: \n", i, heap.degree)
		fmt.Printf("\t%2d(%d) is root\n", heap.key, heap.degree)
		binomialPrintChild(heap.child, heap, 1)
		heap = heap.sibling

	}
	fmt.Println()
}

func binomialPrintChild(node, prev *binomialNode, direction int) {

	for node != nil {
		//printf("%2d \n", node->key);
		if direction == 1 {
			fmt.Printf("\t%2d(%d) is %2d's child\n", node.key, node.degree, prev.key)
		} else {
			fmt.Printf("\t%2d(%d) is %2d's next\n", node.key, node.degree, prev.key)
		}

		if node.child != nil {
			binomialPrintChild(node.child, node, 1)
		}

		// 兄弟节点
		prev = node
		node = node.sibling
		direction = 2
	}

}

// !-
