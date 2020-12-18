package binarytree

import (
	"fmt"
	"log"
)

// binarySearchTree 二叉搜索树

// 先定义结点
type treeNode struct {
	key         int
	left, right *treeNode
	parent      *treeNode
}

// BinArySearchTree 二叉搜索树
type BinArySearchTree struct {
	root *treeNode
}

// 新建结点
func newNode(newKey int) *treeNode {
	return &treeNode{key: newKey}
}

// 插入
func (bst *BinArySearchTree) Insert(key int) {
	var y *treeNode
	y = nil
	node := bst.root
	for node != nil {
		y = node
		if key <= node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	node = newNode(key)
	node.parent = y
	//如果是根节点
	if y == nil {
		bst.root = node
	} else if node.key <= y.key {
		y.left = node
	} else {
		y.right = node
	}
}

// 删除
// 删除辅助函数
func (bst *BinArySearchTree) transplant(u, v *treeNode) {
	//根节点情况
	if u.parent == nil {
		bst.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// 删除函数
func (bst *BinArySearchTree) Delete(key int) {
	z := bst.root
	//先找到要删除的
	for z != nil && z.key != key {
		if key < z.key {
			z = z.left
		} else {
			z = z.right
		}
	}
	if z == nil {
		return
	}
	if z.left == nil {
		bst.transplant(z, z.right)
	} else if z.right == nil {
		bst.transplant(z, z.left)
	} else {
		y, _ := z.right.nodeMin()
		if y.parent != z {
			bst.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		bst.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
	// 如果有重复的
	bst.Delete(key)
}

// 搜索
func (bst *BinArySearchTree) TreeSearch(key int) (*treeNode, bool) {
	return bst.root.searchRecursively(key)
}

// 查询实际函数(递归方法)
// 查到了返回true,否则返回false
func (node *treeNode) searchRecursively(key int) (*treeNode, bool) {
	if node == nil {
		return nil, false
	}
	if node.key == key {
		return node, true
	}
	if key < node.key {
		return node.left.searchRecursively(key)
	} else {
		return node.right.searchRecursively(key)
	}
}

// 搜索(迭代法)
func (node *treeNode) searchIteratively(key int) (*treeNode, bool) {
	root := node
	for root != nil {
		if root.key == key {
			return root, true
		}
		if key < root.key {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil, false
}

// TreeMax 返回最大值
func (bst *BinArySearchTree) TreeMax() (*treeNode, bool) {
	node := bst.root
	if node == nil {
		return nil, false
	} else {
		for node.right != nil {
			node = node.right
		}
	}
	return node, true
}

// TreeMin 返回最小值
func (bst *BinArySearchTree) TreeMin() (*treeNode, bool) {
	return bst.root.nodeMin()
}

func (node *treeNode) nodeMin() (*treeNode, bool) {
	tNode := node
	if tNode == nil {
		return nil, false
	} else {
		for tNode.left != nil {
			tNode = tNode.left
		}
	}
	return tNode, true
}

// 中序遍历二叉搜索树
func (bst *BinArySearchTree) Traverse() {
	bst.root.inorderTreeWalk()
}

// 遍历主函数
func (node *treeNode) inorderTreeWalk() {
	if node == nil {
		return
	}
	node.left.inorderTreeWalk()
	fmt.Printf("%d\n", node.key)
	node.right.inorderTreeWalk()
}

// 前驱后继
// 找一个结点的后继,前驱就是parent了
func (node *treeNode) successor() *treeNode {
	// 找一个结点的后继
	if node.right != nil {
		min, _ := node.right.nodeMin()
		return min
	}
	y := node.parent
	tNode := node
	for y != nil && y.right == tNode {
		tNode = y
		y = y.parent
	}
	return y
}

func printTreeInLog(n *treeNode, front string) {
	if n != nil {

		log.Printf(front+"%d,%s\n", n.key)
		printTreeInLog(n.left, front+"-(l)|")

		printTreeInLog(n.right, front+"-(r)|")
	}
}
