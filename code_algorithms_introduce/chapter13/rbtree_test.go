package redblacktree

import (
	"fmt"
	"log"
	"testing"
)

func TestRBTree_InsertNode(t *testing.T) {
	tree := RBTree{}
	tree.InsertNode(11)
	if tree.Root.parent != tNil {
		fmt.Println("fun insert is wrong", "tree.parent is", tree.Root.GetParent())
	}
	//Root := tree.Root
	tree.InsertNode(2)

	tree.InsertNode(14)

	tree.InsertNode(1)

	tree.InsertNode(7)
	tree.InsertNode(15)

	tree.InsertNode(5)
	tree.InsertNode(8)
	tree.InsertNode(4)
	if tree.Root.key != 7 {
		fmt.Println("fun insert is wrong")
		fmt.Println(tree.Root.key)
	}
	if tree.Root.parent != tNil {
		fmt.Println("fun insert is wrong", "tree.parent is", tree.Root.GetParent())
	}
	//打印红黑树
	//printTreeInLog(tree.Root, "(Root)")

}

// 测试删除结点
func TestRBTree_DeleteNode(t *testing.T) {
	rbtree := RBTree{}

	intarr := [...]int{1, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, num := range intarr {
		rbtree.InsertNode(num)
	}

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < 50; i++ {
	// 	rbtree.Insert(r.Int63n(100))
	// }
	log.Print("输出红黑树@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	printTreeInLog(rbtree.Root, "(Root)")

	log.Print("删除节点@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	rbtree.DeleteNode(14)
	rbtree.DeleteNode(9)
	rbtree.DeleteNode(5)
	printTreeInLog(rbtree.Root, "(Root)")
}
