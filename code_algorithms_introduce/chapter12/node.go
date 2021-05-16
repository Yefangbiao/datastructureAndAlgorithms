package chapter12

// 二叉查找树结点
type node struct {
	val            int
	left, right, p *node
}

// 定义一个全局结点
var (
	NIL *node = &node{}
)
