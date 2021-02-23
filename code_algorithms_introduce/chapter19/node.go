package chapter19

import (
	"fmt"
	"math"
)

const MINNODENUM = math.MinInt32

// node 结点
type node struct {
	p           *node //指向父结点的指针
	child       *node //指向孩子结点的指针
	left, right *node //指向左右兄弟的指针
	degree      int   //结点的度，即孩子数目
	mark        bool  //结点从上一次成为另一个结点孩子后，是否失去过孩子
	data        int   //结点的数据
}

// newNode 返回一个新的结点
func newNode(data int) *node {
	return &node{
		p:      nil,
		child:  nil,
		left:   nil,
		right:  nil,
		degree: 0,
		mark:   false,
		data:   data,
	}
}

// !+
// insertNodeLeft 插入一个新的结点到x结点的左边
func (n *node) insertNodeLeft(x *node) {
	if x == nil {
		return
	}
	left := n.left
	left.right = x
	x.left = left
	x.right = n
	n.left = x
}

// print 辅助打印函数
func (n *node) print(sep string) {
	if n == nil {
		return
	}
	m := make(map[*node]bool)
	t := n
	for {
		if !m[t] {
			m[t] = true
			fmt.Printf("%s key:%d, degree:%d\n", sep, t.data, t.degree)
			t.child.print(sep + sep)
			t = t.right
		} else {
			break
		}
	}
}

func (n *node) remove() {
	n.left.right = n.right
	n.right.left = n.left
}

// !-
