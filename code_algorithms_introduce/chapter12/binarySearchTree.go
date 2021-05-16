package chapter12

// binarySearchTree 二叉查找树
type binarySearchTree struct {
	root *node
}

// GetNewBinarySearchTree 得到一个空二叉查找树
func GetNewBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{
		root: NIL,
	}
}

// InorderTreeWalk 返回二叉查找树的中序遍历
func (tree *binarySearchTree) InorderTreeWalk() []int {
	// 迭代方法
	root := tree.root
	if root == NIL {
		return []int{}
	}
	ans := make([]int, 0)

	stack := make([]*node, 0)

	for root != NIL || len(stack) > 0 {

		for root != NIL {
			stack = append(stack, root)
			root = root.left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.val)
		root = root.right
	}

	return ans
}

// TreeSearch 查找具有给定关键字的结点
// 找到返回结点，否则返回NIL
func (tree *binarySearchTree) TreeSearch(val int) *node {
	x := tree.root
	for x != NIL {
		if x.val == val {
			return x
		}
		if x.val < val {
			x = x.right
		} else {
			x = x.left
		}
	}
	return NIL
}

// 返回二叉查找树给定结点最小值
func (tree *binarySearchTree) TreeMinimum(x *node) *node {
	//x := tree.root
	if x == NIL {
		return NIL
	}
	for x.left != NIL {
		x = x.left
	}
	return x
}

// 返回二叉查找树给定结点最大值
func (tree *binarySearchTree) TreeMaximum(x *node) *node {
	//x := tree.root
	if x == NIL {
		return NIL
	}
	for x.right != NIL {
		x = x.right
	}
	return x
}

// TreeInsert 根据val插入一个结点
func (tree *binarySearchTree) TreeInsert(val int) {
	newNode := &node{
		val:   val,
		left:  NIL,
		right: NIL,
		p:     NIL,
	}
	x := tree.root
	if x == NIL {
		tree.root = newNode
		return
	}

	var y *node = NIL
	for x != NIL {
		y = x
		if val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}
	newNode.p = y
	if val < y.val {
		y.left = newNode
	} else {
		y.right = newNode
	}
}

// TreeDelete 从二叉查找树删除一个结点z
func (tree *binarySearchTree) TreeDelete(z *node) {
	if z.left == NIL {
		tree.transplant(z, z.right)
	} else if z.right == NIL {
		tree.transplant(z, z.left)
	} else {
		y := tree.TreeMinimum(z.right)
		if y.p != z {
			tree.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
		tree.transplant(z, y)
		y.left = z.left
		y.left.p = y
	}
}

// transplant 辅助函数,用v替换u,v成为u双亲的孩子
func (tree *binarySearchTree) transplant(u, v *node) {
	if u.p == NIL {
		tree.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	if v != NIL {
		v.p = u.p
	}
}
