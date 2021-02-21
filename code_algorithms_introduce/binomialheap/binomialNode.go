package binomialheap

import (
	"math"
)

// 结点类型
type NodeType int

// 结点的最大/小值
var Maximum = NodeType(math.MaxInt32)
var Minimum = NodeType(math.MinInt32)

// binomialNode 二项堆结点
type binomialNode struct {
	key     NodeType      //关键字值
	degree  int           //当前结点的度
	parent  *binomialNode //父结点
	child   *binomialNode //孩子结点
	sibling *binomialNode //兄弟结点
}

// !+ 对二项树的操作

// binomialNodeMinimum 寻找最小关键字
// 返回一个指针，指向二项堆的最小关键字结点
func (H *binomialNode) binomialNodeMinimum() *binomialNode {
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

// binomialNodeUnion 合并两个二项堆
// Usage: h1=binomialNodeUnion(H2)
func (H1 *binomialNode) binomialNodeUnion(H2 *binomialNode) *binomialNode {
	//sortedH 已经将H1和H2根据度排序了
	newH := binomialHeapMerge(H1, H2)
	if newH == nil {
		// 合并两个空二项堆
		return newH
	}
	var prevX *binomialNode = nil
	x := newH
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

// BinomialHeapInsert 将结点x插入二项堆中
// 至少结点x的key应该是有的
func (H *binomialNode) binomialNodeInsert(x *binomialNode) *binomialNode {
	x.parent = nil
	x.sibling = nil
	x.child = nil
	x.degree = 0
	return H.binomialNodeUnion(x)
}

// binomialNodeDecreaseKey 减小关键字的值
// k 必须小于 x 当前的值
func (H *binomialNode) binomialNodeDecreaseKey(x *binomialNode, k NodeType) {
	if k >= x.key {
		return
	}
	x.key = k
	y := x
	z := x.parent
	for z != nil && y.key < z.key {
		y.key, z.key = z.key, y.key
		y = z
		z = y.parent
	}
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
func binomialHeapMerge(b1, b2 *binomialNode) *binomialNode {
	sortDegree := &binomialNode{}
	t := sortDegree
	for b1 != nil && b2 != nil {
		if b1.degree <= b2.degree {
			t.sibling = b1
			b1 = b1.sibling
		} else {
			t.sibling = b2
			b2 = b2.sibling
		}
		t = t.sibling
	}
	for b1 != nil {
		t.sibling = b1
		b1 = b1.sibling
		t = t.sibling
	}
	for b2 != nil {
		t.sibling = b2
		b2 = b2.sibling
		t = t.sibling
	}
	return sortDegree.sibling
}

// binomialGetChild 得到结点H的所有孩子,按照度的递增排序
func (b *binomialNode) binomialGetChild() *binomialNode {
	newH := &binomialNode{}
	child := b.child
	for child != nil {
		t := child.sibling
		child.parent = nil
		child.sibling = newH.sibling
		newH.sibling = child
		child = t
	}
	return newH.sibling
}
