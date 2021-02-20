package binomialheap

import "math"

// 结点类型
type NodeType int

// 结点的最大/小值
var Maximum = NodeType(math.MaxInt32)
var Minimum = NodeType(math.MinInt32)

// binomialNode 二项堆结点
// 二项堆结点也可以当做head[H]
type binomialNode struct {
	key     NodeType      //关键字值
	degree  int           //当前结点的度
	parent  *binomialNode //父结点
	child   *binomialNode //孩子结点
	sibling *binomialNode //兄弟结点
}

// !+ 对二项堆的操作

// MakeBinomialHeap 创建一个新的二项堆
// warning:nil
func MakeBinomialHeap() *binomialNode {
	var head *binomialNode //变量逃逸
	return head
}

// BinomialHeapMinimum 寻找最小关键字
// 返回一个指针，指向二项堆的最小关键字结点
func (H *binomialNode) BinomialHeapMinimum() *binomialNode {
	var y *binomialNode
	x := H
	min := Maximum
	for x != nil {
		if min > x.key {
			min = x.key
			y = x
		}
		x = x.sibling
	}

	return y
}

// BinomialHeapUnion 合并两个二项堆
// Usage: h1=BinomialHeapUnion(H2)
func (H1 *binomialNode) BinomialHeapUnion(H2 *binomialNode) *binomialNode {
	newH := MakeBinomialHeap()
	//sortedH 已经将H1和H2根据度排序了
	sortedH := binomialHeapMerge(H1, H2)
	if sortedH == nil {
		// 合并两个空二项堆
		return newH
	}
	var prevX *binomialNode = nil
	x := sortedH
	nextX := x.sibling
	for nextX != nil {
		if (x.degree != nextX.degree) || (nextX.sibling != nil && nextX.sibling.degree == x.degree) {
			//case 1 and 2
			prevX = x
			x = nextX
		} else if x.key <= nextX.key {
			x.sibling = nextX.sibling
			binomialLink(nextX, x)
		} else {
			if prevX == nil {
				newH = nextX
			} else {
				prevX.sibling = nextX
			}
			binomialLink(x, nextX)
			x = nextX
		}
		nextX = x.sibling
	}
	return newH
}

// BinlmialHeapInsert 将结点x插入二项堆中
// 至少结点x的key应该是有的
func (H *binomialNode) BinlmialHeapInsert(x *binomialNode) {
	x.parent = nil
	x.sibling = nil
	x.child = nil
	x.degree = 0
	H.BinomialHeapUnion(x)
}

// BinomialHeapExtractMin 抽取最小结点的关键字，返回一个该结点的指针
func (H *binomialNode) BinomialHeapExtractMin() *binomialNode {
	// 获取min的子结点
	min := H.BinomialHeapMinimum()
	minChild := min.binomialGetChild()
	// 把min从H中移除
	var prev *binomialNode
	t := H
	for t != nil {
		prev = t
		t = t.sibling
	}
	if prev == nil {
		H = H.sibling
	} else {
		prev.sibling = min.sibling
	}
	// 新的H
	H = H.BinomialHeapUnion(minChild)

	min.child = nil
	min.sibling = nil
	return min
}

// BinomialHeapDecareaseKey 减小关键字的值
// k 必须小于 x 当前的值
func (H *binomialNode) BinomialHeapDecareaseKey(x *binomialNode, k NodeType) {
	if k >= x.key {
		return
	}
	x.key = k
	y := x
	z := x.parent
	for z != nil && y.key > z.key {
		y.key, z.key = z.key, y.key
		y = z
		z = y.parent
	}
}

func (H *binomialNode) BinomialHeapDelete(x *binomialNode) {
	H.BinomialHeapDecareaseKey(x, Minimum)
	H.BinomialHeapExtractMin()
}

// !-

// !+ 辅助函数
// binomialLink 连接度数相同的两个二项树
// z成为y的父结点
func binomialLink(y, z *binomialNode) {
	y.parent = z
	y.sibling = z.child
	z.child = y
	z.degree++
}

// binomialHeapMerge 将H1和H2的根表合并成一个按照度数的单调递增依次排序的链表
func binomialHeapMerge(H1, H2 *binomialNode) *binomialNode {
	sortDegree := &binomialNode{}
	t := sortDegree
	for H1 != nil && H2 != nil {
		if H1.degree <= H2.degree {
			t.sibling = H1
			H1 = H1.sibling
		} else {
			t.sibling = H2
			H2 = H2.sibling
		}
		t = t.sibling
	}
	return sortDegree.sibling
}

// binomialGetChild 得到结点H的所有孩子,按照度的递增排序
func (H *binomialNode) binomialGetChild() *binomialNode {
	newH := &binomialNode{}
	child := newH.child
	for child != nil {
		child.parent = nil
		child.sibling = newH.sibling
		newH.sibling = child
		child = child.sibling
	}
	return newH.sibling
}

// !-
