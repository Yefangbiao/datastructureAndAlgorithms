package redblacktree

import (
	"fmt"
	"log"
)

// RBTree 红黑树
type RBTree struct {
	Root *rbNode
}

func (tree *RBTree) InsertNode(key int) {
	//如果是根节点
	if tree.Root == nil {
		tree.Root = NewRBNode(key)
		tree.Root.color = BLACK

		tree.Root.parent = tNil
	} else {
		tree.insertNode(tree.Root, key)
	}
}

// insertNode 红黑树插入算法
func (tree *RBTree) insertNode(node *rbNode, key int) {

	if key < node.key {
		if node.left != tNil {
			tree.insertNode(node.left, key)
		} else {
			temp := NewRBNode(key)
			node.left = temp
			temp.parent = node

			// insert fix
			tree.insertRepairTree(temp)
		}
	} else {
		if node.right != tNil {
			tree.insertNode(node.right, key)
		} else {
			temp := NewRBNode(key)
			node.right = temp
			temp.parent = node

			// insert fix
			tree.insertRepairTree(temp)
		}
	}
}

// 下面两个函数特别强调了根节点的情况
func (tree *RBTree) leftRotate(node *rbNode) {
	//如果是左旋
	temp, err := node.LeftRotate()
	if err != nil {
		fmt.Println("left rotate is wrong")
	}
	if temp.GetParent() == tNil {
		//小心根节点情况
		tree.Root = temp
		temp.parent = tNil
	}
}

func (tree *RBTree) rightRotate(node *rbNode) {
	//如果是右旋
	temp, err := node.RightRotate()
	if err != nil {
		fmt.Println("right rotate is wrong")
	}
	if temp.GetParent() == tNil {
		//小心根节点情况
		tree.Root = temp
		temp.parent = tNil
	}
}

// insertRepairTree插入的修复程序
func (tree *RBTree) insertRepairTree(node *rbNode) {

	if node.GetParent() == tNil {
		//case1: node is Root
		tree.Root = node
		node.color = BLACK
		tree.Root.parent = tNil
		return
	} else if node.GetParent().color == BLACK {
		//case2: node's parent is black,do nothing
		return
	} else if node.GetUncle() != tNil && node.GetUncle().color == RED {
		//case3: node's uncle is red
		//父节点和叔结点转换为黑色，祖父节点变成红色。解决了node的问题，但同时，祖父节点可能违反了红黑树的性质
		node.GetParent().color = BLACK
		node.GetUncle().color = BLACK
		node.GetGrandParent().color = RED
		tree.insertRepairTree(node.GetGrandParent())
	} else {
		//case4: node's uncle is black
		//1.预处理
		p := node.GetParent()
		g := node.GetGrandParent()
		if node == p.right && p == g.left {
			tree.leftRotate(p)
			node = node.left
		} else if node == p.left && p == g.right {
			tree.rightRotate(p)
			node = node.right
		}
		//final process
		//2.最终处理
		p = node.GetParent()
		g = node.GetGrandParent()
		if node == p.left {
			tree.rightRotate(g)
		} else {
			tree.leftRotate(g)
		}
		p.color = BLACK
		g.color = RED
	}
}

// 对外删除算法,成功返回true，失败返回false
func (tree *RBTree) DeleteNode(key int) bool {
	// 如果根节点为空,错误
	if tree.Root == tNil || tree.Root == nil {
		return false
	}
	return tree.deleteChild(tree.Root, key)
}

// 红黑树删除一个结点辅助函数
func (tree *RBTree) rbTransplant(u, v *rbNode) {
	// u一定是存在的,v可能是nil
	if u.parent == tNil {
		//代表是根节点
		tree.Root = v
		//否则父结点一定存在
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

// deleteChild 红黑树的实际删除算法
func (tree *RBTree) deleteChild(node *rbNode, key int) bool {
	if key < node.key {
		// 向左子树找
		if node.left == tNil {
			return false
		}
		return tree.deleteChild(node.left, key)
	} else if key > node.key {
		// 向右子树寻找
		if node.right == tNil {
			return false
		}
		return tree.deleteChild(node.right, key)
	}

	// 否则，找到了删除的结点
	// node:被删结点 y:实际被删除结点 x:拼接结点
	y := node
	yColor := y.color
	var x *rbNode

	if node.left == tNil {
		//case1:左子树为空
		x = node.right
		tree.rbTransplant(node, node.right)
	} else if node.right == tNil {
		x = node.left
		tree.rbTransplant(node, node.left)
	} else {
		y = node.GetRightMin()
		//y 一定存在
		yColor = y.color
		//x 不一定存在
		x = y.right
		if y.parent == node {
			x.parent = y
		} else {
			tree.rbTransplant(y, y.right)
			y.right = node.right
			y.right.parent = y
		}
		tree.rbTransplant(node, y)
		y.left = node.left
		y.left.parent = y
		y.color = node.color
	}
	if yColor == BLACK {
		tree.rbFixUp(x)
	}

	return true
}

func (tree *RBTree) rbFixUp(x *rbNode) {
	for x != tree.Root && x.color == BLACK {
		// x是左子树
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED {
				//case1
				w.color = BLACK
				x.parent.color = RED
				tree.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				//case2
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK {
					//case3
					w.left.color = BLACK
					w.color = RED
					tree.rightRotate(w)
					w = x.parent.right
				}
				//case4 所有情况转换成情况4
				w.color = x.parent.color
				x.parent.color = BLACK
				w.right.color = BLACK
				tree.leftRotate(x.parent)
				x = tree.Root
			}

		} else {
			//x是右边
			w := x.parent.left
			if w.color == RED {
				//case1
				w.color = BLACK
				x.parent.color = RED
				tree.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == BLACK && w.left.color == BLACK {
				//case2
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					tree.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				w.left.color = BLACK
				tree.rightRotate(x.parent)
				x = tree.Root

			}
		}
	}
	x.color = BLACK
}

// log输出树
func printTreeInLog(n *rbNode, front string) {
	if n != tNil {
		var colorstr string
		if n.color == RED {
			colorstr = "红"
		} else {
			colorstr = "黑"
		}
		log.Printf(front+"%d,%s\n", n.key, colorstr)
		printTreeInLog(n.left, front+"-(l)|")

		printTreeInLog(n.right, front+"-(r)|")
	}
}
