package binarytree

import (
	"fmt"
	"testing"
)

func TestBinArySearchTree_Insert(t *testing.T) {
	bst := BinArySearchTree{}
	bst.Insert(12)
	bst.Insert(5)
	bst.Insert(18)
	bst.Insert(2)
	bst.Insert(9)
	bst.Insert(15)
	bst.Insert(19)
	bst.Insert(13)
	bst.Insert(17)
	if bst.root.left == nil {
		fmt.Println("func insert is wrong")
	}
}
func TestBinArySearchTree_TreeMax(t *testing.T) {
	bst := BinArySearchTree{}
	bst.Insert(12)
	bst.Insert(5)
	bst.Insert(18)
	bst.Insert(2)
	bst.Insert(9)
	bst.Insert(15)
	bst.Insert(19)
	bst.Insert(13)
	bst.Insert(17)
	max, _ := bst.TreeMax()
	if max.key != 19 {
		fmt.Println("func max is wrong", max.key)
	}
}

func TestBinArySearchTree_TreeMin(t *testing.T) {
	bst := BinArySearchTree{}
	bst.Insert(12)
	bst.Insert(5)
	bst.Insert(18)
	bst.Insert(2)
	bst.Insert(9)
	bst.Insert(15)
	bst.Insert(19)
	bst.Insert(13)
	bst.Insert(17)
	min, _ := bst.TreeMin()
	if min.key != 2 {
		fmt.Println("func min is wrong", min.key)
	}
}

func TestBinArySearchTree_TreeSearch(t *testing.T) {
	bst := BinArySearchTree{}
	bst.Insert(12)
	bst.Insert(5)
	bst.Insert(18)
	bst.Insert(2)
	bst.Insert(9)
	bst.Insert(15)
	bst.Insert(19)
	bst.Insert(13)
	bst.Insert(17)
	t1, _ := bst.TreeSearch(100)
	t2, _ := bst.TreeSearch(5)
	if t1 != nil {
		fmt.Println("func search is wrong")
	}
	if t2.key != 5 {
		fmt.Println("func search is wrong")
	}
}

func TestBinArySearchTree_successor(t *testing.T) {
	bst := BinArySearchTree{}
	bst.Insert(12)
	bst.Insert(5)
	bst.Insert(18)
	bst.Insert(2)
	bst.Insert(9)
	bst.Insert(15)
	bst.Insert(19)
	bst.Insert(13)
	bst.Insert(17)
	post := bst.root.successor()
	if post.key != 13 {
		fmt.Println("func successor is wrong")
	}
	post1 := bst.root.right.successor()
	if post1.key != 19 {
		fmt.Println("func successor is wrong")
	}
}

func TestBinArySearchTree_Delete(t *testing.T) {
	bst := BinArySearchTree{}
	bst.Insert(12)
	bst.Insert(5)
	bst.Insert(18)
	bst.Insert(2)
	bst.Insert(9)
	bst.Insert(15)
	bst.Insert(19)
	bst.Insert(13)
	bst.Insert(17)
	bst.Insert(17)
	//bst.Traverse()
	bst.Delete(12)
	bst.Delete(100)
	bst.Delete(17)
	bst.Traverse()
}
