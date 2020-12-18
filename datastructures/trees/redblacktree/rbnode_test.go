package redblacktree

import (
	"fmt"
	"testing"
)

func TestRBNode_GetParent(t *testing.T) {
	p := NewRBNode(0)
	child1 := NewRBNode(1)
	p.left = child1
	child1.parent = p
	child2 := NewRBNode(2)
	p.right = child2
	child2.parent = p
	grandson := NewRBNode(3)
	child1.left = grandson
	grandson.parent = child1
	if child1.GetParent() != p {
		fmt.Printf("func GetParent is not succeed")
	}
	if child2.GetParent() != p {
		fmt.Printf("func GetParent is not succeed")
	}
	if child2.GetSibling() != child1 {
		fmt.Println("func GetSibling is not succeed")
	}
	if child1.GetSibling() != child2 {
		fmt.Println("func GetSibling is not succeed")
	}
	if grandson.GetUncle() != child2 {
		fmt.Println("func GetUncle is not succeed")
	}

}

func TestRBNode_LeftRotate(t *testing.T) {
	x := NewRBNode(0)
	y := NewRBNode(1)
	x.right = y
	y.parent = x
	root, _ := x.LeftRotate()
	if root != y {
		fmt.Print("fuc LeftRotate is wrong")
	}
	if x.parent != y {
		fmt.Print("fuc LeftRotate is wrong")
	}
}

func TestRBNode_RightRotate(t *testing.T) {
	y := NewRBNode(0)
	x := NewRBNode(1)
	y.parent = x
	y.left = x
	root, _ := y.RightRotate()
	if root != x {
		fmt.Print("func RightRotate is wrong")
	}
}

func TestRbNode_GetRightMin(t *testing.T) {
	p := NewRBNode(5)

	if p.GetRightMin() != tNil {
		fmt.Println("func GetRightMin is wrong")
	}
	p1 := NewRBNode(10)
	p1.right = tNil
	p.right = p1
	if p.GetRightMin() != p1 {
		fmt.Println("func GetRightMin is wrong")
	}
	p2 := NewRBNode(9)
	p1.left = p2
	if p.GetRightMin() != p2 {
		fmt.Println("func GetRightMin is wrong")
	}
	p3 := NewRBNode(8)
	p2.left = p3
	if p.GetRightMin() != p3 {
		fmt.Println("func GetRightMin is wrong")
	}
	p4 := NewRBNode(7)
	p3.left = p4
	if p.GetRightMin() != p4 {
		fmt.Println("func GetRightMin is wrong")
	}
}
