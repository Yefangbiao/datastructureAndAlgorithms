package redblacktree

import "errors"

const (
	//红节点设置为true
	RED bool = true
	//黑节点设置为false
	BLACK bool = false
)

// rbNode 红黑树
type rbNode struct {
	key                 int
	color               bool
	left, right, parent *rbNode
}

// 定义一个tNil结点
var (
	tNil *rbNode = &rbNode{color: BLACK}
)

// NewRBNode 新建一个结点，默认Color为红色
func NewRBNode(key int) *rbNode {
	return &rbNode{color: RED, key: key, left: tNil, right: tNil}
}

// some helper-function

// GetParent 获取父亲节点
func (rbNode *rbNode) GetParent() *rbNode {
	if rbNode.parent == nil {
		return nil
	}
	return rbNode.parent
}

// GetGrandParent 获取祖父节点
func (rbNode *rbNode) GetGrandParent() *rbNode {
	return rbNode.GetParent().GetParent()
}

// GetSibling获取兄弟节点
func (rbNode *rbNode) GetSibling() *rbNode {
	p := rbNode.GetParent()
	if p == nil {
		return nil
	}
	if p.left == rbNode {
		return p.right
	} else {
		return p.left
	}
}

// GetUncle 获取父节点的兄弟节点
func (rbNode *rbNode) GetUncle() *rbNode {
	gp := rbNode.GetParent()
	if gp == nil {
		return nil
	}
	return gp.GetSibling()
}

// LeftRotate 左旋
func (x *rbNode) LeftRotate() (*rbNode, error) {
	y := x.right
	p := x.GetParent()
	if y == nil {
		return x, errors.New("左旋右结点不能为空")
	}
	x.right = y.left
	y.left = x
	x.parent = y

	// Handle other child/parent pointers.
	if x.right != nil {
		x.right.parent = x
	}

	// Initially x could be the Root.
	if p != nil {
		if x == p.left {
			p.left = y
		} else {
			p.right = y
		}
	}
	y.parent = p

	return y, nil
}

// RightRotate 右旋
func (y *rbNode) RightRotate() (*rbNode, error) {
	x := y.left
	p := y.GetParent()
	if x == nil {
		return y, errors.New("右旋左结点不能为空")
	}
	y.left = x.right
	x.right = y
	y.parent = x

	// Handle other child/parent pointers.
	if y.left != nil {
		y.left.parent = y
	}

	// Initially x could be the Root.
	if p != nil {
		if y == p.left {
			p.left = x
		} else {
			p.right = x
		}
	}
	x.parent = p
	return x, nil
}

// 获取rbNode右子树的最小结点
func (rbNode *rbNode) GetRightMin() *rbNode {
	rbNode = rbNode.right
	if rbNode == tNil {
		return tNil
	}
	for rbNode.left != tNil {
		rbNode = rbNode.left
	}
	return rbNode
}
