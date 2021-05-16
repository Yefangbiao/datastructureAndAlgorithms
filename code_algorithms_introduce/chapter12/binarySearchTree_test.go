package chapter12

import (
	"fmt"
	"testing"
)

func TestGetNewBinarySearchTree(t *testing.T) {
	tree := GetNewBinarySearchTree()
	if tree.root != NIL {
		t.Errorf("error: TestGetNewBinarySearchTree,want:%v , got:%v ", NIL, tree.root)
	}
}

func TestBinarySearchTree_TreeInsert(t *testing.T) {
	tree := GetNewBinarySearchTree()
	tree.TreeInsert(15)
	tree.TreeInsert(6)
	tree.TreeInsert(3)
	tree.TreeInsert(2)
	tree.TreeInsert(4)
	tree.TreeInsert(7)
	tree.TreeInsert(13)
	tree.TreeInsert(9)
	tree.TreeInsert(18)
	tree.TreeInsert(17)
	tree.TreeInsert(20)
	fmt.Println(tree.InorderTreeWalk())
}

func TestBinarySearchTree_TreeDelete(t *testing.T) {
	// case1
	tree := GetNewBinarySearchTree()
	tree.TreeInsert(15)
	tree.TreeInsert(6)
	tree.TreeDelete(tree.root.left)
	fmt.Println(tree.InorderTreeWalk())

	//case2
	tree = GetNewBinarySearchTree()
	tree.TreeInsert(15)
	tree.TreeInsert(18)
	tree.TreeDelete(tree.root.right)
	fmt.Println(tree.InorderTreeWalk())

	//case3
	tree = GetNewBinarySearchTree()
	tree.TreeInsert(15)
	tree.TreeInsert(18)
	tree.TreeInsert(9)
	tree.TreeDelete(tree.root)
	fmt.Println(tree.InorderTreeWalk())

	//case4
	tree = GetNewBinarySearchTree()
	tree.TreeInsert(15)
	tree.TreeInsert(18)
	tree.TreeInsert(16)
	tree.TreeInsert(17)
	tree.TreeInsert(9)
	tree.TreeInsert(20)
	tree.TreeDelete(tree.root.right)
	fmt.Println(tree.InorderTreeWalk())
}

func TestBinarySearchTree_TreeMaximum(t *testing.T) {
	tree := GetNewBinarySearchTree()
	tree.TreeInsert(15)
	tree.TreeInsert(18)
	tree.TreeInsert(16)
	tree.TreeInsert(17)
	tree.TreeInsert(9)
	tree.TreeInsert(20)
	maxNode := tree.TreeMaximum(tree.root)
	fmt.Println(maxNode.val)
}

func TestBinarySearchTree_TreeMinimum(t *testing.T) {
	tree := GetNewBinarySearchTree()
	tree.TreeInsert(15)
	tree.TreeInsert(18)
	tree.TreeInsert(16)
	tree.TreeInsert(17)
	tree.TreeInsert(9)
	tree.TreeInsert(20)
	minNode := tree.TreeMinimum(tree.root)
	fmt.Println(minNode.val)
}
