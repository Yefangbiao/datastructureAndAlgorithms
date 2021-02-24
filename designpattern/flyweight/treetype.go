package main

import "fmt"

type treeType struct {
	name    string
	color   string
	texture string
}

func newTreeType(name string, color string, texture string) *treeType {
	return &treeType{name: name, color: color, texture: texture}
}

func (t *treeType) draw(x, y int) {
	fmt.Printf("在(%d,%d)创建了一棵树, name: %s, color: %s, texture: %s\n", x, y, t.name, t.color, t.texture)
}
